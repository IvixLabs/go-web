package main

import (
	"ivixlabs.com/goweb/internal/app"
	"log"
	"os"
)

func main() {
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

	developmentMode := false
	strDevelopmentMode := os.Getenv("DEVELOPMENT_MODE")
	if strDevelopmentMode != "" {
		developmentMode = true
	}

	app.Run(":"+port, staticDir, dbUrl, sessionsDir, developmentMode)
}
