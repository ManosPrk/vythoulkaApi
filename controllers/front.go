package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterControllers() {

	router := mux.NewRouter()
	router.StrictSlash(true)
	router.Use(ContentTypeMiddleware)

	subrouter := router.PathPrefix("/foods").Subrouter()

	newFoodController(subrouter)

	http.Handle("/", router)
}

func ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
