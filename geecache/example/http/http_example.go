package main

import (
	"log"
	"net/http"
)

type server struct {
}

func (h *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	w.Write([]byte("Hello"))
}

func main() {
	var s server
	http.ListenAndServe("localhost:9999", &s)
}
