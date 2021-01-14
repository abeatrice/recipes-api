package main

import "net/http"

// Show ...
func Show(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("show"))
}
