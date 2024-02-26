package main

import (
	"log"
	"os"

	"ivixlabs.com/goweb/internal/http"
)

func main() {
	log.Println("Web is started")

	dbUrl := os.Getenv("DATABASE_URL")
	log.Println("DbUrl", dbUrl)
	if dbUrl == "" {
		panic("No db url")
	}

	staticDir := os.Getenv("STATIC_DIR")
	if staticDir == "" {
		panic("No static dir")
	}

	sessionsDir := os.Getenv("SESSIONS_DIR")
	if sessionsDir == "" {
		panic("No sessions dir")
	}

	if _, err := os.Stat(sessionsDir); err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(sessionsDir, os.FileMode(0644))
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Port", port)
	log.Println("StaticDir", staticDir)
	log.Println("SessionsDir", sessionsDir)

	http.StartServer(":"+port, staticDir, dbUrl, sessionsDir)
}
