package handlers

import (
	"net/http"

	"github.com/yuhenghenrycai/go_web/pkg/config"
	"github.com/yuhenghenrycai/go_web/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepository(app *config.AppConfig) *Repository {
	return &Repository{
		App: app,
	}
}

// set repository for the handler package
func NewHandler(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.gohtml")
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.gohtml")
}
