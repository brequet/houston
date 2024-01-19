package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strconv"
)

//go:embed front/dist/*
var content embed.FS
var webAppPort int

func main() {
	if len(os.Args) > 1 {
		port, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatalf("Invalid port number: %v", err)
		}
		webAppPort = port
	} else {
		// Default to port 3000 if no argument is provided
		webAppPort = 3000
	}

	serveSpa(webAppPort)
}

func serveSpa(port int) {
	webPage, err := fs.Sub(content, "front/dist")
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", http.FileServer(http.FS(webPage)))

	listenPort := fmt.Sprintf(":%d", port)
	log.Printf("Listening on %s...\n", listenPort)
	err = http.ListenAndServe(listenPort, nil)
	if err != nil {
		log.Fatal(err)
	}
}
