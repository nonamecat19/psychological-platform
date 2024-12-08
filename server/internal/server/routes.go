package server

import "server/internal/server/handlers"

func (s *FiberServer) RegisterUserRoutes() {
	usersRouter := s.Group("/users")

	usersRouter.Get("/", handlers.GetAllUsersHandler)
	usersRouter.Get("/psychologists", handlers.GetAllPsychologistsHandler)
	usersRouter.Post("/", handlers.CreateUserHandler)
	usersRouter.Patch("/", handlers.UpdateUserHandler)
	usersRouter.Delete("/:id", handlers.DeleteUserHandler)
}

func (s *FiberServer) RegisterMessagesRoutes() {
	usersRouter := s.Group("/messages")

	usersRouter.Get("/", handlers.GetAllMessagesHandler)
	usersRouter.Post("/", handlers.CreateMessageHandler)
	usersRouter.Patch("/", handlers.UpdateMessageHandler)
	usersRouter.Delete("/:id", handlers.DeleteMessageHandler)
}

func (s *FiberServer) RegisterTherapyGroupRoutes() {
	usersRouter := s.Group("/therapy-group")

	usersRouter.Get("/", handlers.GetAllTherapyGroupsHandler)
	usersRouter.Post("/", handlers.CreateTherapyGroupHandler)
	usersRouter.Patch("/", handlers.UpdateTherapyGroupHandler)
	usersRouter.Delete("/:id", handlers.DeleteTherapyGroupHandler)
}

func (s *FiberServer) RegisterPrivateChatRoutes() {
	usersRouter := s.Group("/therapy-group")

	usersRouter.Get("/", handlers.GetAllPrivateChatsHandler)
	usersRouter.Post("/", handlers.CreatePrivateChatHandler)
	usersRouter.Patch("/", handlers.UpdatePrivateChatHandler)
	usersRouter.Delete("/:id", handlers.DeletePrivateChatHandler)
}

func (s *FiberServer) RegisterAuthRoutes() {
	authRoutes := s.Group("/auth")

	authRoutes.Post("/login", handlers.AuthLoginHandler)
	authRoutes.Post("/register", handlers.AuthRegisterHandler)
}

func (s *FiberServer) RegisterFiberRoutes() {
	s.RegisterUserRoutes()
	s.RegisterMessagesRoutes()
	s.RegisterTherapyGroupRoutes()
	s.RegisterPrivateChatRoutes()
	s.RegisterAuthRoutes()
}
