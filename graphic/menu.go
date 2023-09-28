package graphic

import (
	"game/utils"
	"game/values"
	"github.com/fatih/color"
	"os"
	"time"
)

func UpdateMenu() {
	utils.ClearTerminal()
	color.Yellow("------------------------------- Pokemon -------------------------------")
	switch values.MenuIndex {
	case 1:
		color.Green("Se balader")
		println("Boutique")
		println("Sac à dos")
		println("Qui sont-ils")
		println("Quitter")
		break
	case 2:
		println("Se balader")
		color.Green("Boutique")
		println("Sac à dos")
		println("Qui sont-ils")
		println("Quitter")
		break
	case 3:
		println("Se balader")
		println("Boutique")
		color.Green("Sac à dos")
		println("Qui sont-ils")
		println("Quitter")
		break
	case 4:
		println("Se balader")
		println("Boutique")
		println("Sac à dos")
		color.Green("Qui sont-ils")
		println("Quitter")
		break
	case 5:
		println("Se balader")
		println("Boutique")
		println("Sac à dos")
		println("Qui sont-ils")
		color.Green("Quitter")
		break
	}
}

func ActionMenuu() {
	switch values.MenuIndex {
	case 1:
		values.MenuIndex = 1
		values.CurrentPage = "balade"
		UpdateBalade()
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
		utils.ClearTerminal()
		utils.Writeanim("Spielberg, Abba et Queen", 20)
		time.Sleep(2 * time.Second)
		UpdateMenu()
		break
	case 5:
		os.Exit(0)
	}
}
