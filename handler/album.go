package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"strconv"

	"../db"
	"../schema"
	"../service"
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

type songsHandler struct {
	postgres *db.Postgres
}

func (handler *songsHandler) Addsong(w http.ResponseWriter, r *http.Request) {

	ctx := db.SetRepository(r.Context(), handler.postgres)
	b, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var song schema.Album

	if err := json.Unmarshal(b, &song); err != nil {
		responseError(w, http.StatusBadRequest, err.Error())
		return
	}
	_, newsong, err := service.Insert(ctx, &song)

	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseOk(w, newsong)

}

func (handler *songsHandler) getAlbum(w http.ResponseWriter, r *http.Request) {
	ctx := db.SetRepository(r.Context(), handler.postgres)

	album, err := service.GetAll(ctx)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseOk(w, album)

}

func (handler *songsHandler) getSong(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ctx := db.SetRepository(r.Context(), handler.postgres)
	id, _ = strconv.Atoi(params["id"])

	song, err := service.GetOne(ctx, id)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseOk(w, song)
}

func (handler *songsHandler) deleteSong(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ctx := db.SetRepository(r.Context(), handler.postgres)
	id, _ = strconv.Atoi(params["id"])

	err := service.Delete(ctx, id)
	if err != nil {
		responseError(w, http.StatusNotFound, "No song with that ID")
	}
	responseOk(w, "Song deleted successfully")

}

func (handler *songsHandler) updateSong(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ctx := db.SetRepository(r.Context(), handler.postgres)
	id, _ = strconv.Atoi(params["id"])
	b, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var song schema.Album

	if err := json.Unmarshal(b, &song); err != nil {
		responseError(w, http.StatusBadRequest, err.Error())
		return
	}
	_, newsong, err := service.UpdateSong(ctx, id, &song)

	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseOk(w, newsong)

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
