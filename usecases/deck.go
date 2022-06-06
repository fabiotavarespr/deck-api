package usecases

import (
	internalerrors "deck-api/errors"
	"deck-api/models"
	"deck-api/repositories"

	"go.uber.org/zap"
)

type DeckUseCase interface {
	Create(cards []models.Card, isShuffled bool) (models.Deck, error)
	Get(id string) (models.Deck, error)
	Draw(id string, count int) ([]string, error)
}

type deckUseCase struct {
	repository repositories.DeckRepository
}

func NewDeckUseCase(repository repositories.DeckRepository) DeckUseCase {
	return deckUseCase{repository: repository}
}

func (uc deckUseCase) Create(cards []models.Card, isShuffled bool) (models.Deck, error) {
	deck, err := models.NewDeck(cards)
	if err != nil {
		return models.Deck{}, err
	}

	if isShuffled {
		deck.Shuffle()
	}

	return uc.repository.Save(deck)
}

func (uc deckUseCase) Get(id string) (models.Deck, error) {
	return uc.repository.FindOne(id)
}

func (uc deckUseCase) Draw(id string, count int) ([]string, error) {
	deck, err := uc.Get(id)
	if err != nil {
		return nil, err
	}

	remaining := deck.GetRemaining()
	if deck.GetRemaining() < count {
		zap.S().Errorf("There are / is only %d card(s) available to be drawn in the deck with id %s", remaining, id)
		return nil, internalerrors.ErrInsufficientResources{Message: "There are no sufficient cards to be drawn in the deck"}
	}

	drawnCards := deck.Cards[:count]
	deck.Cards = deck.Cards[count:]

	if _, err = uc.repository.Update(deck); err != nil {
		return nil, err
	}

	return drawnCards, nil
}
