// app.go

package main

import (
	"database/sql"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"net/http"
	"strconv"
)

type App struct {
	Router *chi.Mux
	DB     *gorm.DB
}

func (a *App) Initialize(user, password, dbname, disable string) {
	//connectionString :=
	//	fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s ", user, password, dbname, disable)

	var err error
	a.DB, err = gorm.Open("postgres", "sslmode=disable host=127.0.0.1 port=5432 user=root dbname=test password=12345")
	if err != nil {
		log.Fatal(err)
	}

	a.Router = chi.NewRouter()
	a.initializeRoutes()
}

//func (a *App) Run(addr string) {
//	log.Printf("Run server")
//	log.Fatal(http.ListenAndServe(":8000", a.Router))
//}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (a *App) getShelfs(w http.ResponseWriter, r *http.Request) {

	shelfs, err := getAllShelfs(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, shelfs)
}

func (a *App) geAllBooks(w http.ResponseWriter, r *http.Request) {
	var err error
	//respondWithJSON(w, http.StatusOK, p)
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	log.Printf("args: %s", id)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	p := &Shelf{NumShelf: id}
	if err := p.getBooks(a.DB.First(p)); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Product not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, p)

}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/shelfs", a.getShelfs)
	a.Router.HandleFunc("/books/{id:[0-9]+}", a.geAllBooks)
}
