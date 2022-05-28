package main

import (
	"log"
	"os"

	"github.com/batt0s/Ninja-Manga-Api/handlers"
)

func main() {
	log.Print("[info] Starting app...")
	appMode := os.Getenv("APP_MODE")
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("[warning] No PORT in env. Defaulting to 8080.")
		port = "8080"
	}
	app := handlers.App{}
	app.Init(appMode)
	log.Print(" OK\n")
	log.Println("[info] Starting server.")
	app.Addr = ":" + port
	app.Run()
}
