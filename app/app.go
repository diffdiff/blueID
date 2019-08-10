package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"

	"github.com/diffdiff/blueID/app/config"
	"github.com/diffdiff/blueID/app/handler"
	"github.com/diffdiff/blueID/app/model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// APP struct
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// Initialize the app
func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Name,
		config.DB.Charset)

	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect database")
	}

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) setRouters() {
	api := a.Router.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/lang", a.CreateLang).Methods("POST")
	api.HandleFunc("/lang", a.GetAllLangs).Methods("GET")
	api.HandleFunc("/lang/{id}", a.GetLang).Methods("GET")
	api.HandleFunc("/lang/{id}", a.DeleteLang).Methods("DELETE")
	api.HandleFunc("/lang/{id}", a.UpdateLang).Methods("PUT")
	api.HandleFunc("/word", a.CreateWord).Methods("POST")
	api.HandleFunc("/word", a.GetAllWord).Methods("GET")
	api.HandleFunc("/word/{ID}", a.GetWord).Methods("GET")
	api.HandleFunc("/word/{ID}", a.DeleteWord).Methods("DELETE")
	api.HandleFunc("/word/{id}", a.UpdateWord).Methods("UPDATE")
	api.HandleFunc("/translate", a.CreateTranslate).Methods("POST")
	api.HandleFunc("/translate", a.GetAllTranslate).Methods("GET")
	api.HandleFunc("/translate/{word}", a.GetTranslate).Methods("GET")
	api.HandleFunc("/translate/{id}", a.DeleteTranslate).Methods("DELETE")
	api.HandleFunc("/translate/{id}", a.UpdateTranslate).Methods("PUT")
}

// CreateLang create new translation lang
func (a *App) CreateLang(w http.ResponseWriter, r *http.Request) {
	handler.CreateLang(a.DB, w, r)
}

//GetAllLangs get all langs
func (a *App) GetAllLangs(w http.ResponseWriter, r *http.Request) {
	handler.GetAllLangs(a.DB, w, r)
}

//GetLang get lang
func (a *App) GetLang(w http.ResponseWriter, r *http.Request) {
	handler.GetLang(a.DB, w, r)
}

//UpdateLang update lang
func (a *App) UpdateLang(w http.ResponseWriter, r *http.Request) {
	handler.UpdateLang(a.DB, w, r)
}

//DeleteLang delete lang
func (a *App) DeleteLang(w http.ResponseWriter, r *http.Request) {
	handler.DeleteLang(a.DB, w, r)
}


// CreateWord create word
func (a *App) CreateWord(w http.ResponseWriter, r *http.Request) {
	handler.CreateWord(a.DB, w, r)
}

// GetAllWord get all words
func (a *App) GetAllWord(w http.ResponseWriter, r *http.Request) {
	handler.GetAllWord(a.DB, w, r)
}

// GetWord get word
func (a *App) GetWord(w http.ResponseWriter, r *http.Request) {
	handler.GetWord(a.DB, w, r)
}

// CreateTranslate create translation mapping between words
func (a *App) CreateTranslate(w http.ResponseWriter, r *http.Request) {
	handler.CreateTranslate(a.DB, w, r)
}

// DeleteWord delete an word
func (a *App) DeleteWord(w http.ResponseWriter, r *http.Request) {
	handler.DeleteWord(a.DB, w, r)
}

// UpdateWord delete an word
func (a *App) UpdateWord(w http.ResponseWriter, r *http.Request) {
	handler.UpdateWord(a.DB, w, r)
}

// GetAllTranslate get all translations
func (a *App) GetAllTranslate(w http.ResponseWriter, r *http.Request) {
	handler.GetAllTranslate(a.DB, w, r)
}

// GetTranslate find translation
func (a *App) GetTranslate(w http.ResponseWriter, r *http.Request) {
	handler.GetTranslate(a.DB, w, r)
}

// DeleteTranslate delete translate
func (a *App) DeleteTranslate(w http.ResponseWriter, r *http.Request) {
	handler.DeleteTranslate(a.DB, w, r)
}

// UpdateTranslate update translate
func (a *App) UpdateTranslate(w http.ResponseWriter, r *http.Request) {
	handler.UpdateTranslate(a.DB, w, r)
}

// Run the main app
func (a *App) Run(host string) {
	c := cors.AllowAll()

	handler := c.Handler(a.Router)

	log.Fatal(http.ListenAndServe(host, handler))
}
