package main

import (
	"mmgweb/config"
	handler "mmgweb/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func routes(app *config.AppConfig) http.Handler {

	r := chi.NewRouter()
	r.Use(NoSurf)
	r.Use(LoadSession)

	r.Get("/", http.HandlerFunc(handler.Repo.Home))
	r.Get("/About", http.HandlerFunc(handler.Repo.About))

	r.Handle("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.Dir("../../assets"))))
	//r.Handle("/assets/*", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets/"))))

	//r.Handle("/assets/css/*", http.StripPrefix("/assets/css/", http.FileServer(http.Dir("../../assets/css"))))
	//r.Handle("/assets/image/*", http.StripPrefix("/assets/image/", http.FileServer(http.Dir("../../assets/image"))))

	return r
}
