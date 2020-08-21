package controllers

import (
	"net/http"
)

var HomePage = func(w http.ResponseWriter, r *http.Request) {
	query := r.URL.RawQuery

	w.Write([]byte(query))
}
