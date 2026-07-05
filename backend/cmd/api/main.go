package main

import (
	"fmt"
	"net/http"

	"github.com/Manwinder4u/knowledge-ai/backend/internals/config"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "KnowledgeAI Backend is running 🚀")
	})

	cfg := config.Load()
	fmt.Printf("%s started on port %s\n", cfg.AppName, cfg.Port)
	http.ListenAndServe(":"+cfg.Port, r)
}
