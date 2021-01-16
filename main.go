package main

import "net/http"

func main() {
	p("ChitChat", version(), "started at", config.Address)

	// handle static assets
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
}
