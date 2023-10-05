package brimstoneesan

import (
	"fmt"
	"log"
	"os"
	"strconv"

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
