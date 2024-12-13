package server

import "server/internal/server/handlers"

func (s *FiberServer) RegisterUserRoutes() {
	usersRouter := s.Group("/users")
	usersRouter.Use(handlers.AuthMiddleware)

	usersRouter.Get("/", handlers.GetAllUsersHandler)
	usersRouter.Get("/psychologists", handlers.GetAllPsychologistsHandler)
	usersRouter.Post("/", handlers.CreateUserHandler)
	usersRouter.Patch("/", handlers.UpdateUserHandler)
	usersRouter.Delete("/:id", handlers.DeleteUserHandler)
}

func (s *FiberServer) RegisterMessagesRoutes() {
	messagesRouter := s.Group("/messages")
	messagesRouter.Use(handlers.AuthMiddleware)

	messagesRouter.Get("/", handlers.GetAllMessagesHandler)
	messagesRouter.Post("/", handlers.CreateMessageHandler)
	messagesRouter.Patch("/", handlers.UpdateMessageHandler)
	messagesRouter.Delete("/:id", handlers.DeleteMessageHandler)
}

func (s *FiberServer) RegisterTherapyGroupRoutes() {
	therapyGroupRouter := s.Group("/therapy-group")
	therapyGroupRouter.Use(handlers.AuthMiddleware)

	therapyGroupRouter.Get("/", handlers.GetAllTherapyGroupsHandler)
	therapyGroupRouter.Post("/", handlers.CreateTherapyGroupHandler)
	therapyGroupRouter.Patch("/", handlers.UpdateTherapyGroupHandler)
	therapyGroupRouter.Delete("/:id", handlers.DeleteTherapyGroupHandler)
}

func (s *FiberServer) RegisterPrivateChatRoutes() {
	privateChatRouter := s.Group("/private-chat")
	privateChatRouter.Use(handlers.AuthMiddleware)

	privateChatRouter.Get("/", handlers.GetAllPrivateChatsHandler)
	privateChatRouter.Get("/my", handlers.GetMyPrivateChatsHandler)
	privateChatRouter.Get("/get-by-user/:id", handlers.GetPrivateChatIdByUserId)
	privateChatRouter.Get("/:id", handlers.GetPrivateChatById)
	privateChatRouter.Post("/", handlers.CreatePrivateChatHandler)
	privateChatRouter.Patch("/", handlers.UpdatePrivateChatHandler)
	privateChatRouter.Delete("/:id", handlers.DeletePrivateChatHandler)
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
