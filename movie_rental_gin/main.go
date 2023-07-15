package main

import (
	"log"
	"movierental/api"
	"movierental/config"
	"movierental/service"
)

func main() {
	config := config.GetConfig()
	movieService := service.NewMovieService(config)
	server := api.NewServer(movieService)

	if err := server.Start(config.ServerAddress); err != nil {
		log.Fatalf("Could not start server: %s", err.Error())
	}
}
