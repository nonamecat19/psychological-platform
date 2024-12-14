package main

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"server/internal/config"
	"server/internal/server"
)

func main() {
	appConfig := config.GetAppConfig()

	api := server.New()
	api.RegisterFiberRoutes()
	api.RegisterFiberWsRoutes()

	err := api.Listen(fmt.Sprintf(":%d", appConfig.Port))
	log.Print(fmt.Sprintf("Server started on port: %d", appConfig.Port))
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
