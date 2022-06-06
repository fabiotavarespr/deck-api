package repositories

import (
	internalerrors "deck-api/errors"
	"deck-api/models"

	"go.uber.org/zap"

	"github.com/jinzhu/gorm"
)

type DeckRepository interface {
	Save(deck models.Deck) (models.Deck, error)
	Update(deck models.Deck) (models.Deck, error)
	FindOne(id string) (models.Deck, error)
}

type deckRepository struct {
	conn *gorm.DB
}

func NewDeckRepository(conn *gorm.DB) DeckRepository {
	return deckRepository{conn: conn}
}

func (r deckRepository) Save(deck models.Deck) (models.Deck, error) {
	if err := r.conn.Create(&deck).Error; err != nil {
		zap.S().Errorf("Something went wrong to save a new deck: %s", err.Error())
		return models.Deck{}, err
	}
	return deck, nil
}

func (r deckRepository) Update(deck models.Deck) (models.Deck, error) {
	if err := r.conn.Save(&deck).Error; err != nil {
		zap.S().Errorf("Something went wrong to update deck: %s", err.Error())
		return models.Deck{}, err
	}
	return deck, nil
}

func (r deckRepository) FindOne(id string) (models.Deck, error) {
	deck := models.Deck{}
	result := r.conn.First(&deck, "id = ?", id)

	if result.RecordNotFound() {
		zap.S().Errorf("Deck with id %s was not found", id)
		return deck, internalerrors.ErrNotFound{Message: "Deck not found"}
	}

	err := result.Error
	if err != nil {
		zap.S().Errorf("Something went wrong to find Deck with id %s: %s", id, err.Error())
		return deck, err
	}

	return deck, nil
}
