package main

import (
	"fmt"
	"game/bdd"
	"game/graphic"
	"game/utils"
	"game/values"
	"github.com/briandowns/spinner"
	"github.com/mattn/go-tty"
	"log"
	"os"
	"time"
)

func main() {
	utils.ClearTerminal()
	bdd.Database = bdd.NewQuickDB("database.json")
	if bdd.Database.Get("pokemon_selected") != nil && bdd.Database.Get("pokemon_selected").(bool) {
		values.CurrentPage = "menu"
		graphic.UpdateMenu()
		SavedCharacterName := bdd.Database.Get("character_name")
		if SavedCharacterName != nil {
			values.MainCharacter.Nom = SavedCharacterName.(string)
		}
		bdd.Database.LoadPokemons()
		bdd.Database.LoadBadges()
		if bdd.Database.Get("character_pokemax") != nil {
			values.MainCharacter.PokemonMax = int(bdd.Database.Get("character_pokemax").(float64))
		}
		if bdd.Database.Get("character_argent") != nil {
			values.MainCharacter.Argent = int(bdd.Database.Get("character_argent").(float64))
		}
		if bdd.Database.Get("character_pokeball") != nil {
			values.MainCharacter.Pokeball = int(bdd.Database.Get("character_pokeball").(float64))
		}
		if bdd.Database.Get("character_potionvie") != nil {
			values.MainCharacter.PotionVie = int(bdd.Database.Get("character_potionvie").(float64))
		}
		if bdd.Database.Get("character_potionpoison") != nil {
			values.MainCharacter.PotionPoison = int(bdd.Database.Get("character_potionpoison").(float64))
		}
		if bdd.Database.Get("character_pierrenuit") != nil {
			values.MainCharacter.PierreNuit = int(bdd.Database.Get("character_pierrenuit").(float64))
		}
		if bdd.Database.Get("character_xp") != nil {
			values.MainCharacter.Xp = int(bdd.Database.Get("character_xp").(float64))
		}
		if bdd.Database.Get("character_level") != nil {
			values.MainCharacter.Level = int(bdd.Database.Get("character_level").(float64))
		}
		StartListening()
	} else {

		utils.Writeanim("                                  ,'\\\n    _.----.        ____         ,'  _\\   ___    ___     ____\n_,-'       `.     |    |  /`.   \\,-'    |   \\  /   |   |    \\  |`.\n\\      __    \\    '-.  | /   `.  ___    |    \\/    |   '-.   \\ |  |\n \\.    \\ \\   |  __  |  |/    ,','_  `.  |          | __  |    \\|  |\n   \\    \\/   /,' _`.|      ,' / / / /   |          ,' _`.|     |  |\n    \\     ,-'/  /   \\    ,'   | \\/ / ,`.|         /  /   \\  |     |\n     \\    \\ |   \\_/  |   `-.  \\    `'  /|  |    ||   \\_/  | |\\    |\n      \\    \\ \\      /       `-.`.___,-' |  |\\  /| \\      /  | |   |\n       \\    \\ `.__,'|  |`-._    `|      |__| \\/ |  `.__,'|  | |   |\n        \\_.-'       |__|    `-._ |              '-.|     '-.| |   |\n                                `'                            '-._|", 1)
		fmt.Println()
		s := spinner.New(spinner.CharSets[38], 100*time.Millisecond)
		s.FinalMSG = "Game loaded !\n"
		s.Color("bold")
		s.Start()
		time.Sleep(6000 * time.Millisecond)
		fmt.Println()
		s.Stop()

		graphic.SelectPokemon()
		go playMusic()
		fmt.Println("Séléctionnez votre pokémon de départ :")
		valid, input := utils.WaitForNumberInput()
		if !valid {
			println("Vous devez entrer des chiffres")
			os.Exit(0)
		} else {
			switch input {
			case 1:
				bdd.Database.Set("pokemon_selected", true)
				values.MainCharacter.Pokemons = append(values.MainCharacter.Pokemons, values.FindPokemon("Bulbizarre"))
				bdd.Database.AddPokemon("Bulbizarre")
				ChoseUrName()
				break
			case 2:
				bdd.Database.Set("pokemon_selected", true)
				values.MainCharacter.Pokemons = append(values.MainCharacter.Pokemons, values.FindPokemon("Salamèche"))
				bdd.Database.AddPokemon("Salamèche")
				ChoseUrName()
				break
			case 3:
				bdd.Database.Set("pokemon_selected", true)
				values.MainCharacter.Pokemons = append(values.MainCharacter.Pokemons, values.FindPokemon("Carapuce"))
				bdd.Database.AddPokemon("Carapuce")
				ChoseUrName()
				break
			default:
				println("Pokémon invalide")
				os.Exit(0)
			}
		}
	}

}

func playMusic() {
	go utils.PlaySound(values.Sounds["opening"], 1)
}

func ChoseUrName() {
	utils.ClearTerminal()
	fmt.Println(`  ,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,  
 /                                                                \ 
|    Quel sera le nom de votre chasseur ?                   	 |
|                                                                |
|    (Appuyez sur Entrée pour confirmer)                           |
|                                                                |
 \                                                                /
  \,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,/
`)
	input := utils.WaitForInput()
	values.MainCharacter.Nom = input
	values.MainCharacter.PokemonMax = 3
	values.MainCharacter.Argent = 60
	bdd.Database.Set("character_name", input)
	bdd.Database.Set("character_pokemax", values.MainCharacter.PokemonMax)
	bdd.Database.Set("character_argent", values.MainCharacter.Argent)
	bdd.Database.SaveBadges(values.MainCharacter.Badges)
	utils.ClearTerminal()
	utils.Writeanim("Bienvenue, "+input, 10)
	utils.Writeanim("\nHeureux de partir à l'aventure avec vous !", 5)
	utils.Writeanim("\nAppuyez sur Entrée pour continuer...", 10)
	fmt.Scanln()
	values.CurrentPage = "menu"
	graphic.UpdateMenu()
	StartListening()
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
				maxIndex = 5
				funcType = graphic.UpdateMenu
				break
			case "bag":
				maxIndex = 4
				funcType = graphic.UpdateBag
				break
			case "pokemons":
				maxIndex = len(values.MainCharacter.Pokemons) + 1
				funcType = graphic.UpdatePokemonsBag
				break
			case "shop":
				maxIndex = 6
				funcType = graphic.UpdateShop
				break
			case "pokemon_info":
				if values.MainCharacter.PotionVie < 1 {
					maxIndex = 3
				} else {
					maxIndex = 4
				}
				funcType = graphic.DisplayPokemonInfo
				break
			case "attacks":
				maxIndex = len(values.PokemonInfo.GetAttacks()) + 1
				funcType = graphic.UpdatePokemonAttacks
				break
			case "upgrade":
				maxIndex = 2
				funcType = graphic.AskForUpgrade
				break
			case "balade":
				maxIndex = 3
				funcType = graphic.UpdateBalade
				break
			case "balade_pokemon":
				maxIndex = len(values.MainCharacter.Pokemons) + 1
				funcType = graphic.SelectPokemonBalade
				break
			case "in_battle":
				if values.MainCharacter.PotionPoison < 1 {
					maxIndex = len(values.PokemonBalade.GetAttacks()) + 1
				} else {
					maxIndex = len(values.PokemonBalade.GetAttacks()) + 2
				}
				funcType = graphic.BattleBalade
				break
			case "select_arene":
				maxIndex = len(values.Maitres) + 1
				funcType = graphic.SelectArene
				break
			case "select_arene_pokemon":
				maxIndex = len(values.MainCharacter.Pokemons) + 1
				funcType = graphic.SelectPokemonArene
				break
			case "in_battle_arene":
				if values.MainCharacter.PotionPoison < 1 {
					maxIndex = len(values.PokemonBalade.GetAttacks())
				} else {
					maxIndex = len(values.PokemonBalade.GetAttacks()) + 1
				}
				funcType = graphic.BattleArene
				break
			case "badge":
				maxIndex = len(values.MainCharacter.Badges) + 1
				funcType = graphic.UpdateBadge
				break
			default:
				continue
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
			if values.CurrentPage != "in_battle" && values.CurrentPage != "in_battle_arene" {
				utils.PlaySound(values.Sounds["navigate"], 3)
			}
		} else {
			if r == 13 { // 13 = touche entrée
				if values.CurrentPage != "in_battle" && values.CurrentPage != "in_battle_arene" {
					utils.PlaySound(values.Sounds["click"], 3)
				}
				switch values.CurrentPage {
				case "menu":
					graphic.ActionMenuu()
					break
				case "bag":
					graphic.ActionBag()
					break
				case "pokemons":
					graphic.ActionPokemonsBag()
					break
				case "pokemon_info":
					graphic.ActionPokemonInfo()
					break
				case "inventory":
					values.CurrentPage = "bag"
					graphic.UpdateBag()
					break
				case "shop":
					graphic.ActionShop()
					break
				case "attacks":
					graphic.ActionPokemonAttacks()
					break
				case "upgrade":
					graphic.ActionAskUpgrade()
					break
				case "balade":
					graphic.ActionBalade()
					break
				case "balade_pokemon":
					graphic.ConfirmSelectBalade()
					break
				case "in_battle":
					graphic.ActionBattle()
					break
				case "select_arene":
					graphic.ActionSelectArene()
					break
				case "select_arene_pokemon":
					graphic.ConfirmSelectArene()
					break
				case "in_battle_arene":
					graphic.ActionArene()
					break
				case "badge":
					graphic.ActionBadge()
					break
				default:
					continue
				}
			}
		}
	}
}
