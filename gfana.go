package gfana

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/rkritchat/gfana/pkg"
	"net/http"
)

func New(router *chi.Mux, values map[string]string){
	router.Get("/", checkHealth)
	router.Post("/search", search)
	router.Post("/query", func(w http.ResponseWriter, r *http.Request) {
		result := pkg.Query(values)
		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(&result)
	})
}

func checkHealth(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
}

func search(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type","application/json")
	json.NewEncoder(w).Encode(pkg.InitSearch())
}