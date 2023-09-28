package graphic

import (
	"fmt"
	"game/utils"
	"game/values"
	"github.com/fatih/color"
)

func UpdateBalade() {
	utils.ClearTerminal()
	color.Yellow("------------------------------- Pokemon -------------------------------")
	switch values.MenuIndex {
	case 1:
		fmt.Println("-----------------------------")
		fmt.Println("|                           |")
		color.Green("> Dans les hautes herbes    |")
		fmt.Println("| Arène                     |")
		fmt.Println("| Quitter                   |")
		fmt.Println("|                           |")
		fmt.Println("-----------------------------")
		break
	case 2:
		fmt.Println("-----------------------------")
		fmt.Println("|                           |")
		fmt.Println("| Dans les hautes herbes    |")
		color.Green("> Arène                     |")
		fmt.Println("| Quitter                   |")
		fmt.Println("|                           |")
		fmt.Println("-----------------------------")
		break
	case 3:
		fmt.Println("-----------------------------")
		fmt.Println("|                           |")
		fmt.Println("| Dans les hautes herbes    |")
		fmt.Println("| Arène                     |")
		color.Green("> Quitter                   |")
		fmt.Println("|                           |")
		fmt.Println("-----------------------------")
		break
	}
}

func ActionBalade() {
	switch values.MenuIndex {
	case 1:
		values.MenuIndex = 1
		values.CurrentPage = "balade_pokemon"
		SelectPokemonBalade()
		break
	case 2:
		values.MenuIndex = 1
		values.CurrentPage = "select_arene"
		SelectArene()
		break
	case 3:
		values.MenuIndex = 1
		values.CurrentPage = "menu"
		UpdateMenu()
		break
	}
}

func SelectArene() {
	utils.ClearTerminal()
	color.Yellow("------------------------------- Pokemon -------------------------------")
	for i, maitre := range values.Maitres {
		if i+1 == values.MenuIndex {
			color.Green("> %s", maitre.Arene)
			color.Blue("\t[%s]", maitre.Nom)
			color.Yellow("	Pokemon : [%s]", maitre.Pokemon.Nom)
			color.Magenta("	Vie : [%d/%d]", maitre.Pokemon.Vie, maitre.Pokemon.MaxVie)
		} else {
			fmt.Println(maitre.Arene)
		}
	}
	if values.MenuIndex > len(values.Maitres) {
		color.Green("> Quitter")
	} else {
		fmt.Println("Quitter")
	}
}

func ActionSelectArene() {
	if values.MenuIndex > len(values.Maitres) {
		values.MenuIndex = 1
		values.CurrentPage = "balade"
		UpdateBalade()
	} else {
		values.MaitrePlaying = &values.Maitres[values.MenuIndex-1]
		values.EnemyBalade = &values.MaitrePlaying.Pokemon
		values.MenuIndex = 1
		values.CurrentPage = "select_arene_pokemon"
		SelectPokemonArene()
	}
}

func SelectPokemonArene() {
	utils.ClearTerminal()
	color.Yellow("------------------------------- Pokemon -------------------------------")
	for i, pokemon := range values.MainCharacter.Pokemons {
		if i+1 == values.MenuIndex {
			color.Green("> %s", pokemon.GetName())
			color.Blue("\t[%s]", pokemon.GetDescription())
			color.Yellow("	Vie : [%d/%d]", pokemon.GetLife(), pokemon.GetMaxLife())
			color.Magenta("	Classe : [%s]", pokemon.GetClass())
			color.Cyan("\t[%d] attaque(s)", len(pokemon.GetAttacks()))
		} else {
			fmt.Println(pokemon.GetName())
		}
	}
	if values.MenuIndex > len(values.MainCharacter.Pokemons) {
		color.Green("> Quitter")
	} else {
		fmt.Println("Quitter")
	}
}

func ConfirmSelectArene() {
	if values.MenuIndex > len(values.MainCharacter.Pokemons) {
		values.MenuIndex = 1
		values.CurrentPage = "balade"
		UpdateBalade()
	} else {
		if values.MainCharacter.Pokemons[values.MenuIndex-1].Vie < 1 {
			utils.ClearTerminal()
			utils.Writeanim("Ce pokemon n'a plus de PV", 10)
			fmt.Scanln()
			SelectPokemonArene()
		} else {
			values.PokemonBalade = &values.MainCharacter.Pokemons[values.MenuIndex-1]
			values.CurrentPage = "announce_arene"
			AnnounceArene()
		}
	}
}

func SelectPokemonBalade() {
	utils.ClearTerminal()
	color.Yellow("------------------------------- Pokemon -------------------------------")
	for i, pokemon := range values.MainCharacter.Pokemons {
		if i+1 == values.MenuIndex {
			color.Green("> %s", pokemon.GetName())
			color.Blue("\t[%s]", pokemon.GetDescription())
			color.Yellow("	Vie : [%d/%d]", pokemon.GetLife(), pokemon.GetMaxLife())
			color.Magenta("	Classe : [%s]", pokemon.GetClass())
			color.Cyan("\t[%d] attaque(s)", len(pokemon.GetAttacks()))
		} else {
			fmt.Println(pokemon.GetName())
		}
	}
	if values.MenuIndex > len(values.MainCharacter.Pokemons) {
		color.Green("> Quitter")
	} else {
		fmt.Println("Quitter")
	}
}

func ConfirmSelectBalade() {
	if values.MenuIndex > len(values.MainCharacter.Pokemons) {
		values.MenuIndex = 1
		values.CurrentPage = "balade"
		UpdateBalade()
	} else {
		if values.MainCharacter.Pokemons[values.MenuIndex-1].Vie < 1 {
			utils.ClearTerminal()
			utils.Writeanim("Ce pokemon n'a plus de PV", 10)
			fmt.Scanln()
			SelectPokemonBalade()
		} else {
			values.PokemonBalade = &values.MainCharacter.Pokemons[values.MenuIndex-1]
			values.CurrentPage = "look_enemy"
			LookForEnemy()
		}
	}
}
