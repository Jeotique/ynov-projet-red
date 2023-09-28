package values

import "game/utils"

type Pokemon struct {
	nom           string
	classe        string
	description   string
	vie           int
	maxVie        int
	chanceCapture int
	attaques      []Attaque
}

type Attaque struct {
	nom         string
	description string
	classe      string
	degatMin    int
	degatMax    int
	prixVie     int
}

func (p Pokemon) GetName() string {
	return p.nom
}

func FindPokemon(nom string) Pokemon {
	var Result Pokemon
	for _, M := range Pokemons {
		if M.nom == nom {
			Result = M
		}
	}
	return Result
}

func GetRandomPokemon() Pokemon {
	return Pokemons[utils.RandomNumber(0, len(Pokemons)-1)]
}
