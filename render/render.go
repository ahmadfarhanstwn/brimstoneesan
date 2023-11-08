package render

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/CloudyKit/jet/v6"
	"github.com/alexedwards/scs/v2"
)

type Render struct {
	Renderer   string
	RootPath   string
	Secure     bool
	Port       string
	JetViews   *jet.Set
	ServerName string
	Session    *scs.SessionManager
}

type TemplateData struct {
	IsAuthenticated bool
	IntMap          map[string]int
	StringMap       map[string]string
	FloatMap        map[string]float32
	Data            map[string]interface{}
	CSRFToken       string
	Port            string
	ServerName      string
	Secure          bool
}

func (rend *Render) Page(w http.ResponseWriter, r *http.Request, view string, variables, data interface{}) error {
	switch strings.ToLower(rend.Renderer) {
	case "go":
		return rend.GoPage(w, r, view, data)
	case "jet":
		return rend.JetPage(w, r, view, variables, data)
	default:
	}

	return errors.New("No rendering engine specified")
}

func (rend *Render) GoPage(w http.ResponseWriter, r *http.Request, view string, data interface{}) error {
	template, err := template.ParseFiles(fmt.Sprintf("%s/views/%s.page.tmpl", rend.RootPath, view))
	if err != nil {
		return err
	}

	templateData := &TemplateData{}
	if data != nil {
		templateData = data.(*TemplateData)
	}

	err = template.Execute(w, &templateData)
	if err != nil {
		return err
	}

	return nil
}

func (rend *Render) JetPage(w http.ResponseWriter, r *http.Request, view string, variables, data interface{}) error {
	var vars jet.VarMap

	if variables == nil {
		vars = make(jet.VarMap)
	} else {
		vars = variables.(jet.VarMap)
	}

	templateData := &TemplateData{}
	if data != nil {
		templateData = data.(*TemplateData)
	}

	template, err := rend.JetViews.GetTemplate(fmt.Sprintf("%s.jet", view))
	if err != nil {
		log.Println(err)
		return err
	}

	if err = template.Execute(w, vars, templateData); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
