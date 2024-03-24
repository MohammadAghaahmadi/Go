package main

import (
	"fmt"
	"log"
	"mmgweb/config"
	handler "mmgweb/handlers"
	render "mmgweb/helpers"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.Name = "mmgMain"
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session
	//create a temlate cache
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Println("error creating static template: ", err)
	}
	log.Println("create template is: ", tc)
	app.TemplateCache = tc
	app.UseCache = true

	repo := handler.SetRepo(&app)
	handler.NewHandlers(repo)

	render.SetTemplates(&app)

	fmt.Println(fmt.Sprintf("Starting Application on: %d", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Println("ListenAndServe has error : ", tc)
	}

}
