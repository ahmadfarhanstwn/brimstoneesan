package render

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type Render struct {
	Renderer   string
	RootPath   string
	Secure     bool
	Port       string
	ServerName string
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

	}

	return nil
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

func (rend *Render) JetPage(w http.ResponseWriter, r *http.Request, view string, data interface{}) error {
	return nil
}
