package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"strconv"

	"github.com/gorilla/mux"
)

type Album struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Artist    string    `json:"artist"`
	DateAdded time.Time `json:"dateadded"`
}

var album []Album

var id = 1

type songsHandler struct{}

func (handler *songsHandler) Addsong(w http.ResponseWriter, req *http.Request) {
	var song Album
	_ = json.NewDecoder(req.Body).Decode(&song)
	song.ID = id + 1
	id++
	song.DateAdded = time.Now()
	album = append(album, song)
	json.NewEncoder(w).Encode(album)
}

func (handler *songsHandler) getAlbum(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(album)

}

func (handler *songsHandler) getSong(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range album {

		id, _ = strconv.Atoi(params["id"])
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func (handler *songsHandler) deleteSong(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for i, item := range album {
		id, _ = strconv.Atoi(params["id"])
		if item.ID == id {
			album = append(album[:i], album[i+1:]...)
			responseOk(w, "Song deleted successfully")
			return
		}
	}
	responseError(w, http.StatusNotFound, "No song with that ID")

}

func responseOk(w http.ResponseWriter, body interface{}) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(body)
}

func responseError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	body := map[string]string{
		"error": message,
	}
	json.NewEncoder(w).Encode(body)
}
