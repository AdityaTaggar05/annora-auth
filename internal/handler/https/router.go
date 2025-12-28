package https

import (
	authhandler "github.com/AdityaTaggar05/annora-auth/internal/handler/https/auth"
	tokenhandler "github.com/AdityaTaggar05/annora-auth/internal/handler/https/token"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(authHandler *authhandler.Handler, tokenHandler *tokenhandler.Handler) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/auth/register", authHandler.HandleRegister)
	r.Post("/auth/login", authHandler.HandleLogin)
	r.Post("/auth/logout", authHandler.HandleLogout)
	r.Post("/auth/refresh", tokenHandler.HandleRefresh)
	r.Get("/.well-known/jwks.json", tokenHandler.HandleJWKS)
	
	return r
}