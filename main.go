package main

import (
	"embed"
	"flag"
	"io/fs"
	"log"
	"net/http"

	"golang.org/x/crypto/acme/autocert"
)

//go:embed static
var static embed.FS

func main() {
	var local = flag.Bool("local", false, "run locally")
	flag.Parse()
	mux := http.NewServeMux()
	static, err := fs.Sub(static, "static")
	if err != nil {
		log.Fatal(err)
	}
	mux.Handle("/", http.FileServer(http.FS(static)))
	if *local {
		log.Println("running locally")
		log.Fatal(http.ListenAndServe(":8080", mux))
	}
	log.Fatal(http.Serve(autocert.NewListener("binjip978.net"), mux))
}
