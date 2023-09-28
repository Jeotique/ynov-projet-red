package values

var MenuIndex = 1
var MainCharacter Personnage
var CurrentPage = "menu"
var PokemonInfo *Pokemon
var AttackInfo *Attaque
var PokemonBalade *Pokemon
var EnemyBalade *Pokemon
var PlayerTour bool = true
var EnemyVeski bool
var PlayerDebrief bool
var EnemyDebrief bool
var LastDamage int
var TryingCapture bool
var MaitrePlaying *Maitre

var Sounds = map[string]string{
	"error":    "./sounds/beep-error.mp3",
	"success":  "./sounds/beep-success.mp3",
	"navigate": "./sounds/navigate.mp3",
	"click":    "./sounds/click.mp3",
	"heal":     "./sounds/healv2.mp3",
	"opening":  "./sounds/opening.mp3",
	"upgrade":  "./sounds/upgrade.mp3",
	"battle":   "./sounds/battle.mp3",
	"waiting":  "./sounds/waiting.mp3",
	"capture":  "./sounds/capture.mp3",
	"win":      "./sounds/win.mp3",
	"lose":     "./sounds/lose.mp3",
}
