package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/googollee/go-socket.io"
	"server/internal/database"
)

type FiberServer struct {
	*fiber.App
	socketServer *socketio.Server
}

func New() *FiberServer {
	database.InitPostgres()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:  "*",
		ExposeHeaders: "*",
	}))

	socketServer := socketio.NewServer(nil)

	socketServer.OnEvent("", "CUSTOM_EVENT", func(s socketio.Conn, msg string) {
		fmt.Printf("Received CUSTOM_EVENT: %s\n", msg)
	})

	socketServer.OnConnect("", func(s socketio.Conn) error {
		fmt.Printf("New connection: %s\n", s.ID())
		return nil
	})

	socketServer.OnDisconnect("", func(s socketio.Conn, reason string) {
		fmt.Printf("Disconnected: %s, Reason: %s\n", s.ID(), reason)
	})

	app.Get("/socket.io", adaptor.HTTPHandlerFunc(socketServer.ServeHTTP))
	app.Post("/socket.io", adaptor.HTTPHandlerFunc(socketServer.ServeHTTP))

	go func() {
		if err := socketServer.Serve(); err != nil {
			log.Fatalf("Socket.IO server error: %s", err)
		}
	}()

	return &FiberServer{
		App:          app,
		socketServer: socketServer,
	}
}
