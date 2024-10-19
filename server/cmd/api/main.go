package main

import (
	"fmt"
	"os"
	"server/internal/server"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

func main() {

	api := server.New()

	api.RegisterFiberRoutes()
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	err := api.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
