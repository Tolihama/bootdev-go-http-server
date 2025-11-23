package main

import (
	"io"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/app/", http.StripPrefix("/app/", http.FileServer(http.Dir("./server"))))
	mux.HandleFunc("/healthz", healthz)
	server := http.Server{}
	server.Addr = ":8080"
	server.Handler = mux
	server.ListenAndServe()
}

func healthz(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "OK")
}
