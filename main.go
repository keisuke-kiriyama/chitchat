package main

import "net/http"

func main() {
	p("ChitChat", version(), "started at", config.Address)

	// handle static assets
	mux := http.NewServeMux()

	// TODO: bootstrapでなんかやるっぽい。後で
	// files := http.FileServer(http.Dir(config.Static))
	// mux.Handle("/static/", http.StripPrefix("/static/", files))

	//
	// all route patterns matched here
	// route handler functions defined in other files
	//

	// index
	mux.HandleFunc("/", index)
	// error
	mux.HandleFunc("/err", err)

}
