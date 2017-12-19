package server

import (
	"fmt"
	"net/http"
)

func VersionHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("0.0.1"))
}

func NewApplicationServer(port int) *http.Server {

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/version", VersionHandler)

	return &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}
}
