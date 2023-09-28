package graphic

import (
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
		println("Pokémons")
		println("Inventaire")
		color.Green("Badges")
		println("Quitter")
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
		[]string{"Pokemon(s)", strconv.Itoa(len(values.MainCharacter.Pokemons))},
		[]string{"Capacité max", strconv.Itoa(values.MainCharacter.PokemonMax)},
		[]string{"Pokeball", strconv.Itoa(values.MainCharacter.Pokeball)},
		[]string{"Potion de vie", strconv.Itoa(values.MainCharacter.PotionVie)},
		[]string{"Potion de poison", strconv.Itoa(values.MainCharacter.PotionPoison)},
	}
	table := tablewriter.NewWriter(os.Stdout)
	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}
