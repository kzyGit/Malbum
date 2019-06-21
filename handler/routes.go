package handler

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func SetUpRouting() *mux.Router {
	songsHandler := &songsHandler{}
	usersHandler := &usersHandler{}

	mux := mux.NewRouter()
	album = append(album, Album{ID: 1, Title: "Ayo", Artist: "Simi", DateAdded: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)})
	users = append(users, Users{ID: 1, Username: "kezzy", Password: "kezzy"})

	mux.HandleFunc("/album", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			songsHandler.getAlbum(w, r)
		case http.MethodPost:
			songsHandler.Addsong(w, r)
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

	mux.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			usersHandler.userSignup(w, r)
		default:
			responseError(w, http.StatusNotFound, "Invalid request method")
		}
	})

	return mux
}
