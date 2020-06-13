package gfana

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/rkritchat/gfana/pkg"
	"net/http"
)

func New(router *chi.Mux, valueFunc func() map[string]string){
	router.Get("/", checkHealth)
	router.Post("/search", search)
	router.Post("/query", func(w http.ResponseWriter, r *http.Request) {
		resp(w, pkg.Query(valueFunc))
	})
}

func checkHealth(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
}

func search(w http.ResponseWriter, r *http.Request){
	resp(w, pkg.InitSearch())
}

func resp(w http.ResponseWriter, resp interface{}){
	w.Header().Set("content-type","application/json")
	json.NewEncoder(w).Encode(resp)
}