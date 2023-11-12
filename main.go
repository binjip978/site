package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	"golang.org/x/crypto/acme/autocert"
)

//go:embed static
var static embed.FS

func main() {
	mux := http.NewServeMux()
	static, err := fs.Sub(static, "static")
	if err != nil {
		log.Fatal(err)
	}
	mux.Handle("/", http.FileServer(http.FS(static)))
	log.Fatal(http.Serve(autocert.NewListener("binjip978.net"), mux))
}
