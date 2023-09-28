package values

import "game/bdd"

var MenuIndex = 1
var MainCharacter Personnage
var CurrentPage = "menu"
var Database *bdd.QuickDB

type MainGame struct {
	Game Progress `json:"progress"`
}

type Progress struct {
	PokemonChoice bool    `json:"prokemon_selected"`
	Pokemon       Pokemon `json:"pokemon"`
}
