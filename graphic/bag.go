package graphic

import (
	"fmt"
	"game/utils"
	"game/values"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
)

func UpdateBag() {
	utils.ClearTerminal()
	color.Yellow("------------------------------- Pokemon -------------------------------")
	switch values.MenuIndex {
	case 1:
		color.Green("Pokémons")
		println("Inventaire")
		println("Badges")
		println("Quitter")
		break
	case 2:
		println("Pokémons")
		color.Green("Inventaire")
		println("Badges")
		println("Quitter")
		break
	case 3:
		println("Pokémons")
		println("Inventaire")
		color.Green("Badges")
		println("Quitter")
		break
	case 4:
		println("Pokémons")
		println("Inventaire")
		println("Badges")
		color.Green("Quitter")
		break
	}
}

func ActionBag() {
	switch values.MenuIndex {
	case 1:
		values.MenuIndex = 1
		values.CurrentPage = "pokemons"
		UpdatePokemonsBag()
		break
	case 2:
		values.CurrentPage = "inventory"
		DisplayInventory()
		break
	case 3:
		values.MenuIndex = 1
		values.CurrentPage = "badge"
		UpdateBadge()
		break
	case 4:
		values.MenuIndex = 1
		values.CurrentPage = "menu"
		UpdateMenu()
		break
	}
}

func DisplayInventory() {
	utils.ClearTerminal()
	color.Yellow("------------------------------- Pokemon -------------------------------")
	data := [][]string{
		[]string{"Argent", strconv.Itoa(values.MainCharacter.Argent)},
		[]string{"Expérience", strconv.Itoa(values.MainCharacter.Xp) + "/50"},
		[]string{"Niveau", strconv.Itoa(values.MainCharacter.Level)},
		[]string{"Pokemon(s)", strconv.Itoa(len(values.MainCharacter.Pokemons))},
		[]string{"Capacité max", strconv.Itoa(values.MainCharacter.PokemonMax)},
		[]string{"Badge(s)", strconv.Itoa(len(values.MainCharacter.Badges))},
		[]string{"Pokeball", strconv.Itoa(values.MainCharacter.Pokeball)},
		[]string{"Potion de vie", strconv.Itoa(values.MainCharacter.PotionVie)},
		[]string{"Potion de poison", strconv.Itoa(values.MainCharacter.PotionPoison)},
		[]string{"Pierre de nuit", strconv.Itoa(values.MainCharacter.PierreNuit)},
	}
	table := tablewriter.NewWriter(os.Stdout)
	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}

func UpdateBadge() {
	utils.ClearTerminal()
	color.Yellow("------------------------------- Pokemon -------------------------------")
	for i, badge := range values.MainCharacter.Badges {
		if i+1 == values.MenuIndex {
			color.Green("> %s", badge.Nom)
			color.Blue("\t[%s]", badge.Description)
			color.Yellow("	Gardien : [%s]", badge.Maitre)
			color.Magenta("	Arene : [%s]", badge.Arene)
			color.Cyan("\tPokemon : [%s]", badge.Pokemon)
		} else {
			fmt.Println(badge.Nom)
		}
	}
	if len(values.MainCharacter.Badges) < 1 {
		color.Red("Vous n'avez aucun badge. Rendez-vous dans l'arene ;)")
	}
	if values.MenuIndex > len(values.MainCharacter.Badges) {
		color.Green("> Quitter")
	} else {
		fmt.Println("Quitter")
	}
}

func ActionBadge() {
	if values.MenuIndex > len(values.MainCharacter.Badges) {
		values.MenuIndex = 1
		values.CurrentPage = "bag"
		UpdateBag()
	}
}
