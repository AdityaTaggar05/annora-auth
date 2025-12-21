package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/AdityaTaggar05/connectify-auth/internal/config"
	"github.com/AdityaTaggar05/connectify-auth/internal/database"
	"github.com/joho/godotenv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// LOADING ENV VARIABLES
	err := godotenv.Load()
	if err != nil {
		log.Fatal("[ERR] Error loading .env file")
	}
	cfg := config.Load()

	// LOADING DATABASE
	ctx := context.Background()
	_ = database.Connect(ctx, cfg.DB_URL)

	// SETUP ROUTER & ROUTES
	r := chi.NewRouter()

	r.Use(middleware.Logger)
  	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	fmt.Printf("[DEBUG] Serving on PORT: %s\n", cfg.PORT)
	http.ListenAndServe(":" + cfg.PORT, r)
}