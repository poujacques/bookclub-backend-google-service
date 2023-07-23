package main

import (
	"googleBooks"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Route("/volumes", func(r chi.Router) {
		r.Get("/", googleBooks.GetVolumesResult)
		r.Get("/{volumeID}", googleBooks.GetVolumeByID)
	})
	log.Print("Listening...")
	http.ListenAndServe(":8080", r)
}
