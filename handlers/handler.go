package handler

import (
	"mmgweb/config"
	render "mmgweb/helpers"
	"mmgweb/models"
	"net/http"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

func SetRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home-page.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	//m.App.Session Access to the Session

	sideKickMap := make(map[string]string)
	sideKickMap["morty"] = "Ooh, wee!"
	sideKickMap["remote_ip"] = m.App.Session.GetString(r.Context(), "remote_ip")

	render.RenderTemplate(w, "about-page.html", &models.TemplateData{StringMap: sideKickMap})
}
