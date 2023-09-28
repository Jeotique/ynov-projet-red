package values

type Personnage struct {
	Nom          string
	Pokemons     []Pokemon
	PokemonMax   int
	PotionVie    int
	PotionPoison int
	Pokeball     int
	PierreNuit   int
	Argent       int
	Xp           int
	Level        int
	Badges       []Badge
}

func (p *Personnage) AddXp(nb int) {
	p.Xp += nb
	if p.Xp >= 50 {
		p.Level++
		p.Xp = 0
	}
}

func (p *Personnage) AddLevel(nb int) {
	p.Level += nb
}

type Badge struct {
	Nom         string
	Description string
	Maitre      string
	Arene       string
	Pokemon     string
}

type Maitre struct {
	Nom     string
	Arene   string
	Pokemon Pokemon
	Badge   Badge
	Vaincu  bool
}

var Maitres []Maitre = []Maitre{
	{
		Nom:   "Henry",
		Arene: "Natural Park",
		Pokemon: Pokemon{
			Nom:           "Torterra",
			Classe:        "Plante",
			Description:   "Un Pokémon Plante/Sol puissant et imposant",
			Vie:           120,
			MaxVie:        120,
			ChanceCapture: 0,
			Attaques: []Attaque{
				{
					Nom:         "Fouet Géant",
					Description: "Fouette l'adversaire avec une liane géante",
					Classe:      "Plante",
					DegatMin:    20,
					DegatMax:    40,
					PrixVie:     0,
				},
				{
					Nom:         "Séisme",
					Description: "Déclenche un séisme pour engloutir l'adversaire",
					Classe:      "Sol",
					DegatMin:    30,
					DegatMax:    60,
					PrixVie:     0,
				},
				{
					Nom:         "Morsure Féroce",
					Description: "Mord l'adversaire avec une force redoutable",
					Classe:      "Normal",
					DegatMin:    25,
					DegatMax:    50,
					PrixVie:     0,
				},
				{
					Nom:         "Éclosion Sylvestre",
					Description: "Déclenche une éruption de la nature tout autour de lui",
					Classe:      "Plante",
					DegatMin:    40,
					DegatMax:    80,
					PrixVie:     0,
				},
			},
		},
		Badge: Badge{
			Nom:         "Natural Badge",
			Description: "Un badge remis au courageux chasseur ayant battu Henry et Torterra",
			Maitre:      "Henry",
			Arene:       "Natural Park",
			Pokemon:     "Torterra",
		},
	},
	{
		Nom:   "Thomas",
		Arene: "Oceanic Arena",
		Pokemon: Pokemon{
			Nom:           "Laggron",
			Classe:        "Eau",
			Description:   "Un Pokémon Eau/Sol massif et redoutable",
			Vie:           140,
			MaxVie:        140,
			ChanceCapture: 0,
			Attaques: []Attaque{
				{
					Nom:         "Raz-de-Marée",
					Description: "Crée un raz-de-marée dévastateur",
					Classe:      "Eau",
					DegatMin:    40,
					DegatMax:    80,
					PrixVie:     0,
				},
				{
					Nom:         "Séisme",
					Description: "Déclenche un séisme pour engloutir l'adversaire",
					Classe:      "Sol",
					DegatMin:    30,
					DegatMax:    60,
					PrixVie:     0,
				},
				{
					Nom:         "Hydrocanon",
					Description: "Déchaîne un puissant jet d'eau sur l'adversaire",
					Classe:      "Eau",
					DegatMin:    25,
					DegatMax:    50,
					PrixVie:     0,
				},
				{
					Nom:         "Morsure Féroce",
					Description: "Mord l'adversaire avec une force redoutable",
					Classe:      "Normal",
					DegatMin:    25,
					DegatMax:    50,
					PrixVie:     0,
				},
			},
		},
		Badge: Badge{
			Nom:         "Oceanic Badge",
			Description: "Un badge remis au courageux chasseur ayant battu Thomas et Laggron à l'Oceanic Arena",
			Maitre:      "Thomas",
			Arene:       "Oceanic Arena",
			Pokemon:     "Laggron",
		},
	},
	{
		Nom:   "Alan",
		Arene: "Volcano Arena",
		Pokemon: Pokemon{
			Nom:           "Typhlosion",
			Classe:        "Feu",
			Description:   "Un Pokémon de feu explosif et féroce",
			Vie:           160,
			MaxVie:        160,
			ChanceCapture: 0,
			Attaques: []Attaque{
				{
					Nom:         "Explosion de Feu",
					Description: "Déclenche une explosion de feu intense",
					Classe:      "Feu",
					DegatMin:    40,
					DegatMax:    80,
					PrixVie:     0,
				},
				{
					Nom:         "Déflagration",
					Description: "Déclenche une déflagration de feu",
					Classe:      "Feu",
					DegatMin:    30,
					DegatMax:    60,
					PrixVie:     0,
				},
				{
					Nom:         "Lance-Flammes",
					Description: "Lance un puissant jet de flammes sur l'adversaire",
					Classe:      "Feu",
					DegatMin:    25,
					DegatMax:    50,
					PrixVie:     0,
				},
				{
					Nom:         "Griffe Incendiaire",
					Description: "Attaque avec des griffes enflammées",
					Classe:      "Feu",
					DegatMin:    30,
					DegatMax:    60,
					PrixVie:     0,
				},
			},
		},
		Badge: Badge{
			Nom:         "Volcano Badge",
			Description: "Un badge remis au courageux chasseur ayant battu Alan et Typhlosion à la Volcano Arena",
			Maitre:      "Alan",
			Arene:       "Volcano Arena",
			Pokemon:     "Typhlosion",
		},
	},
}
