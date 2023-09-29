package values

import "game/utils"

type Pokemon struct {
	Nom                  string    `json:"Nom"`
	Classe               string    `json:"Classe"`
	Description          string    `json:"Description"`
	Vie                  int       `json:"Vie"`
	MaxVie               int       `json:"MaxVie"`
	ChanceCapture        int       `json:"ChanceCapture"`
	DefaultChanceCapture int       `json:"DefaultChanceCapture"`
	Attaques             []Attaque `json:"Attaques"`
}

type Attaque struct {
	Nom         string `json:"Nom"`
	Description string `json:"Description"`
	Classe      string `json:"Classe"`
	DegatMin    int    `json:"DegatMin"`
	DegatMax    int    `json:"DegatMax"`
	PrixVie     int    `json:"PrixVie"`
}

func (p Pokemon) GetName() string {
	return p.Nom
}

func (p Pokemon) GetClass() string {
	return p.Classe
}

func (p Pokemon) GetDescription() string {
	return p.Description
}

func (p Pokemon) GetLife() int {
	return p.Vie
}

func (p Pokemon) GetMaxLife() int {
	return p.MaxVie
}

func (p Pokemon) GetCaptureChance() int {
	return p.ChanceCapture
}

func (p Pokemon) GetAttacks() []Attaque {
	return p.Attaques
}

func (p *Pokemon) Heal() {
	p.Vie = p.MaxVie
}

func FindPokemon(nom string) Pokemon {
	var Result Pokemon
	for _, M := range Pokemons {
		if M.Nom == nom {
			Result = M
		}
	}
	return Result
}

func GetRandomPokemon() Pokemon {
	return Pokemons[utils.RandomNumber(0, len(Pokemons)-1)]
}

func IsCounter(poke1 string, poke2 string) bool {
	if poke1 == poke2 {
		return false
	} else if poke1 == "Psychique" && poke2 == "Psychique" {
		return true
	} else if poke1 == "Feu" && poke2 == "Eau" {
		return true
	} else if poke1 == "Eau" && poke2 == "Normal" {
		return true
	} else if poke1 == "Plante" && poke2 == "Feu" {
		return true
	} else if poke1 == "Eau" && poke2 == "Plante" {
		return true
	} else if poke1 != "Psychique" && poke2 == "Psychique " {
		return true
	} else if poke1 == "Normal" && poke2 == "Vol" {
		return true
	} else if poke1 == "Électrique" && poke2 == "Roche" {
		return true
	} else if poke1 == "Roche" && poke2 == "Feu" {
		return true
	} else if poke1 == "Roche" && poke2 == "Eau" {
		return true
	} else {
		return false
	}
}
