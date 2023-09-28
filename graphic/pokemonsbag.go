package graphic

import (
	"fmt"
	"game/bdd"
	"game/utils"
	"game/values"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
)

func UpdatePokemonsBag() {
	utils.ClearTerminal()
	color.Yellow("------------------------------- Pokemon -------------------------------")
	for i, pokemon := range values.MainCharacter.Pokemons {
		if i+1 == values.MenuIndex {
			color.Green("> %s", pokemon.GetName())
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

func ActionPokemonsBag() {
	if values.MenuIndex > len(values.MainCharacter.Pokemons) {
		values.MenuIndex = 1
		values.CurrentPage = "bag"
		UpdateBag()
	} else {
		values.MenuIndex = 1
		values.CurrentPage = "pokemon_info"
		values.PokemonInfo = &values.MainCharacter.Pokemons[values.MenuIndex-1]
		DisplayPokemonInfo()
	}
}

func DisplayPokemonInfo() {
	pokemon := values.PokemonInfo
	utils.ClearTerminal()
	color.Yellow("------------------------------- Pokemon -------------------------------")
	data := [][]string{
		[]string{"Nom", pokemon.GetName()},
		[]string{"Type", pokemon.GetClass()},
		[]string{"Description", pokemon.GetDescription()},
		[]string{"Vie", strconv.Itoa(pokemon.GetLife()) + "/" + strconv.Itoa(pokemon.GetMaxLife())},
	}
	table := tablewriter.NewWriter(os.Stdout)
	for _, v := range data {
		table.Append(v)
	}
	table.Render()

	if values.MainCharacter.PotionVie < 1 {
		switch values.MenuIndex {
		case 1:
			fmt.Println("-------------------------")
			fmt.Println("|                       |")
			color.Green("> Relâcher              |")
			fmt.Println("| Quitter               |")
			fmt.Println("|                       |")
			fmt.Println("-------------------------")
			break
		case 2:
			fmt.Println("-------------------------")
			fmt.Println("|                       |")
			fmt.Println("| Relâcher              |")
			color.Green("> Quitter               |")
			fmt.Println("|                       |")
			fmt.Println("-------------------------")
			break
		}
	} else {
		switch values.MenuIndex {
		case 1:
			fmt.Println("-------------------------")
			fmt.Println("|                       |")
			color.Green("> Utiliser potion vie   |")
			fmt.Println("| Relâcher              |")
			fmt.Println("| Quitter               |")
			fmt.Println("|                       |")
			fmt.Println("-------------------------")
			break
		case 2:
			fmt.Println("-------------------------")
			fmt.Println("|                       |")
			fmt.Println("| Utiliser potion vie   |")
			color.Green("> Relâcher              |")
			fmt.Println("| Quitter               |")
			fmt.Println("|                       |")
			fmt.Println("-------------------------")
			break
		case 3:
			fmt.Println("-------------------------")
			fmt.Println("|                       |")
			fmt.Println("| Utiliser potion vie   |")
			fmt.Println("| Relâcher              |")
			color.Green("> Quitter               |")
			fmt.Println("|                       |")
			fmt.Println("-------------------------")
			break
		}
	}
}

func ActionPokemonInfo() {
	switch values.MenuIndex {
	case 1:
		if values.MainCharacter.PotionVie < 1 {
			if len(values.MainCharacter.Pokemons) < 2 {
				utils.ClearTerminal()
				utils.Writeanim("Vous ne pouvez pas relâcher votre seul et unique pokemon", 10)
				fmt.Scanln()
				DisplayPokemonInfo()
			} else {
				var newPokemons []values.Pokemon
				for _, poke := range values.MainCharacter.Pokemons {
					if poke.GetName() == values.PokemonInfo.GetName() {
						continue
					} else {
						newPokemons = append(newPokemons, poke)
					}
				}
				values.MainCharacter.Pokemons = newPokemons
				bdd.Database.RemovePokemon(values.PokemonInfo.GetName())
				values.MenuIndex = 1
				values.CurrentPage = "pokemons"
				UpdatePokemonsBag()
			}
		} else {
			if values.PokemonInfo.GetLife() < values.PokemonInfo.GetMaxLife() {
				values.MainCharacter.PotionVie--
				values.PokemonInfo.Heal()
				bdd.Database.Set("character_potionvie", values.MainCharacter.PotionVie)
				utils.PlaySound(values.Sounds["heal"], 1)
				bdd.Database.SavePokemons(values.MainCharacter.Pokemons)
				DisplayPokemonInfo()
			} else {
				utils.PlaySound(values.Sounds["error"])
			}
		}
		break
	case 2:
		if values.MainCharacter.PotionVie < 1 {
			values.MenuIndex = 1
			values.CurrentPage = "pokemons"
			UpdatePokemonsBag()
		} else {
			if len(values.MainCharacter.Pokemons) < 2 {
				utils.ClearTerminal()
				utils.Writeanim("Vous ne pouvez pas relâcher votre seul et unique pokemon", 10)
				fmt.Scanln()
				DisplayPokemonInfo()
			} else {
				var newPokemons []values.Pokemon
				for _, poke := range values.MainCharacter.Pokemons {
					if poke.GetName() == values.PokemonInfo.GetName() {
						continue
					} else {
						newPokemons = append(newPokemons, poke)
					}
				}
				values.MainCharacter.Pokemons = newPokemons
				bdd.Database.RemovePokemon(values.PokemonInfo.GetName())
				values.MenuIndex = 1
				values.CurrentPage = "pokemons"
				UpdatePokemonsBag()
			}
		}
		break
	case 3:
		values.MenuIndex = 1
		values.CurrentPage = "pokemons"
		UpdatePokemonsBag()
		break
	}
}
