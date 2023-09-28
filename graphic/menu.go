package graphic

import (
	"game/utils"
	"game/values"
	"github.com/fatih/color"
	"os"
)

func UpdateMenu() {
	utils.ClearTerminal()
	color.Yellow("------------------------------- Pokemon -------------------------------")
	switch values.MenuIndex {
	case 1:
		color.Green("Se balader")
		println("Boutique")
		println("Sac à dos")
		println("Quitter")
		break
	case 2:
		println("Se balader")
		color.Green("Boutique")
		println("Sac à dos")
		println("Quitter")
		break
	case 3:
		println("Se balader")
		println("Boutique")
		color.Green("Sac à dos")
		println("Quitter")
		break
	case 4:
		println("Se balader")
		println("Boutique")
		println("Sac à dos")
		color.Green("Quitter")
		break
	}
}

func ActionMenuu() {
	switch values.MenuIndex {
	case 1:
		color.Green("Se balader")
		println("Boutique")
		println("Sac à dos")
		println("Quitter")
		break
	case 2:
		values.MenuIndex = 1
		values.CurrentPage = "shop"
		UpdateShop()
		break
	case 3:
		values.MenuIndex = 1
		values.CurrentPage = "bag"
		UpdateBag()
		break
	case 4:
		os.Exit(0)
	}
}
