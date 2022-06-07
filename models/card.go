package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

const (
	Ace      = "A"
	Two      = "2"
	Three    = "3"
	Four     = "4"
	Five     = "5"
	Six      = "6"
	Seven    = "7"
	Eight    = "8"
	Nine     = "9"
	Ten      = "0"
	Jack     = "J"
	Queen    = "Q"
	King     = "K"
	Spades   = "S"
	Clubs    = "C"
	Diamonds = "D"
	Hearts   = "H"
)

var (
	valuesSequence = []string{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}
	suitsSequence  = []string{Spades, Diamonds, Clubs, Hearts}
)

type Card struct {
	Code string
}

func (c Card) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Code, validation.Required, validation.Length(2, 2), validation.In(
			"AS", "2S", "3S", "4S", "5S", "6S", "7S", "8S", "9S", "0S", "JS", "QS", "KS",
			"AD", "2D", "3D", "4D", "5D", "6D", "7D", "8D", "9D", "0D", "JD", "QD", "KD",
			"AC", "2C", "3C", "4C", "5C", "6C", "7C", "8C", "9C", "0C", "JC", "QC", "KC",
			"AH", "2H", "3H", "4H", "5H", "6H", "7H", "8H", "9H", "0H", "JH", "QH", "KH",
		)),
	)
}
