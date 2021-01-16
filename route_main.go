package main

import "net/http"

// GET /err?msg=
// shows the error message page
func err(writer http.ResponseWriter, request *http.Request) {
	// TODO: (仮)
	vals := request.URL.Query()
	generateHTML(writer, vals.Get("msg"), "layout", "public.navbar", "error")

	// vals := request.URL.Query()
	// _, err := session(writer, request)
	// if err != nil {
	// 	generateHTML(writer, vals.Get("msg"), "layout", "public.navbar", "error")
	// } else {
	// 	generateHTML(writer, vals.Get("msg"), "layout", "private.navbar", "error")
	// }
}
