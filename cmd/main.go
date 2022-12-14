package main

import (
	"github.com/agirenko/hello-api/handlers/rest"
	"log"
	"net/http"
)

func main() {
	addr := ":8080"
	mux := http.NewServeMux()
	mux.HandleFunc("/translate/hello", rest.TranslateHandler)
	log.Printf("listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
