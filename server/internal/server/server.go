package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/googollee/go-socket.io"
	"log"
	"server/internal/database"
)

type FiberServer struct {
	*fiber.App
	SocketServer *socketio.Server
}

func New() *FiberServer {
	database.InitPostgres()

	app := fiber.New()

	app.Use(cors.New(cors.Config{}))

	socketServer := socketio.NewServer(nil)

	app.Get("/socket.io", adaptor.HTTPHandlerFunc(socketServer.ServeHTTP))
	app.Post("/socket.io", adaptor.HTTPHandlerFunc(socketServer.ServeHTTP))

	go func() {
		if err := socketServer.Serve(); err != nil {
			log.Fatalf("Socket.IO server error: %s", err)
		}
	}()

	return &FiberServer{
		App:          app,
		SocketServer: socketServer,
	}
}
