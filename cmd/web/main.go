package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"ivixlabs.com/goweb/internal/app"
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

	clickhouseHost := os.Getenv("CLICKHOUSE_HOST")
	if clickhouseHost == "" {
		clickhouseHost = "localhost"
	}

	var clickhousePort uint64 = 9000
	strClickhousePort := os.Getenv("CLICKHOUSE_PORT")
	if strClickhousePort != "" {
		var err error
		clickhousePort, err = strconv.ParseUint(strClickhousePort, 10, 32)
		if err != nil {
			panic(err)
		}
	}

	clickhouseAddr := []string{fmt.Sprintf("%s:%d", clickhouseHost, clickhousePort)}

	app.Run(":"+port, staticDir, dbUrl, sessionsDir, developmentMode, clickhouseAddr)
}
