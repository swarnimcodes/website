package handlers

import "net/http"

func HandleHello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	msg := "hello, world"
	w.Write([]byte(msg))
	return
}
