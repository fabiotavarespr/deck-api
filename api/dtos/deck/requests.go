package dtos

import (
	internalerrors "deck-api/errors"
	"deck-api/models"
)

// CreateRequest represents the entries to crate a new Deck.
type CreateRequest struct {
	Cards   []string
	Shuffle bool
}

func (c CreateRequest) ToCards() ([]models.Card, error) {
	var cards []models.Card
	for _, code := range c.Cards {
		if code == "" {
			continue
		}
		if err := c.validateCodeLength(code); err != nil {
			return []models.Card{}, err
		}
		cards = append(cards, models.Card{Code: code})
	}
	return cards, nil
}

func (c CreateRequest) validateCodeLength(code string) error {
	codeLength := len(code)
	if codeLength != 2 {
		return internalerrors.ErrInvalidEntry{Message: "The set of cards contains one or more invalid entries"}
	}
	return nil
}
