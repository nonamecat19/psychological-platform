package main

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"server/internal/config"
	"server/internal/server"
)

func main() {
	appConfig := config.GetAppConfig()

	api := server.New()
	api.RegisterFiberRoutes()
	err := api.Listen(fmt.Sprintf(":%d", appConfig.Port))
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
