package dtos

import "deck-api/models"

type (
	Suit  string
	Value string
)

// Possible CardDto suits.
const (
	Spades   Suit = "SPADES"
	Diamonds Suit = "DIAMONDS"
	Clubs    Suit = "CLUBS"
	Hearts   Suit = "HEARTS"
)

// Possible CardDto values.
const (
	Ace   Value = "ACE"
	Two   Value = "TWO"
	Three Value = "THREE"
	Four  Value = "FOUR"
	Five  Value = "FIVE"
	Six   Value = "SIX"
	Seven Value = "SEVEN"
	Eight Value = "EIGHT"
	Nine  Value = "NINE"
	Ten   Value = "TEN"
	Jack  Value = "JACK"
	Queen Value = "QUEEN"
	King  Value = "KING"
)

var suitDictionary = map[string]Suit{
	models.Spades:   Spades,
	models.Diamonds: Diamonds,
	models.Clubs:    Clubs,
	models.Hearts:   Hearts,
}

var valueDictionary = map[string]Value{
	models.Ace:   Ace,
	models.Two:   Two,
	models.Three: Three,
	models.Four:  Four,
	models.Five:  Five,
	models.Six:   Six,
	models.Seven: Seven,
	models.Eight: Eight,
	models.Nine:  Nine,
	models.Ten:   Ten,
	models.Jack:  Jack,
	models.Queen: Queen,
	models.King:  King,
}

type CardDto struct {
	Value Value  `json:"value"`
	Suit  Suit   `json:"suit"`
	Code  string `json:"code"`
}

type CreateResponse struct {
	ID        string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
}

type OpenResponse struct {
	CreateResponse
	Cards []CardDto `json:"cards"`
}

type DrawResponse struct {
	Cards []CardDto `json:"cards"`
}

func ToCreateResponse(deck models.Deck) CreateResponse {
	createResponse := CreateResponse{}
	createResponse.ID = deck.ID.String()
	createResponse.Shuffled = deck.IsShuffled
	createResponse.Remaining = deck.GetRemaining()
	return createResponse
}

func ToOpenResponse(deck models.Deck) OpenResponse {
	openResponse := OpenResponse{}
	openResponse.Cards = ToCardDtos(deck.Cards)
	openResponse.ID = deck.ID.String()
	openResponse.Shuffled = deck.IsShuffled
	openResponse.Remaining = deck.GetRemaining()
	return openResponse
}

func ToDrawResponse(drawnCards []string) DrawResponse {
	drawResponse := DrawResponse{}
	drawResponse.Cards = ToCardDtos(drawnCards)
	return drawResponse
}

func ToCardDtos(codes []string) []CardDto {
	cardDtos := make([]CardDto, 0)
	for _, code := range codes {
		cardDto := CardDto{
			Value: valueDictionary[string(code[0])],
			Suit:  suitDictionary[string(code[1])],
			Code:  code,
		}
		cardDtos = append(cardDtos, cardDto)
	}
	return cardDtos
}
