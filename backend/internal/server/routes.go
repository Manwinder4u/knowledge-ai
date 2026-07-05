package server

func (s *Server) registerRoutes() {
	s.router.Get("/health", s.health)
}
