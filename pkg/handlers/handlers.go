package handlers

import (
	"log"
	"net/http"

	"github.com/Bayvao/hotel-reservation/pkg/config"
	"github.com/Bayvao/hotel-reservation/pkg/models"
	"github.com/Bayvao/hotel-reservation/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// New handlers set the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	log.Println("remote_ip: " + remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	log.Println("Get remote_ip: " + remoteIP)

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Contact page handler
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {

	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	log.Println("Get remote_ip: " + remoteIP)

	render.RenderTemplate(w, "contact.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Generals  page handler
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {

	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	log.Println("Get remote_ip: " + remoteIP)

	render.RenderTemplate(w, "generals.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Majors page handler
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {

	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	log.Println("Get remote_ip: " + remoteIP)

	render.RenderTemplate(w, "majors.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Reservation page handler
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {

	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	log.Println("Get remote_ip: " + remoteIP)

	render.RenderTemplate(w, "make-reservation.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Availability page handler
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {

	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	log.Println("Get remote_ip: " + remoteIP)

	render.RenderTemplate(w, "search-availability.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
