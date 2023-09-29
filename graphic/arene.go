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
	"time"
)

func AnnounceArene() {
	go utils.PlaySound(values.Sounds["battle"])
	utils.ClearTerminal()
	utils.Writeanim(values.PokeArt[values.EnemyBalade.GetName()], 0)
	fmt.Println()
	fmt.Println()
	utils.Writeanim("Un "+values.EnemyBalade.GetName()+" vient d'apparaître avec "+values.MaitrePlaying.Nom+" !", 10)
	time.Sleep(2000 * time.Millisecond)
	values.MenuIndex = 1
	values.CurrentPage = "in_battle_arene"
	BattleArene()
}

func BattleArene() {
	utils.ClearTerminal()
	fmt.Println()
	fmt.Println()
	data := [][]string{
		[]string{values.PokeArt[values.PokemonBalade.Nom], values.PokeArt[values.EnemyBalade.Nom]},
		[]string{"\n\n", "\n\n"},
		//[]string{"Vie : " + strconv.Itoa(values.PokemonBalade.Vie) + "/" + strconv.Itoa(values.PokemonBalade.MaxVie), "Vie : " + strconv.Itoa(values.EnemyBalade.Vie) + "/" + strconv.Itoa(values.EnemyBalade.MaxVie)},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	for _, v := range data {
		table.Append(v)
	}
	table.SetFooter([]string{"Type : " + values.PokemonBalade.Classe + "\nVie : " + strconv.Itoa(values.PokemonBalade.Vie) + "/" + strconv.Itoa(values.PokemonBalade.MaxVie),
		"Type : " + values.EnemyBalade.Classe + "\nVie : " + strconv.Itoa(values.EnemyBalade.Vie) + "/" + strconv.Itoa(values.EnemyBalade.MaxVie),
	})
	table.Render()

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
			BattleArene()
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

				} else if values.MenuIndex == len(values.PokemonBalade.GetAttacks())+1 {
					color.Green("> Utiliser le poison")
				} else {
					fmt.Println("Utiliser le poison")
				}
			} else {
				if values.MainCharacter.PotionPoison > 0 {
					fmt.Println("Utiliser le poison")
				}
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
			BattleArene()
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
				BattleArene()
			} else {
				if values.IsCounter(values.EnemyBalade.Classe, values.PokemonBalade.Classe) {
					ToRemove /= 2
				}
				values.PokemonBalade.Vie -= ToRemove
				if values.PokemonBalade.Vie < 0 {
					values.PokemonBalade.Vie = 0
				}
				if values.PokemonBalade.Vie == 0 {
					AreneLoser()
				} else {
					values.PlayerTour = true
					values.MenuIndex = 1
					values.EnemyDebrief = true
					values.LastDamage = ToRemove
					BattleArene()
				}
			}
		}
	}

}

func ActionArene() {
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
					AreneWinner()
				} else {
					values.PlayerTour = false
					values.MenuIndex = 1
					values.PlayerDebrief = true
					values.LastDamage = 30
					BattleArene()
				}
			} else {

			}
		} else {

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
			BattleArene()
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
				AreneWinner()
			} else {
				values.PlayerTour = false
				values.MenuIndex = 1
				values.PlayerDebrief = true
				values.LastDamage = ToRemove
				BattleArene()
			}
		}
	}
}

func AreneWinner() {
	values.MenuIndex = 1
	values.CurrentPage = "winner"
	bdd.Database.SavePokemons(values.MainCharacter.Pokemons)
	values.MainCharacter.Badges = append(values.MainCharacter.Badges, values.MaitrePlaying.Badge)
	bdd.Database.SaveBadges(values.MainCharacter.Badges)
	values.EnemyBalade.Vie = values.EnemyBalade.MaxVie
	xpWin := utils.RandomNumber(10, 25)
	values.MainCharacter.AddXp(xpWin)
	values.MainCharacter.AddLevel(utils.RandomNumber(3, 8))
	bdd.Database.Set("character_xp", values.MainCharacter.Xp)
	bdd.Database.Set("character_level", values.MainCharacter.Level)

	argentWin := utils.RandomNumber(50, 150)
	values.MainCharacter.Argent += argentWin
	bdd.Database.Set("character_argent", values.MainCharacter.Argent)
	utils.ClearTerminal()
	fmt.Println()
	fmt.Println()
	go utils.PlaySound(values.Sounds["win"])
	color.Green(values.PokeArt[values.EnemyBalade.GetName()])
	fmt.Println()
	fmt.Println()
	utils.Writeanim("Vous avez battu un "+values.EnemyBalade.GetName()+" et "+values.MaitrePlaying.Nom, 10)
	utils.Writeanim("\nArgent gagné : "+strconv.Itoa(argentWin), 10)
	time.Sleep(3000 * time.Millisecond)
	values.MenuIndex = 1
	values.CurrentPage = "menu"
	UpdateMenu()
}

func AreneLoser() {
	values.MenuIndex = 1
	values.CurrentPage = "loser"
	values.PokemonBalade.Vie = values.PokemonBalade.MaxVie / 2
	bdd.Database.SavePokemons(values.MainCharacter.Pokemons)
	values.EnemyBalade.Vie = values.EnemyBalade.MaxVie
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
