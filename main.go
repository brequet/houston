package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

var (
	//go:embed frontend/dist/*
	content embed.FS

	webAppPort int
)

func main() {
	if len(os.Args) > 1 {
		port, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatalf("Invalid port number: %v", err)
		}
		webAppPort = port
	} else {
		// Default
		webAppPort = 3000
	}

	r := mux.NewRouter() // Create a new Gorilla Mux router
	handleSpa(r)         // Pass the router to handleSpa
	handleApi(r)         // Pass the router to handleApi

	webAppPort := fmt.Sprintf(":%d", webAppPort)
	log.Printf("Listening on %s...\n", webAppPort)
	err := http.ListenAndServe(webAppPort, corsMiddleware(r))
	if err != nil {
		log.Fatal(err)
	}
}

// CORS middleware to handle preflight requests
func corsMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173") // or "*" for any origin
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		// Check if it's a preflight request
		if r.Method == "OPTIONS" {
			// If so, handle the preflight and stop the chain
			w.WriteHeader(http.StatusOK)
			return
		}

		// Next
		handler.ServeHTTP(w, r)
	})
}

func handleSpa(r *mux.Router) {
	webPage, err := fs.Sub(content, "frontend/dist")
	if err != nil {
		log.Fatal(err)
	}

	r.Handle("/", http.FileServer(http.FS(webPage)))
}

func handleApi(r *mux.Router) {
	r.HandleFunc("/api/scripts", handleGetScripts).Methods("GET")
	r.HandleFunc("/api/script/execute", handlePostExecuteScript).Methods("POST")

	r.HandleFunc("/api/services", handleGetServices).Methods("GET")
}

func handleGetScripts(w http.ResponseWriter, r *http.Request) {
	scripts, err := GetAllScripts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(scripts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func handlePostExecuteScript(w http.ResponseWriter, r *http.Request) {
	var reqBody struct {
		ScriptName string `json:"scriptName"`
	}
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	output, err := ExecuteScriptByName(reqBody.ScriptName)
	if err != nil {
		log.Println("ExecuteScriptByName went wrong:", output, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"output":     output,
		"scriptName": reqBody.ScriptName,
	})
}

func handleGetServices(w http.ResponseWriter, r *http.Request) {
	services, err := GetAllServices()
	if err != nil {
		log.Println("Could not fetch all services", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(services)
	if err != nil {
		log.Println("Could marshal services", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
