package handler

import (
	"net/http"
	"time"
	"github.com/gorilla/mux"
	"../db"
)

func SetUpRouting(postgres *db.Postgres) *mux.Router {
	songsHandler := &songsHandler{
		// postgres: postgres,
	}
	
	mux := mux.NewRouter()
	album = append(album, Album{ID: 1, Title: "Nic", Artist: "Raboy", DateAdded: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)})

	mux.HandleFunc("/album", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodPost:
            songsHandler.Addsong(w, r)
        case http.MethodGet:
            songsHandler.getAlbum(w, r)
        default:
            responseError(w, http.StatusNotFound, "Invalid route or request method")
        }
	})

	mux.HandleFunc("/album/{id}", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodGet:
            songsHandler.getSong(w, r)
        case http.MethodDelete:
            songsHandler.deleteSong(w, r)
        default:
            responseError(w, http.StatusNotFound, "Invalid route or request method")
        }
	})
	return mux
}
