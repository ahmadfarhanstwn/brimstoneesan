package brimstoneesan

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/CloudyKit/jet/v6"
	"github.com/ahmadfarhanstwn/brimstoneesan/render"
	"github.com/ahmadfarhanstwn/brimstoneesan/session"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

const version = "1.0.0"

type Brimstoneesan struct {
	AppName  string
	Debug    bool
	Version  string
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	RootPath string
	Routes   *chi.Mux
	Render   *render.Render
	Session  *scs.SessionManager
	JetView  *jet.Set
	config   config
}

type config struct {
	port        string
	renderer    string
	cookie      cookieConfig
	sessionType string
}

func (b *Brimstoneesan) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath:    rootPath,
		folderNames: []string{"handlers", "migrations", "views", "data", "public", "tmp", "logs", "middleware"},
	}

	err := b.Init(pathConfig)
	if err != nil {
		return err
	}

	err = b.checkForEnv(rootPath)
	if err != nil {
		return err
	}

	//read env
	err = godotenv.Load(rootPath + "/.env")
	if err != nil {
		return err
	}

	//start logger
	infoLog, errorLog := b.startLoggers()
	b.InfoLog = infoLog
	b.ErrorLog = errorLog
	b.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	b.Version = version
	b.RootPath = rootPath
	b.Routes = b.routes().(*chi.Mux)
	b.config = config{
		port:     os.Getenv("PORT"),
		renderer: os.Getenv("RENDERER"),
		cookie: cookieConfig{
			name:     os.Getenv("COOKIE NAME"),
			lifetime: os.Getenv("COOKIE_LIFETIME"),
			persist:  os.Getenv("COOKIE_PERSIST"),
			secure:   os.Getenv("COOKIE_SECURE"),
			domain:   os.Getenv("COOKIE_DOMAIN"),
		},
		sessionType: os.Getenv("SESSION_TYPE"),
	}

	//CREATE SESSION
	sess := session.Session{
		CookieLifetime: b.config.cookie.lifetime,
		CookiePersist:  b.config.cookie.persist,
		CookieName:     b.config.cookie.name,
		SessionType:    b.config.sessionType,
		CookieDomain:   b.config.cookie.domain,
	}

	b.Session = sess.InitSession()

	views := jet.NewSet(
		jet.NewOSFileSystemLoader(fmt.Sprintf("%s/views", rootPath)),
		jet.InDevelopmentMode(),
	)

	b.JetView = views

	b.CreateRenderer()

	return nil
}

func (b *Brimstoneesan) Init(p initPaths) error {
	root := p.rootPath
	for _, path := range p.folderNames {
		err := b.CreateDirIfNotExist(root + "/" + path)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *Brimstoneesan) ListenAndServer() {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", os.Getenv("PORT")),
		ErrorLog:     b.ErrorLog,
		Handler:      b.Routes,
		IdleTimeout:  30 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 600 * time.Second,
	}

	b.InfoLog.Printf("Listening on port %s", os.Getenv("PORT"))
	err := srv.ListenAndServe()
	b.ErrorLog.Fatal(err)
}

func (b *Brimstoneesan) checkForEnv(path string) error {
	err := b.CreateFileIfNotExist(fmt.Sprintf("%s/.env", path))
	if err != nil {
		return err
	}

	return nil
}

func (b *Brimstoneesan) startLoggers() (*log.Logger, *log.Logger) {
	var infoLog *log.Logger
	var errorLog *log.Logger

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	return infoLog, errorLog
}

func (b *Brimstoneesan) CreateRenderer() {
	myRenderer := render.Render{
		RootPath: b.RootPath,
		Renderer: b.config.renderer,
		Port:     b.config.port,
		JetViews: b.JetView,
	}

	b.Render = &myRenderer
}
