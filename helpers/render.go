package render

import (
	"bytes"
	"html/template"
	"log"
	"mmgweb/config"
	"mmgweb/models"
	"net/http"
	"path/filepath"
)

func AddDefaultData(td *models.TemplateData) *models.TemplateData {

	return td
}

var app *config.AppConfig

func SetTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tpml string, td *models.TemplateData) {

	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	//get the right template from cash
	t, ok := tc[tpml]
	if !ok {
		log.Println("error creating static template: ")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	err := t.Execute(buf, td)
	if err != nil {
		log.Println("error parsing template: ", err)
	}

	//remder that template

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println("error parsing template: ", err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	theCache := map[string]*template.Template{}

	pages, err := filepath.Glob("../../pages/*.html")
	if err != nil {
		return theCache, err
	}

	for _, page := range pages {
		log.Println(page + " found ")
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return theCache, err
		}

		matches, err := filepath.Glob("../../layouts/*.html")
		if err != nil {
			return theCache, err
		}
		//چون هر فایلی ممکنه لی اوت خودش رو داشته باشه
		// پس جدا باید لی اوتش لود بشه و فایل اصلی در کش قرار بگیره
		log.Println("matches parsing template: ", matches)
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("../../layouts/*.html")
			if err != nil {
				return theCache, err
			}
		}

		theCache[name] = ts
	}
	return theCache, nil

}
