package main

import (
	_ "embed"
	"log"
	"net/http"

	"golang.org/x/crypto/acme/autocert"
)

//go:embed content/index.html
var index []byte

//go:embed content/styles.css
var style []byte

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		log.Println("/index.html")
		w.Header().Set("Content-Type", "text/html")
		w.Write(index)
	})
	mux.HandleFunc("/styles.css", func(w http.ResponseWriter, _ *http.Request) {
		log.Println("/styles.css")
		w.Header().Set("Content-Type", "text/css")
		w.Write(style)
	})

	log.Fatal(http.Serve(autocert.NewListener("binjip978.net"), mux))
}
