package server

import (
	"fmt"
	"net/http"

	"github.com/deanrtaylor1/backend-go/internal/config"
	"github.com/deanrtaylor1/backend-go/internal/controllers/basecontrollers"
	"github.com/deanrtaylor1/backend-go/internal/controllers/domain/authcontrollers"
	db "github.com/deanrtaylor1/backend-go/internal/db/sqlc"
	"github.com/deanrtaylor1/backend-go/internal/middleware"
	"github.com/go-chi/chi"
)

type Server struct {
	Config config.EnvConfig
	Store  db.Store
	Router *chi.Mux
}

func (s *Server) Start() {

	http.ListenAndServe(fmt.Sprintf(":%s", s.Config.Port), s.Router)
}

func NewServer(config config.EnvConfig, store db.Store, router *chi.Mux) *Server {
	return &Server{
		Config: config,
		Store:  store,
		Router: router,
	}
}

func (s *Server) RegisterMiddlewares() {
	s.Router.Use(middleware.ColorLoggingMiddleware)
}

func (s *Server) RegisterRoutes(router *chi.Mux) {
	s.Router.Route("/api/v1", func(r chi.Router) {
		// r.Mount("/users", routes.UsersRouter())
		baseController := basecontrollers.NewBaseController(s.Store, s.Config)
		authController := authcontrollers.NewAuthController(*baseController)
		r.Mount("/auth", authController.Routes())

	})
}
