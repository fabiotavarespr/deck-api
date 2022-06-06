package dtos

import (
	"testing"

	internalerrors "deck-api/errors"

	"github.com/stretchr/testify/assert"
)

func TestDeckRequestsDtos(t *testing.T) {
	t.Run("covert to models.Cards successfully", func(t *testing.T) {
		codes := []string{"AS", "4D", "6C", "7C"}
		createRequest := CreateRequest{Cards: codes}

		cards, err := createRequest.ToCards()

		assert.NoError(t, err)
		assert.Equal(t, len(codes), len(cards))
	})

	t.Run("covert to models.Cards with error because the code 'A' is invalid", func(t *testing.T) {
		codes := []string{"A", "4D", "6C", "7C"}
		createRequest := CreateRequest{Cards: codes}

		cards, err := createRequest.ToCards()

		assert.Error(t, err)
		assert.IsType(t, internalerrors.ErrInvalidEntry{}, err)
		assert.Empty(t, cards)
	})

	t.Run("covert to models.Cards with error because the code 'AS1' is invalid", func(t *testing.T) {
		codes := []string{"AS1", "4D", "6C", "7C"}
		createRequest := CreateRequest{Cards: codes}

		cards, err := createRequest.ToCards()

		assert.Error(t, err)
		assert.IsType(t, internalerrors.ErrInvalidEntry{}, err)
		assert.Empty(t, cards)
	})
}
