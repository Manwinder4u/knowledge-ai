package middleware

import (
	"net/http"

	"github.com/Manwinder4u/knowledge-ai/backend/internal/logger"
	"github.com/Manwinder4u/knowledge-ai/backend/internal/response"
)

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			if err := recover(); err != nil {

				logger.Log.Error().
					Any("panic", err).
					Str("method", r.Method).
					Str("path", r.URL.Path).
					Msg("Recovered from panic")

				response.Error(w, http.StatusInternalServerError, "Internal server error")
			}
		}()

		next.ServeHTTP(w, r)
	})
}
