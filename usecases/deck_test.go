package usecases

import (
	"errors"
	"testing"

	internalerrors "deck-api/errors"
	"deck-api/models"
	"deck-api/test/mocks"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	saveFunc    = "Save"
	findOneFunc = "FindOne"
	updateFunc  = "Update"
)

func TestDeckUseCase(t *testing.T) {
	t.Run("Create a new deck successfully", func(t *testing.T) {
		deckRepository := new(mocks.DeckRepository)
		uc := NewDeckUseCase(deckRepository)

		cards := []models.Card{{Code: "2S"}, {Code: "9D"}, {Code: "KD"}}

		deckID := uuid.New()
		deck := newFakeDeck(deckID)

		deckRepository.On(saveFunc, mock.Anything).Return(deck, nil)

		deck, err := uc.Create(cards, false)

		assert.NoError(t, err)
		assert.Equal(t, deckID, deck.ID)
		deckRepository.AssertNumberOfCalls(t, saveFunc, 1)
		deckRepository.AssertNotCalled(t, updateFunc)
		deckRepository.AssertNotCalled(t, findOneFunc)
	})

	t.Run("Create a new deck with error", func(t *testing.T) {
		deckRepository := new(mocks.DeckRepository)
		uc := NewDeckUseCase(deckRepository)

		cards := []models.Card{{Code: "invalid"}}

		_, err := uc.Create(cards, false)

		assert.Error(t, err)
		assert.IsType(t, internalerrors.ErrInvalidEntry{}, err)
		deckRepository.AssertNotCalled(t, saveFunc)
		deckRepository.AssertNotCalled(t, updateFunc)
		deckRepository.AssertNotCalled(t, findOneFunc)
	})

	t.Run("Get a deck successfully", func(t *testing.T) {
		deckRepository := new(mocks.DeckRepository)
		uc := NewDeckUseCase(deckRepository)

		deckID := uuid.New()
		deckIDStr := deckID.String()
		deck := newFakeDeck(deckID)

		deckRepository.On(findOneFunc, deckIDStr).Return(deck, nil)

		foundDeck, err := uc.Get(deckIDStr)

		assert.NoError(t, err)
		assert.NotNil(t, foundDeck)
		deckRepository.AssertNumberOfCalls(t, findOneFunc, 1)
		deckRepository.AssertNotCalled(t, saveFunc)
		deckRepository.AssertNotCalled(t, updateFunc)
	})

	t.Run("Get a deck with error", func(t *testing.T) {
		deckRepository := new(mocks.DeckRepository)
		uc := NewDeckUseCase(deckRepository)

		deckIDStr := uuid.NewString()

		deckRepository.On(findOneFunc, deckIDStr).Return(models.Deck{}, internalerrors.ErrNotFound{})

		_, err := uc.Get(deckIDStr)

		assert.Error(t, err)
		assert.IsType(t, internalerrors.ErrNotFound{}, err)
		deckRepository.AssertNumberOfCalls(t, findOneFunc, 1)
	})

	t.Run("Draw cards from a deck successfully", func(t *testing.T) {
		deckRepository := new(mocks.DeckRepository)
		uc := NewDeckUseCase(deckRepository)

		deckID := uuid.New()
		deckIDStr := deckID.String()
		deck := newFakeDeck(deckID)

		count := deck.GetRemaining() - 1

		deckRepository.On(findOneFunc, deckIDStr).Return(deck, nil)

		deck.Cards = deck.Cards[count:]
		deckRepository.On(updateFunc, deck).Return(deck, nil)

		drawnCards, err := uc.Draw(deckIDStr, count)

		assert.NoError(t, err)
		assert.Equal(t, count, len(drawnCards))
		deckRepository.AssertNumberOfCalls(t, findOneFunc, 1)
		deckRepository.AssertNumberOfCalls(t, updateFunc, 1)
		deckRepository.AssertNotCalled(t, saveFunc)
	})

	t.Run("Draw cards from a deck with error because the deck was not found", func(t *testing.T) {
		deckRepository := new(mocks.DeckRepository)
		uc := NewDeckUseCase(deckRepository)

		deckIDStr := uuid.NewString()

		deckRepository.On(findOneFunc, deckIDStr).Return(models.Deck{}, internalerrors.ErrNotFound{})

		drawnCards, err := uc.Draw(deckIDStr, 1)

		assert.Error(t, err)
		assert.IsType(t, internalerrors.ErrNotFound{}, err)
		assert.Empty(t, drawnCards)
		deckRepository.AssertNumberOfCalls(t, findOneFunc, 1)
		deckRepository.AssertNotCalled(t, updateFunc)
		deckRepository.AssertNotCalled(t, saveFunc)
	})

	t.Run("Draw cards from a deck with error because the deck does not have sufficient cards to be drawn", func(t *testing.T) {
		deckRepository := new(mocks.DeckRepository)
		uc := NewDeckUseCase(deckRepository)

		deckID := uuid.New()
		deckIDStr := deckID.String()
		deck := newFakeDeck(deckID)

		count := deck.GetRemaining() + 1

		deckRepository.On(findOneFunc, deckIDStr).Return(deck, nil)

		drawnCards, err := uc.Draw(deckIDStr, count)

		assert.Error(t, err)
		assert.IsType(t, internalerrors.ErrInsufficientResources{}, err)
		assert.Empty(t, drawnCards)
		deckRepository.AssertNumberOfCalls(t, findOneFunc, 1)
		deckRepository.AssertNotCalled(t, updateFunc)
		deckRepository.AssertNotCalled(t, saveFunc)
	})

	t.Run("Draw cards from a deck with error because the update has failed", func(t *testing.T) {
		deckRepository := new(mocks.DeckRepository)
		uc := NewDeckUseCase(deckRepository)

		deckID := uuid.New()
		deckIDStr := deckID.String()
		deck := newFakeDeck(deckID)

		count := deck.GetRemaining() - 1

		deckRepository.On(findOneFunc, deckIDStr).Return(deck, nil)

		deck.Cards = deck.Cards[count:]
		deckRepository.On(updateFunc, deck).Return(deck, errors.New("internal error"))

		drawnCards, err := uc.Draw(deckIDStr, count)

		assert.Error(t, err)
		assert.Empty(t, drawnCards)
		deckRepository.AssertNumberOfCalls(t, findOneFunc, 1)
		deckRepository.AssertNumberOfCalls(t, updateFunc, 1)
		deckRepository.AssertNotCalled(t, saveFunc)
	})
}

// To improve this, maybe a lib like faker(https://github.com/bxcodec/faker) could be used.
func newFakeDeck(id uuid.UUID) models.Deck {
	return models.Deck{
		ID:    id,
		Cards: []string{"2S", "9D", "KD"},
	}
}
