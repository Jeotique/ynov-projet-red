package graphic

import (
	"fmt"
	"game/bdd"
	"game/utils"
	"game/values"
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
	"time"
)

func LookForEnemy() {
	utils.ClearTerminal()
	color.Yellow("------------------------------- Pokemon -------------------------------")
	fmt.Println()
	color.Green("                  _._       _,._\n                        _.'   `. ' .'   _`.\n                ,\"\"\"/`\"\"-.-.,/. ` V'\\-,`.,--/\"\"\".\"-..\n              ,'    `...,' . ,\\-----._|     `.   /   \\\n             `.            .`  -'`\"\" .._   :> `-'   `.\n            ,'  ,-.  _,.-'| `..___ ,'   |'-..__   .._ L\n           .    \\_ -'   `-'     ..      `.-' `.`-.'_ .|\n           |   ,',-,--..  ,--../  `.  .-.    , `-.  ``.\n           `.,' ,  |   |  `.  /'/,,.\\/  |    \\|   |\n                `  `---'    `j   .   \\  .     '   j\n              ,__`\"        ,'|`'\\_/`.'\\'        |\\-'-, _,.\n       .--...`-. `-`. /    '- ..      _,    /\\ ,' .--\"'  ,'\".\n     _'-\"\"-    --  _`'-.../ __ '.'`-^,_`-\"\"\"\"---....__  ' _,-`\n   _.----`  _..--.'        |  \"`-..-\" __|'\"'         .\"\"-. \"\"'--.._\n  /        '    /     ,  _.+-.'  ||._'   \"\"\"\". .          `     .__\\\n `---    /        /  / j'       _/|..`  -. `-`\\ \\   \\  \\   `.  \\ `-..\n,\" _.-' /    /` ./  /`_|_,-\"   ','|       `. | -'`._,   L  \\ .  `.   |\n`\"' /  /  / ,__...-----| _.,  ,'            `|----.._`-.|' |. .` ..  .\n   /  '| /.,/   \\--.._ `-,' ,          .  '`.'  __,., '  ''``._ \\ \\`,'\n  /_,'---  ,     \\`._,-` \\ //  / . \\    `._,  -`,  / / _   |   `-L -\n   /       `.     ,  ..._ ' `_/ '| |\\ `._'       '-.'   `.,'     |\n  '         /    /  ..   `.  `./ | ; `.'    ,\"\" ,.  `.    \\      |\n   `.     ,'   ,'   | |\\  |       \"        |  ,'\\ |   \\    `    ,L\n   /|`.  /    '     | `-| '                  /`-' |    L    `._/  \\\n  / | .`|    |  .   `._.'                   `.__,'   .  |     |  (`\n '-\"\"-'_|    `. `.__,._____     .    _,        ____ ,-  j     \".-'\"'\n        \\      `-.  \\/.    `\"--.._    _,.---'\"\"\\/  \"_,.'     /-'\n         )        `-._ '-.        `--\"      _.-'.-\"\"        `.\n        ./            `,. `\".._________...\"\"_.-\"`.          _j\n       /_\\.__,\"\".   ,.'  \"`-...________.---\"     .\".   ,.  / \\\n              \\_/\"\"\"-'                           `-'--(_,`\"`-`")
	fmt.Println()
	color.Red("		      En recherche d'adversaire")
	go utils.PlaySound(values.Sounds["waiting"])
	s := spinner.New(spinner.CharSets[38], 100*time.Millisecond)
	s.FinalMSG = "Game loaded !\n"
	s.Color("bold")
	s.Start()
	time.Sleep(6000 * time.Millisecond)
	fmt.Println()
	s.Stop()
	values.EnemyBalade = &values.Pokemons[utils.RandomNumber(0, len(values.Pokemons)-1)]
	values.CurrentPage = "announce_battle"
	AnnounceBattle()
}

func AnnounceBattle() {
	go utils.PlaySound(values.Sounds["battle"])
	utils.ClearTerminal()
	utils.Writeanim(values.PokeArt[values.EnemyBalade.GetName()], 1)
	fmt.Println()
	fmt.Println()
	utils.Writeanim("Un "+values.EnemyBalade.GetName()+" vient d'apparaître !", 10)
	time.Sleep(2000 * time.Millisecond)
	values.MenuIndex = 1
	values.CurrentPage = "in_battle"
	BattleBalade()
}

func BattleBalade() {
	utils.ClearTerminal()
	fmt.Println()
	fmt.Println()
	data := [][]string{
		[]string{values.PokeArt[values.PokemonBalade.Nom], values.PokeArt[values.EnemyBalade.Nom]},
		[]string{"\n\n", "\n\n"},
		[]string{"Vie : " + strconv.Itoa(values.PokemonBalade.Vie) + "/" + strconv.Itoa(values.PokemonBalade.MaxVie), "Vie : " + strconv.Itoa(values.EnemyBalade.Vie) + "/" + strconv.Itoa(values.EnemyBalade.MaxVie)},
		[]string{"", "Chance de capture : " + strconv.Itoa(values.EnemyBalade.ChanceCapture) + "%"},
	}
	table := tablewriter.NewWriter(os.Stdout)
	for _, v := range data {
		table.Append(v)
	}
	table.Render()

	if !values.TryingCapture {

		if values.PlayerTour {
			if values.EnemyDebrief {
				if values.EnemyVeski {
					utils.Writeanim("Vous avez esquivé !", 10)
				} else {
					utils.Writeanim("Vous avez perdu "+strconv.Itoa(values.LastDamage)+"PV", 10)
				}
				values.TryingCapture = false
				values.EnemyDebrief = false
				values.PlayerTour = true
				values.PlayerDebrief = false
				values.LastDamage = 0
				values.EnemyVeski = false
				time.Sleep(1500 * time.Millisecond)
				BattleBalade()
			} else {
				for i, attack := range values.PokemonBalade.GetAttacks() {
					if i+1 == values.MenuIndex {
						color.Green("> %s", attack.Nom)
						color.Blue("     Classe : [%s]", attack.Classe)
						color.Yellow("     Dégat min/max : [%d/%d]", attack.DegatMin, attack.DegatMax)
						color.Magenta("     [%s]", attack.Description)
					} else {
						fmt.Println(attack.Nom)
					}
				}
				if values.MenuIndex > len(values.PokemonBalade.GetAttacks()) {
					if values.MainCharacter.PotionPoison < 1 {
						canCapture := true
						for _, m := range values.MainCharacter.Pokemons {
							if m.Nom == values.EnemyBalade.Nom {
								canCapture = false
							}
						}
						if canCapture {
							if values.MainCharacter.Pokeball > 0 {
								if len(values.MainCharacter.Pokemons) < values.MainCharacter.PokemonMax {
									color.Green("> Capturer %s [%d pokeball]", values.EnemyBalade.Nom, values.MainCharacter.Pokeball)
								} else {
									color.Red("> Capturer %s", values.EnemyBalade.Nom)
									color.Red("	[X] Capacité max atteinte")
								}
							} else {
								color.Red("> Capturer %s", values.EnemyBalade.Nom)
								color.Red("	[X] Vous n'avez pas de pokeball")
							}
						} else {
							color.Red("> Capturer %s", values.EnemyBalade.Nom)
							color.Red("	[X] Vous avez déjà ce pokemon")
						}
					} else if values.MenuIndex == len(values.PokemonBalade.GetAttacks())+1 {
						color.Green("> Utiliser le poison")
						fmt.Println("Capturer", values.EnemyBalade.Nom)
					} else {
						fmt.Println("Utiliser le poison")
						canCapture := true
						for _, m := range values.MainCharacter.Pokemons {
							if m.Nom == values.EnemyBalade.Nom {
								canCapture = false
							}
						}
						if canCapture {
							if values.MainCharacter.Pokeball > 0 {
								if len(values.MainCharacter.Pokemons) < values.MainCharacter.PokemonMax {
									color.Green("> Capturer %s [%d pokeball]", values.EnemyBalade.Nom, values.MainCharacter.Pokeball)
								} else {
									color.Red("> Capturer %s", values.EnemyBalade.Nom)
									color.Red("	[X] Capacité max atteinte")
								}
							} else {
								color.Red("> Capturer %s", values.EnemyBalade.Nom)
								color.Red("	[X] Vous n'avez pas de pokeball")
							}
						} else {
							color.Red("> Capturer %s", values.EnemyBalade.Nom)
							color.Red("	[X] Vous avez déjà ce pokemon")
						}
					}
				} else {
					if values.MainCharacter.PotionPoison > 0 {
						fmt.Println("Utiliser le poison")
					}
					fmt.Println("Capturer", values.EnemyBalade.Nom)
				}
			}
		} else {
			if values.PlayerDebrief {
				if values.EnemyVeski {
					utils.Writeanim(values.EnemyBalade.Nom+" a esquivé !", 10)
				} else {
					utils.Writeanim("Vous avez retiré "+strconv.Itoa(values.LastDamage)+"PV", 10)
				}
				time.Sleep(1500 * time.Millisecond)
				values.TryingCapture = false
				values.EnemyDebrief = false
				values.PlayerTour = false
				values.PlayerDebrief = false
				values.LastDamage = 0
				values.EnemyVeski = false
				BattleBalade()
			} else {
				attack := values.EnemyBalade.GetAttacks()[utils.RandomNumber(0, len(values.EnemyBalade.GetAttacks())-1)]
				utils.Writeanim(values.EnemyBalade.Nom+" utilise "+attack.Nom, 10)
				time.Sleep(1500 * time.Millisecond)
				ToRemove := utils.RandomNumber(attack.DegatMin, attack.DegatMax)
				if utils.RandomNumber(0, 3) == 1 {
					values.EnemyVeski = true
					values.MenuIndex = 1
					values.EnemyDebrief = true
					values.PlayerTour = true
					BattleBalade()
				} else {
					if values.IsCounter(values.EnemyBalade.Classe, values.PokemonBalade.Classe) {
						ToRemove /= 2
					}
					values.PokemonBalade.Vie -= ToRemove
					if values.PokemonBalade.Vie < 0 {
						values.PokemonBalade.Vie = 0
					}
					if values.PokemonBalade.Vie == 0 {
						BaladeLoser()
					} else {
						values.PlayerTour = true
						values.MenuIndex = 1
						values.EnemyDebrief = true
						values.LastDamage = ToRemove
						BattleBalade()
					}
				}
			}
		}
	} else {
		utils.Writeanim("Vous tentez de capturer "+values.EnemyBalade.Nom+" !\n", 10)
		utils.Writeanim("Pokeball jeté !\n", 10)
		time.Sleep(1500 * time.Millisecond)
		utils.Writeanim("Hmm...\n", 10)
		time.Sleep(1500 * time.Millisecond)
		utils.Writeanim("Hmmmmm...\n", 10)
		time.Sleep(1500 * time.Millisecond)
		utils.Writeanim("Hmmmmmmm...\n", 10)
		var chance []int
		for i := 0; i < 100; i++ {
			if i <= values.EnemyBalade.GetCaptureChance() {
				chance = append(chance, 1)
			} else {
				chance = append(chance, 0)
			}
		}
		IsCaptured := chance[utils.RandomNumber(0, len(chance)-1)]
		if IsCaptured == 1 {
			values.TryingCapture = false
			values.EnemyDebrief = false
			values.PlayerTour = true
			values.PlayerDebrief = false
			values.LastDamage = 0
			values.EnemyVeski = false
			BaladeCaptured()
		} else {
			color.Red("Raté !")
			values.TryingCapture = false
			values.EnemyDebrief = false
			values.PlayerTour = false
			values.PlayerDebrief = false
			values.LastDamage = 0
			values.EnemyVeski = false
			time.Sleep(1500 * time.Millisecond)
			BattleBalade()
		}
	}
}

func ActionBattle() {
	if values.MenuIndex > len(values.PokemonBalade.GetAttacks()) {
		if values.MainCharacter.PotionPoison > 0 {
			if values.MenuIndex == len(values.PokemonBalade.GetAttacks())+1 {
				values.MainCharacter.PotionPoison--
				bdd.Database.Set("character_potionpoison", values.MainCharacter.PotionPoison)
				values.EnemyBalade.Vie -= 30
				values.EnemyBalade.ChanceCapture += 20
				if values.EnemyBalade.Vie < 0 {
					values.EnemyBalade.Vie = 0
				}
				if values.EnemyBalade.Vie == 0 {
					BaladeWinner()
				} else {
					values.PlayerTour = false
					values.MenuIndex = 1
					values.PlayerDebrief = true
					values.LastDamage = 30
					BattleBalade()
				}
			} else {
				canCapture := true
				for _, m := range values.MainCharacter.Pokemons {
					if m.Nom == values.EnemyBalade.Nom {
						canCapture = false
					}
				}
				if canCapture {
					if values.MainCharacter.Pokeball > 0 {
						if len(values.MainCharacter.Pokemons) < values.MainCharacter.PokemonMax {
							values.MenuIndex = 1
							values.TryingCapture = true
							values.PlayerTour = false
							values.MainCharacter.Pokeball--
							bdd.Database.Set("character_pokeball", values.MainCharacter.Pokeball)
							BattleBalade()
						}
					}
				}
			}
		} else {
			canCapture := true
			for _, m := range values.MainCharacter.Pokemons {
				if m.Nom == values.EnemyBalade.Nom {
					canCapture = false
				}
			}
			if canCapture {
				if values.MainCharacter.Pokeball > 0 {
					if len(values.MainCharacter.Pokemons) < values.MainCharacter.PokemonMax {
						values.MenuIndex = 1
						values.TryingCapture = true
						values.PlayerTour = false
						values.MainCharacter.Pokeball--
						bdd.Database.Set("character_pokeball", values.MainCharacter.Pokeball)
						BattleBalade()
					}
				}
			}
		}
	} else {
		// attaque classique
		attack := values.PokemonBalade.GetAttacks()[values.MenuIndex-1]
		ToRemove := utils.RandomNumber(attack.DegatMin, attack.DegatMax)
		if utils.RandomNumber(0, 3) == 1 {
			values.EnemyVeski = true
			values.MenuIndex = 1
			values.PlayerDebrief = true
			values.PlayerTour = false
			BattleBalade()
		} else {
			if values.IsCounter(values.PokemonBalade.Classe, values.EnemyBalade.Classe) {
				ToRemove /= 2
			}
			values.EnemyBalade.Vie -= ToRemove
			values.EnemyBalade.ChanceCapture += ToRemove / 2
			if values.EnemyBalade.Vie < 0 {
				values.EnemyBalade.Vie = 0
			}
			if values.EnemyBalade.Vie == 0 {
				BaladeWinner()
			} else {
				values.PlayerTour = false
				values.MenuIndex = 1
				values.PlayerDebrief = true
				values.LastDamage = ToRemove
				BattleBalade()
			}
		}
	}
}

func BaladeWinner() {
	values.MenuIndex = 1
	values.CurrentPage = "winner"
	bdd.Database.SavePokemons(values.MainCharacter.Pokemons)
	values.EnemyBalade.Vie = values.EnemyBalade.MaxVie
	values.EnemyBalade.ChanceCapture = values.EnemyBalade.DefaultChanceCapture
	xpWin := utils.RandomNumber(10, 25)
	values.MainCharacter.AddXp(xpWin)
	bdd.Database.Set("character_xp", values.MainCharacter.Xp)
	bdd.Database.Set("character_level", values.MainCharacter.Level)

	argentWin := utils.RandomNumber(15, 50)
	values.MainCharacter.Argent += argentWin
	bdd.Database.Set("character_argent", values.MainCharacter.Argent)
	utils.ClearTerminal()
	fmt.Println()
	fmt.Println()
	go utils.PlaySound(values.Sounds["win"])
	color.Green(values.PokeArt[values.EnemyBalade.GetName()])
	fmt.Println()
	fmt.Println()
	utils.Writeanim("Vous avez battu un "+values.EnemyBalade.GetName(), 10)
	utils.Writeanim("\nArgent gagné : "+strconv.Itoa(argentWin), 10)
	time.Sleep(3000 * time.Millisecond)
	values.MenuIndex = 1
	values.CurrentPage = "menu"
	UpdateMenu()
}

func BaladeLoser() {
	values.MenuIndex = 1
	values.CurrentPage = "loser"
	values.PokemonBalade.Vie = values.PokemonBalade.MaxVie / 2
	bdd.Database.SavePokemons(values.MainCharacter.Pokemons)
	values.EnemyBalade.Vie = values.EnemyBalade.MaxVie
	values.EnemyBalade.ChanceCapture = values.EnemyBalade.DefaultChanceCapture

	utils.ClearTerminal()
	fmt.Println()
	fmt.Println()
	go utils.PlaySound(values.Sounds["lose"])
	color.Red(values.PokeArt[values.EnemyBalade.GetName()])
	fmt.Println()
	fmt.Println()
	utils.Writeanim("Vous avez été vaincu...\n", 10)
	if values.MainCharacter.PotionVie < 1 {
		values.MainCharacter.PotionVie = 1
		bdd.Database.Set("character_potionvie", values.MainCharacter.PotionVie)
		utils.Writeanim("Une potion de vie ajouté à l'inventaire.\n", 10)
	}
	time.Sleep(3000 * time.Millisecond)
	values.MenuIndex = 1
	values.CurrentPage = "menu"
	UpdateMenu()
}

func BaladeCaptured() {
	values.MenuIndex = 1
	values.CurrentPage = "captured"
	pokemonToAdd := values.FindPokemon(values.EnemyBalade.GetName())
	pokemonToAdd.Vie = values.EnemyBalade.Vie
	values.MainCharacter.Pokemons = append(values.MainCharacter.Pokemons, pokemonToAdd)
	bdd.Database.AddPokemon(values.EnemyBalade.GetName())

	values.EnemyBalade.Vie = values.EnemyBalade.MaxVie
	values.EnemyBalade.ChanceCapture = values.EnemyBalade.DefaultChanceCapture

	utils.ClearTerminal()
	fmt.Println()
	fmt.Println()
	go utils.PlaySound(values.Sounds["capture"])
	color.Yellow(values.PokeArt[values.EnemyBalade.GetName()])
	fmt.Println()
	fmt.Println()
	utils.Writeanim("Vous avez capturé un "+values.EnemyBalade.GetName(), 10)
	time.Sleep(3000 * time.Millisecond)
	values.MenuIndex = 1
	values.CurrentPage = "menu"
	UpdateMenu()
}
