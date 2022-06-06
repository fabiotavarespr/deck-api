package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCardModel(t *testing.T) {
	t.Run("Validate cards successfully", func(t *testing.T) {
		allPossibleCards := []string{
			"AS", "2S", "3S", "4S", "5S", "6S", "7S", "8S", "9S", "0S", "JS", "QS", "KS",
			"AD", "2D", "3D", "4D", "5D", "6D", "7D", "8D", "9D", "0D", "JD", "QD", "KD",
			"AC", "2C", "3C", "4C", "5C", "6C", "7C", "8C", "9C", "0C", "JC", "QC", "KC",
			"AH", "2H", "3H", "4H", "5H", "6H", "7H", "8H", "9H", "0H", "JH", "QH", "KH",
		}

		var cards []Card
		for _, code := range allPossibleCards {
			cards = append(cards, Card{Code: code})
		}

		for _, card := range cards {
			assert.NoError(t, card.Validate())
		}
	})

	t.Run("Validate cards with error because the cards are invalid", func(t *testing.T) {
		invalidCards := []string{
			"10S", "10D", "10C", "10H", "99", "invalid", "AS5", "61H", "(%", "#$", "9Hj",
		}

		var cards []Card
		for _, code := range invalidCards {
			cards = append(cards, Card{Code: code})
		}

		for _, card := range cards {
			assert.Error(t, card.Validate())
		}
	})
}
