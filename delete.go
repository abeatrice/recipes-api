package main

import "net/http"

// Delete ...
func Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete"))
}
