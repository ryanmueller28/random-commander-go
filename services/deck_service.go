package services

import (
	"fmt"
	"random-commander/api"
	"random-commander/models"
	"strings"
)

type DeckService struct{}

func (s *DeckService) NewDeck() *models.Deck {
	var commanders []*models.Card

	commander := s.GetRandomCommander()

	colorIdentity := commander.ColorIdentity

	commanders = append(commanders, commander)

	secondMechanics := s.DetermineMultipleCommanderMechanic(commander)
	cmdr2 := s.FetchSecondCommander(secondMechanics)
	if cmdr2 != nil {
		commanders = append(commanders, cmdr2)
		colorIdentity = append(colorIdentity, cmdr2.ColorIdentity...)
	}
	deck := &models.Deck{
		Commanders: commanders,
		Cards:      []*models.Card{},
	}

	removalSpells := FetchRemovalSpells(10, strings.Join(colorIdentity, ""))
	boardWipes := FetchBoardWipes(5, strings.Join(colorIdentity, ""))
	cardAdvantageSpells := FetchCardAdvantageSpells(10, strings.Join(colorIdentity, ""))
	rampSpells := FetchRampSpells(5, strings.Join(colorIdentity, ""))

	deck.Cards = append(deck.Cards, removalSpells...)
	deck.Cards = append(deck.Cards, boardWipes...)
	deck.Cards = append(deck.Cards, cardAdvantageSpells...)
	deck.Cards = append(deck.Cards, rampSpells...)
	return deck
}

func (s *DeckService) GetRandomCommander() *models.Card {
	params := make(map[string]string)
	params["is"] = "commander"
	//params["game"] = "paper"
	sfCards, err := api.FetchCard(params, "random")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	card := (*sfCards)[0]
	return models.NewCardFromScryfall(&card)
}

func (s *DeckService) DetermineMultipleCommanderMechanic(commander *models.Card) string {
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

func (s *DeckService) FetchSecondCommander(secondCommanderMechanic string) *models.Card {
	if secondCommanderMechanic == "No multiple commanders" {
		return nil
	}
	params := make(map[string]string)
	//params["game"] = "paper"

	if secondCommanderMechanic == "Choose a Background" {
		params["oracle"] = "Choose a Background"
	} else if secondCommanderMechanic == "Doctor's Companion" {
		params["oracle"] = "Doctor's Companion"
	} else if secondCommanderMechanic == "Partner With" {
		fmt.Println("NOT YET IMPLEMENTED")
	} else if secondCommanderMechanic == "Partner" {
		params["oracle"] = "Partner"
	} else if secondCommanderMechanic == "Friends Forever" {
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
	return models.NewCardFromScryfall(&card)
}

func FetchBoardWipes(count int, colorIdentity string) []*models.Card {
	return fetchCoreCategory(count, colorIdentity, "boardwipe")
}

func FetchCardAdvantageSpells(count int, colorIdentity string) []*models.Card {
	return fetchCoreCategory(count, colorIdentity, "draw")
}

func FetchRemovalSpells(count int, colorIdentity string) []*models.Card {
	return fetchCoreCategory(count, colorIdentity, "removal")
}

func FetchRampSpells(count int, colorIdentity string) []*models.Card {
	return fetchCoreCategory(count, colorIdentity, "ramp")
}

func fetchCoreCategory(count int, colorIdentity string, cardFunction string) []*models.Card {
	params := make(map[string]string)
	params["function"] = cardFunction
	params["commander"] = colorIdentity
	cards := []*models.Card{}

	for i := 0; i < count; i++ {
		sfCards, err := api.FetchCard(params, "random")
		if err != nil {
			fmt.Println(err)
			return nil
		}
		card := (*sfCards)[0]
		cards = append(cards, models.NewCardFromScryfall(&card))
	}

	return cards
}
