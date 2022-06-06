package main

import (
	"deck-api/api/handlers"
	"deck-api/config/db"
	"deck-api/config/http"
	"deck-api/config/log"
	"deck-api/repositories"
	"deck-api/usecases"

	"go.uber.org/zap"
)

func main() {
	log.SetupLogger()

	conn, err := db.SetupDB()
	if err != nil {
		zap.S().Fatalf("Postgres connection has failed: %s", err.Error())
	}

	deckRepo := repositories.NewDeckRepository(conn)
	deckUc := usecases.NewDeckUseCase(deckRepo)
	deckHandler := handlers.NewDeckHandler(deckUc)

	server := http.SetupHTTPServer(deckHandler)
	zap.S().Info("Starting deck-api http server on port 8080...")
	if err := server.ListenAndServe(); err != nil {
		zap.S().Errorf("Something went wrong starting http server %s", err.Error())
	}
}
