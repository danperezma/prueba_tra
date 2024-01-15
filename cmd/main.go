package main

import (
	"fmt"
	"net/http"
	"back_go/pkg/handler"
	"back_go/pkg/api"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

func main() {
	handler.Index()

	r := chi.NewRouter()
	port := 8080

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})

	r.Use(corsMiddleware.Handler)
	r.Post("/api/query", api.HandlerPostQuery)

	fmt.Printf("Server listening on port %d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
