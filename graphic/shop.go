package graphic

import (
	"fmt"
	"game/bdd"
	"game/utils"
	"game/values"
	"github.com/fatih/color"
)

func UpdateShop() {
	utils.ClearTerminal()
	color.Yellow("------------------------------- Pokemon -------------------------------")
	switch values.MenuIndex {
	case 1:
		fmt.Println("-----------------------------")
		fmt.Println("|                           |")
		color.Green("> Pokeball (10€)            |")
		fmt.Println("| Potion de vie             |")
		fmt.Println("| Potion de poison          |")
		fmt.Println("| Augmenter capacité        |")
		fmt.Println("| Pierre de nuit            |")
		fmt.Println("| Quitter                   |")
		fmt.Println("|                           |")
		fmt.Println("-----------------------------")
		break
	case 2:
		fmt.Println("-----------------------------")
		fmt.Println("|                           |")
		fmt.Println("| Pokeball                  |")
		color.Green("> Potion de vie (15€)       |")
		fmt.Println("| Potion de poison          |")
		fmt.Println("| Augmenter capacité        |")
		fmt.Println("| Pierre de nuit            |")
		fmt.Println("| Quitter                   |")
		fmt.Println("|                           |")
		fmt.Println("-----------------------------")
		break
	case 3:
		fmt.Println("-----------------------------")
		fmt.Println("|                           |")
		fmt.Println("| Pokeball                  |")
		fmt.Println("| Potion de vie             |")
		color.Green("> Potion de poison (25€)    |")
		fmt.Println("| Augmenter capacité        |")
		fmt.Println("| Pierre de nuit            |")
		fmt.Println("| Quitter                   |")
		fmt.Println("|                           |")
		fmt.Println("-----------------------------")
		break
	case 4:
		fmt.Println("-----------------------------")
		fmt.Println("|                           |")
		fmt.Println("| Pokeball                  |")
		fmt.Println("| Potion de vie             |")
		fmt.Println("| Potion de poison          |")
		color.Green("> Augmenter capacité (100€) |")
		fmt.Println("| Pierre de nuit            |")
		fmt.Println("| Quitter                   |")
		fmt.Println("|                           |")
		fmt.Println("-----------------------------")
		break
	case 5:
		fmt.Println("-----------------------------")
		fmt.Println("|                           |")
		fmt.Println("| Pokeball                  |")
		fmt.Println("| Potion de vie             |")
		fmt.Println("| Potion de poison          |")
		fmt.Println("| Augmenter capacité        |")
		color.Green("> Pierre de nuit (50€)      |")
		fmt.Println("| Quitter                   |")
		fmt.Println("|                           |")
		fmt.Println("-----------------------------")
		break
	case 6:
		fmt.Println("-----------------------------")
		fmt.Println("|                           |")
		fmt.Println("| Pokeball                  |")
		fmt.Println("| Potion de vie             |")
		fmt.Println("| Potion de poison          |")
		fmt.Println("| Augmenter capacité        |")
		fmt.Println("| Pierre de nuit            |")
		color.Green("> Quitter                   |")
		fmt.Println("|                           |")
		fmt.Println("-----------------------------")
		break
	}
	fmt.Println()
	color.Blue("Argent restant : %d", values.MainCharacter.Argent)
}

func ActionShop() {
	switch values.MenuIndex {
	case 1:
		if values.MainCharacter.Argent < 10 {
			utils.PlaySound(values.Sounds["error"])
		} else {
			values.MainCharacter.Argent -= 10
			bdd.Database.Set("character_argent", values.MainCharacter.Argent)
			values.MainCharacter.Pokeball++
			bdd.Database.Set("character_pokeball", values.MainCharacter.Pokeball)
			utils.PlaySound(values.Sounds["success"])
			UpdateShop()
		}
		break
	case 2:
		if values.MainCharacter.Argent < 15 {
			utils.PlaySound(values.Sounds["error"])
		} else {
			values.MainCharacter.Argent -= 15
			bdd.Database.Set("character_argent", values.MainCharacter.Argent)
			values.MainCharacter.PotionVie++
			bdd.Database.Set("character_potionvie", values.MainCharacter.PotionVie)
			utils.PlaySound(values.Sounds["success"])
			UpdateShop()
		}
		break
	case 3:
		if values.MainCharacter.Argent < 25 {
			utils.PlaySound(values.Sounds["error"])
		} else {
			values.MainCharacter.Argent -= 25
			bdd.Database.Set("character_argent", values.MainCharacter.Argent)
			values.MainCharacter.PotionPoison++
			bdd.Database.Set("character_potionpoison", values.MainCharacter.PotionPoison)
			utils.PlaySound(values.Sounds["success"])
			UpdateShop()
		}
		break
	case 4:
		if values.MainCharacter.Argent < 100 {
			utils.PlaySound(values.Sounds["error"])
		} else {
			values.MainCharacter.Argent -= 100
			bdd.Database.Set("character_argent", values.MainCharacter.Argent)
			values.MainCharacter.PokemonMax++
			bdd.Database.Set("character_pokemax", values.MainCharacter.PokemonMax)
			utils.PlaySound(values.Sounds["success"])
			UpdateShop()
		}
		break
	case 5:
		if values.MainCharacter.Argent < 50 {
			utils.PlaySound(values.Sounds["error"])
		} else {
			values.MainCharacter.Argent -= 50
			bdd.Database.Set("character_argent", values.MainCharacter.Argent)
			values.MainCharacter.PierreNuit++
			bdd.Database.Set("character_pierrenuit", values.MainCharacter.PierreNuit)
			utils.PlaySound(values.Sounds["success"])
			UpdateShop()
		}
		break
	case 6:
		values.MenuIndex = 1
		values.CurrentPage = "menu"
		UpdateMenu()
		break
	}
}
