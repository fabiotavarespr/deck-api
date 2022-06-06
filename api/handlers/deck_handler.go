package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"

	dtos "deck-api/api/dtos/deck"
	"deck-api/usecases"

	"github.com/gin-gonic/gin"
)

const (
	CardsParam   = "cards"
	ShuffleParam = "shuffle"
	CountParam   = "count"
	IDParam      = "id"

	ErrInvalidShuffleParam = "Invalid shuffle parameter"
	ErrInvalidCountParam   = "Invalid count parameter"
	ErrInvalidID           = "Invalid id"
)

type DeckHandler struct {
	uc usecases.DeckUseCase
}

func NewDeckHandler(uc usecases.DeckUseCase) DeckHandler {
	return DeckHandler{uc: uc}
}

func (h DeckHandler) Routes(router *gin.Engine) {
	router.POST("/decks", h.Create)
	router.GET("/decks/:id", h.Open)
	router.PATCH("/decks/:id/draw", h.Draw)
}

func (h DeckHandler) Create(context *gin.Context) {
	shuffleParam := context.DefaultQuery(ShuffleParam, "false")
	shuffle, err := strconv.ParseBool(shuffleParam)
	if err != nil {
		context.JSON(http.StatusBadRequest, ErrInvalidShuffleParam)
		return
	}

	cardsParam := context.DefaultQuery(CardsParam, "")
	createRequest := dtos.CreateRequest{
		Cards:   strings.Split(cardsParam, ","),
		Shuffle: shuffle,
	}

	cards, err := createRequest.ToCards()
	if err != nil {
		HandleErr(context, err)
		return
	}

	deck, err := h.uc.Create(cards, createRequest.Shuffle)
	if err != nil {
		HandleErr(context, err)
		return
	}

	context.JSON(http.StatusOK, dtos.ToCreateResponse(deck))
}

func (h DeckHandler) Open(context *gin.Context) {
	id := context.Params.ByName(IDParam)
	if !isValidUUID(id) {
		context.JSON(http.StatusBadRequest, ErrInvalidID)
		return
	}

	deck, err := h.uc.Get(id)
	if err != nil {
		HandleErr(context, err)
		return
	}

	context.JSON(http.StatusOK, dtos.ToOpenResponse(deck))
}

func (h DeckHandler) Draw(context *gin.Context) {
	countParam := context.DefaultQuery(CountParam, "1")
	count, err := strconv.Atoi(countParam)
	if err != nil || count < 1 || count > 52 {
		context.JSON(http.StatusBadRequest, ErrInvalidCountParam)
		return
	}

	id := context.Params.ByName(IDParam)

	drawnCards, err := h.uc.Draw(id, count)
	if err != nil {
		HandleErr(context, err)
		return
	}

	context.JSON(http.StatusOK, dtos.ToDrawResponse(drawnCards))
}

func isValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
