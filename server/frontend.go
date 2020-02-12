package server

import (
	"net/http"
)

//ServeFrontend serves the react frontend
func ServeFrontend(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
