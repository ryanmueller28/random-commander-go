package models

type Deck struct {
	Commanders []*Card
	Cards      []*Card
}

func (deck *Deck) AddCommander(card *Card) {
	deck.Commanders = append(deck.Commanders, card)
}

func (deck *Deck) AddCard(card *Card) {
	deck.Cards = append(deck.Cards, card)
}
