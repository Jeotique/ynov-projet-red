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
		values.CurrentPage = "pokemon_info"
		values.PokemonInfo = &values.MainCharacter.Pokemons[values.MenuIndex-1]
		values.MenuIndex = 1
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
			fmt.Println("| Voir attaques         |")
			fmt.Println("| Quitter               |")
			fmt.Println("|                       |")
			fmt.Println("-------------------------")
			break
		case 2:
			fmt.Println("-------------------------")
			fmt.Println("|                       |")
			fmt.Println("| Relâcher              |")
			color.Green("> Voir attaques         |")
			fmt.Println("| Quitter               |")
			fmt.Println("|                       |")
			fmt.Println("-------------------------")
			break
		case 3:
			fmt.Println("-------------------------")
			fmt.Println("|                       |")
			fmt.Println("| Relâcher              |")
			fmt.Println("| Voir attaques         |")
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
			fmt.Println("| Voir attaques         |")
			fmt.Println("| Quitter               |")
			fmt.Println("|                       |")
			fmt.Println("-------------------------")
			break
		case 2:
			fmt.Println("-------------------------")
			fmt.Println("|                       |")
			fmt.Println("| Utiliser potion vie   |")
			color.Green("> Relâcher              |")
			fmt.Println("| Voir attaques         |")
			fmt.Println("| Quitter               |")
			fmt.Println("|                       |")
			fmt.Println("-------------------------")
			break
		case 3:
			fmt.Println("-------------------------")
			fmt.Println("|                       |")
			fmt.Println("| Utiliser potion vie   |")
			fmt.Println("| Relâcher              |")
			color.Green("> Voir attaques         |")
			fmt.Println("| Quitter               |")
			fmt.Println("|                       |")
			fmt.Println("-------------------------")
			break
		case 4:
			fmt.Println("-------------------------")
			fmt.Println("|                       |")
			fmt.Println("| Utiliser potion vie   |")
			fmt.Println("| Relâcher              |")
			fmt.Println("| Voir attaques         |")
			color.Green("> Quitter               |")
			fmt.Println("|                       |")
			fmt.Println("-------------------------")
			break
		}
	}
	data2 := [][]string{
		[]string{values.PokeArt[pokemon.GetName()]},
	}
	table2 := tablewriter.NewWriter(os.Stdout)
	for _, v := range data2 {
		table2.Append(v)
	}
	table2.Render()
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
				go utils.PlaySound(values.Sounds["error"])
			}
		}
		break
	case 2:
		if values.MainCharacter.PotionVie < 1 {
			values.MenuIndex = 1
			values.CurrentPage = "attacks"
			UpdatePokemonAttacks()
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
		if values.MainCharacter.PotionVie < 1 {
			values.MenuIndex = 1
			values.CurrentPage = "pokemons"
			UpdatePokemonsBag()
		} else {
			values.MenuIndex = 1
			values.CurrentPage = "attacks"
			UpdatePokemonAttacks()
		}
		break
	case 4:
		values.MenuIndex = 1
		values.CurrentPage = "pokemons"
		UpdatePokemonsBag()
		break
	}
}

func UpdatePokemonAttacks() {
	utils.ClearTerminal()
	color.Yellow("------------------------------- Pokemon -------------------------------")
	pokemon := values.PokemonInfo
	for i, attack := range pokemon.GetAttacks() {
		if i+1 == values.MenuIndex {
			color.Green("> %s", attack.Nom)
			color.Blue("     Classe : [%s]", attack.Classe)
			color.Yellow("     Dégat min/max : [%d/%d]", attack.DegatMin, attack.DegatMax)
			color.Magenta("     [%s]", attack.Description)
		} else {
			fmt.Println(attack.Nom)
		}
	}
	if values.MenuIndex > len(pokemon.GetAttacks()) {
		color.Green("> Quitter")
	} else {
		fmt.Println("Quitter")
	}
}

func ActionPokemonAttacks() {
	if values.MenuIndex > len(values.PokemonInfo.GetAttacks()) {
		values.MenuIndex = 1
		values.CurrentPage = "pokemon_info"
		DisplayPokemonInfo()
	} else {
		if values.MainCharacter.PierreNuit < 1 {
			utils.ClearTerminal()
			utils.Writeanim("Vous n'avez aucune pierre de nuit pour améliorer cette attaque", 10)
			fmt.Scanln()
			UpdatePokemonAttacks()
		} else {
			values.AttackInfo = &values.PokemonInfo.GetAttacks()[values.MenuIndex-1]
			values.MenuIndex = 1
			values.CurrentPage = "upgrade"
			AskForUpgrade()
		}
	}
}

func AskForUpgrade() {
	utils.ClearTerminal()
	color.Yellow("------------------------------- Pokemon -------------------------------")
	color.Yellow("Voulez-vous vraiment utiliser x1 pierre de nuit pour améliorer [%s] ?", values.AttackInfo.Nom)
	switch values.MenuIndex {
	case 1:
		fmt.Println("--------------------------")
		fmt.Println("|                        |")
		color.Green("> Améliorer              |")
		fmt.Println("| Annuler                |")
		fmt.Println("|                        |")
		fmt.Println("--------------------------")
		break
	case 2:
		fmt.Println("--------------------------")
		fmt.Println("|                        |")
		fmt.Println("| Améliorer              |")
		color.Green("> Annuler                |")
		fmt.Println("|                        |")
		fmt.Println("--------------------------")
		break
	}
}

func ActionAskUpgrade() {
	switch values.MenuIndex {
	case 1:
		values.MainCharacter.PierreNuit--
		bdd.Database.Set("character_pierrenuit", values.MainCharacter.PierreNuit)
		values.AttackInfo.DegatMin += 10
		values.AttackInfo.DegatMax += 10
		bdd.Database.SavePokemons(values.MainCharacter.Pokemons)
		utils.ClearTerminal()
		go utils.PlaySound(values.Sounds["upgrade"])
		utils.Writeanim("Félicitation ! Vous venez d'améliorer votre attaque !", 10)
		fmt.Scanln()
		values.MenuIndex = 1
		values.CurrentPage = "attacks"
		UpdatePokemonAttacks()
		break
	case 2:
		values.MenuIndex = 1
		values.CurrentPage = "attacks"
		UpdatePokemonAttacks()
		break
	}
}
