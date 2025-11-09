package main

import (
	"github.com/ilam072/profiler/internal/rest"
	"log"
	"net/http"
)

func main() {
	handler := rest.NewHandler()

	http.HandleFunc("/sum", handler.Sum)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
