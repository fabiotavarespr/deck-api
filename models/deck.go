package models

import (
	"math/rand"
	"time"

	internalerrors "deck-api/errors"

	"go.uber.org/zap"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Deck struct {
	ID         uuid.UUID
	IsShuffled bool
	Cards      pq.StringArray `gorm:"type:varchar(2)[]"`
}

func (d *Deck) GetRemaining() int {
	return len(d.Cards)
}

func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
	d.IsShuffled = true
}

func NewDeck(cards []Card) (Deck, error) {
	if len(cards) == 0 {
		return newStandardDeck(), nil
	}
	return newCustomDeck(cards)
}

func newStandardDeck() Deck {
	var cards []string
	for _, suit := range suitsSequence {
		for _, value := range cardsSequence {
			code := value + suit
			cards = append(cards, code)
		}
	}
	return Deck{
		ID:    uuid.New(),
		Cards: cards,
	}
}

func newCustomDeck(cards []Card) (Deck, error) {
	var codes []string
	for _, card := range cards {
		if err := card.Validate(); err != nil {
			zap.S().Errorf("An invalid card was found: %v", card)
			return Deck{}, internalerrors.ErrInvalidEntry{Message: "The set of cards contains one or more invalid entries"}
		}
		codes = append(codes, card.Code)
	}
	return Deck{
		ID:    uuid.New(),
		Cards: codes,
	}, nil
}
