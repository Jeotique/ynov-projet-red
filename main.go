package main

import (
	"fmt"
	"game/bdd"
	"game/graphic"
	"game/utils"
	"game/values"
	"github.com/mattn/go-tty"
	"log"
	"os"
)

func main() {
	values.Database = bdd.NewQuickDB("database.json")
	if values.Database.Get("pokemon_selected") != nil && values.Database.Get("pokemon_selected").(bool) {
		values.CurrentPage = "menu"
		graphic.UpdateMenu()
		StartListening()
	} else {
		graphic.SelectPokemon()
		fmt.Println("Séléctionnez votre pokémon de départ :")
		valid, input := utils.WaitForNumberInput()
		if !valid {
			println("Vous devez entrer des chiffres")
			os.Exit(0)
		} else {
			switch input {
			case 1:
				values.Database.Set("pokemon_selected", true)
				values.MainCharacter.Pokemons = append(values.MainCharacter.Pokemons, values.FindPokemon("Bulbizarre"))
				values.Database.Set("character", values.MainCharacter)
				break
			case 2:
				values.Database.Set("pokemon_selected", true)
				values.MainCharacter.Pokemons = append(values.MainCharacter.Pokemons, values.FindPokemon("Salamèche"))
				values.Database.Set("character", values.MainCharacter)
				break
			case 3:
				values.Database.Set("pokemon_selected", true)
				values.MainCharacter.Pokemons = append(values.MainCharacter.Pokemons, values.FindPokemon("Carapuce"))
				values.Database.Set("character", values.MainCharacter)
				break
			default:
				println("Pokémon invalide")
				os.Exit(0)
			}
		}
	}

}

func StartListening() {
	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()

	for {
		r, err := tty.ReadRune()
		if err != nil {
			log.Fatal(err)
		}
		if r == 65 || r == 66 { // 65 = fleche haut & 66 = fleche bas
			var maxIndex int
			var funcType func()
			switch values.CurrentPage {
			case "menu":
				maxIndex = 4
				funcType = graphic.UpdateMenu
				break
			}
			if r == 65 {
				if values.MenuIndex == 1 {
					values.MenuIndex = maxIndex
				} else {
					values.MenuIndex--
				}
			} else if r == 66 {
				if values.MenuIndex == maxIndex {
					values.MenuIndex = 1
				} else {
					values.MenuIndex++
				}
			}
			funcType()
		} else {
			if r == 13 { // 13 = touche entrée

			}
		}
	}
}
