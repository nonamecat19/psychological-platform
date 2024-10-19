package server

import "server/internal/server/routes"

func (s *FiberServer) RegisterFiberRoutes() {
	routes.ConnectUsersRouter(s)
}
