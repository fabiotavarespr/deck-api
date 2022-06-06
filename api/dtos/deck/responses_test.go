package dtos

import (
	"testing"

	"deck-api/models"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestDeckResponsesDtos(t *testing.T) {
	t.Run("convert models.Deck to CreateResponse successfully", func(t *testing.T) {
		deckID := uuid.New()
		deck := models.Deck{
			ID:         deckID,
			IsShuffled: false,
			Cards:      []string{"AS", "0H", "KH"},
		}

		createResponse := ToCreateResponse(deck)

		assert.NotNil(t, createResponse)
		assert.False(t, createResponse.Shuffled)
		assert.Equal(t, len(deck.Cards), createResponse.Remaining)
		assert.Equal(t, createResponse.ID, deckID.String())
	})

	t.Run("convert models.Deck to OpenResponse successfully", func(t *testing.T) {
		ahCode := "AH"
		zerosCode := "0S"
		ksCode := "KS"

		deckID := uuid.New()
		deck := models.Deck{
			ID:         deckID,
			IsShuffled: false,
			Cards:      []string{ahCode, zerosCode, ksCode},
		}

		openResponse := ToOpenResponse(deck)

		assert.NotNil(t, openResponse)
		assert.False(t, openResponse.Shuffled)
		assert.Equal(t, len(deck.Cards), openResponse.Remaining)
		assert.Equal(t, openResponse.ID, deckID.String())

		// Assert card with code AH.
		ahCardDto := openResponse.Cards[0]
		assert.Equal(t, Ace, ahCardDto.Value)
		assert.Equal(t, Hearts, ahCardDto.Suit)
		assert.Equal(t, ahCode, ahCardDto.Code)

		// Assert card with code 0S.
		zerohCardDto := openResponse.Cards[1]
		assert.Equal(t, Ten, zerohCardDto.Value)
		assert.Equal(t, Spades, zerohCardDto.Suit)
		assert.Equal(t, zerosCode, zerohCardDto.Code)

		// Assert card with code KS.
		ksCardDto := openResponse.Cards[2]
		assert.Equal(t, King, ksCardDto.Value)
		assert.Equal(t, Spades, ksCardDto.Suit)
		assert.Equal(t, ksCode, ksCardDto.Code)
	})

	t.Run("convert codes to DrawResponse successfully", func(t *testing.T) {
		qcCode := "QC"
		qdCode := "QD"
		jhCode := "JH"

		codes := []string{qcCode, qdCode, jhCode}

		drawResponse := ToDrawResponse([]string{"QC", "QD", "JH"})

		assert.NotNil(t, drawResponse)
		assert.Equal(t, len(codes), len(drawResponse.Cards))

		// Assert card with code QC.
		qcCardDto := drawResponse.Cards[0]
		assert.Equal(t, Queen, qcCardDto.Value)
		assert.Equal(t, Clubs, qcCardDto.Suit)
		assert.Equal(t, qcCode, qcCardDto.Code)

		// Assert card with code QD.
		qdCardDto := drawResponse.Cards[1]
		assert.Equal(t, Queen, qdCardDto.Value)
		assert.Equal(t, Diamonds, qdCardDto.Suit)
		assert.Equal(t, qdCode, qdCardDto.Code)

		// Assert card with code JH.
		jhCardDto := drawResponse.Cards[2]
		assert.Equal(t, Jack, jhCardDto.Value)
		assert.Equal(t, Hearts, jhCardDto.Suit)
		assert.Equal(t, jhCode, jhCardDto.Code)
	})
}
