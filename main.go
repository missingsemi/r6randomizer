package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/missingsemi/r6randomizer/assets"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html><body><a href="/attacker">Attacker</a><br /><a href="/defender">Defender</a></body></html>`))
	})
	r.Get("/attacker", func(w http.ResponseWriter, r *http.Request) {
		loadout := GetRandomAttacker()
		RenderTemplate(w, loadout)
	})

	r.Get("/defender", func(w http.ResponseWriter, r *http.Request) {
		loadout := GetRandomDefender()
		RenderTemplate(w, loadout)
	})

	r.Handle("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.FS(assets.Assets))))

	http.ListenAndServe(":5000", r)
}
