package main

import "net/http"

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("index"))
}
