package server

import "server/internal/server/handlers"

func (s *FiberServer) RegisterUserRoutes() {
	usersRouter := s.Group("/users")

	usersRouter.Get("/", handlers.GetAllUsersHandler)
	usersRouter.Post("/", handlers.CreateUserHandler)
	usersRouter.Patch("/:id", handlers.UpdateUserHandler)
	usersRouter.Delete("/:id", handlers.DeleteUserHandler)
}

func (s *FiberServer) RegisterFiberRoutes() {
	s.RegisterUserRoutes()
}
