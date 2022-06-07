package http

import (
	"net/http"
	"time"

	"deck-api/api/handlers"

	"github.com/gin-gonic/gin"
)

// SetupHTTPServer configure a new router for Handler implementations (to make it easy to add more Handlers).
func SetupHTTPServer(handlers ...handlers.Handler) *http.Server {
	router := gin.Default()
	for _, handler := range handlers {
		handler.Routes(router)
	}

	// TODO: Include env vars.
	return &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
}
