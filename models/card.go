package models

import "random-commander/api"

type Card struct {
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	OracleText    string   `json:"oracle_text"`
	ColorIdentity []string `json:"color_identity"`
	Colors        []string `json:"colors"`
	TypeLine      string   `json:"type_line"`
}

func NewCardFromScryfall(scryfallCard *api.ScryfallCard) *Card {
	return &Card{
		Name:          scryfallCard.Name,
		OracleText:    scryfallCard.OracleText,
		ColorIdentity: scryfallCard.ColorIdentity,
		Colors:        scryfallCard.Colors,
		TypeLine:      scryfallCard.TypeLine,
	}
}
