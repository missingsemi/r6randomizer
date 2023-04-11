package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html><body><a href="/attacker">Attacker</a><br /><a href="/defender">Defender</a></body></html>`))
	})
	r.Get("/attacker", func(w http.ResponseWriter, r *http.Request) {
		loadout := GetRandomAttacker()
		w.Write([]byte(loadout.String()))
	})

	r.Get("/defender", func(w http.ResponseWriter, r *http.Request) {
		loadout := GetRandomDefender()
		w.Write([]byte(loadout.String()))
	})

	http.ListenAndServe(":5000", r)
}
