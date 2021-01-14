package main

import "net/http"

// Store ...
func Store(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("store"))
}
