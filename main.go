package main

import (
	"net/http"
	"time"
)

func main() {
	p("ChitChat", version(), "started at", config.Address)

	// handle static assets
	mux := http.NewServeMux()

	// TODO: 後でCSS系のいじったら必要になる？
	// files := http.FileServer(http.Dir(config.Static))
	// mux.Handle("/static/", http.StripPrefix("/static/", files))

	//
	// all route patterns matched here
	// route handler functions defined in other files
	//

	// index
	// mux.HandleFunc("/", index)
	// error
	mux.HandleFunc("/err", err)

	// starting up the server
	server := &http.Server{
		Addr:           config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
