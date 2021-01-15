package main

import "net/http"

// Update ...
func Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("update"))
}
