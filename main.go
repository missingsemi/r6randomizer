package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/missingsemi/r6randomizer/assets"
)

var (
	port uint
)

func init() {
	flag.UintVar(&port, "port", 8080, "Port to host the website on")
	flag.Parse()
}

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

	http.ListenAndServe(fmt.Sprintf(":%v", port), r)
}
