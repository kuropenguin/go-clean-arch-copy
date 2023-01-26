package main

import (
	"log"
	"net/http"

	"github.com/kuropenguin/go-clean-arch-copy/handler"
)

func main() {
	http.HandleFunc("/", handler.Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
