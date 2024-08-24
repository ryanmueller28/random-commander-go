package models

import (
	"fmt"
	"random-commander/api"
	"strings"
)

type Deck struct {
	Commanders []*Card
	Cards      []*Card
}

func NewDeck() *Deck {
	var commanders []*Card

	commander := getRandomCommander()

	commanders = append(commanders, commander)

	secondMechanics := determineMultipleCommanderMechanic(commander)
	cmdr2 := fetchSecondCommander(secondMechanics)
	if cmdr2 != nil {
		commanders = append(commanders, cmdr2)
	}
	return &Deck{
		Commanders: commanders,
		Cards:      []*Card{},
	}
}

func (deck *Deck) AddCommander(card *Card) {
	deck.Commanders = append(deck.Commanders, card)
}

func (deck *Deck) AddCard(card *Card) {
	deck.Cards = append(deck.Cards, card)
}

func getRandomCommander() *Card {
	params := make(map[string]string)
	params["is"] = "commander"
	params["game"] = "paper"
	sfCards, err := api.FetchCard(params, "random")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	card := (*sfCards)[0]
	return NewCardFromScryfall(&card)
}

func determineMultipleCommanderMechanic(commander *Card) string {
	fmt.Println(commander)
	if strings.Contains(commander.TypeLine, "Background") {
		return "Choose a Background"
	} else if strings.Contains(commander.TypeLine, "Time Lord Doctor") {
		return "Doctor's Companion"
	}

	if strings.Contains(commander.OracleText, "Partner") {
		return "Partner"
	} else if strings.Contains(commander.OracleText, "Partner With") {

	} else if strings.Contains(commander.OracleText, "Doctor's Companion") {
		return "Time Lord Doctor"
	} else if strings.Contains(commander.OracleText, "Choose a Background") {
		return "Background"
	} else if strings.Contains(commander.OracleText, "Friends forever") {
		return "Friends Forever"
	}

	return "No multiple commanders"
}

func fetchSecondCommander(secondCommanderMechanic string) *Card {
	if secondCommanderMechanic == "No multiple commanders" {
		return nil
	}
	params := make(map[string]string)
	params["game"] = "paper"

	if secondCommanderMechanic == "Choose a Background" {
		params["oracle"] = "Choose a Background"
	} else if secondCommanderMechanic == "Doctor's Companion" {
		params["oracle"] = "Doctor's Companion"
	} else if secondCommanderMechanic == "Partner With" {
		fmt.Println("NOT YET IMPLEMENTED")
	} else if secondCommanderMechanic == "Partner" {
		params["oracle"] = "Partner"
	} else if secondCommanderMechanic == "Friends forever" {
		params["oracle"] = "Friends Forever"
	} else if secondCommanderMechanic == "Time Lord Doctor" {
		params["type"] = "Time Lord Doctor"
	} else if secondCommanderMechanic == "Background" {
		params["type"] = "Background"
	}

	sfCards, err := api.FetchCard(params, "random")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	card := (*sfCards)[0]
	return NewCardFromScryfall(&card)
}
