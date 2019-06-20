package handler

import (
	"encoding/json"
	// "io/ioutil"
	"net/http"
	"time"

	"strconv"

	// "../db"
	// "../schema"
	// "../service"

	"github.com/gorilla/mux"
)

type Album struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Artist    string    `json:"artist"`
	DateAdded time.Time `json:"dateadded"`
}

var album []Album

var id = 0

type songsHandler struct {
	// postgres *db.Postgres
}

func (handler *songsHandler) Addsong(w http.ResponseWriter, req *http.Request) {
	var song Album
	_ = json.NewDecoder(req.Body).Decode(&song)
	song.ID = id + 1
	id++
	song.DateAdded = time.Now()
	album = append(album, song)
	json.NewEncoder(w).Encode(album)

}

// func (handler *songsHandler) Addsong(w http.ResponseWriter, r *http.Request) {
// 	ctx := db.SetRepository(r.Context(), handler.postgres)

// 	b, err := ioutil.ReadAll(r.Body)

// 	panic(err)
// 	if err != nil {
// 		responseError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	var song schema.Album
// 	if err := json.Unmarshal(b, &song); err != nil {
// 		responseError(w, http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	id, err := service.Insert(ctx, &song)
// 	if err != nil {
// 		responseError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	responseOk(w, id)
// }

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
