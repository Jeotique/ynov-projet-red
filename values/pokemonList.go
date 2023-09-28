package values

var Pokemons []Pokemon = []Pokemon{
	{
		Nom:           "Bulbizarre",
		Classe:        "plante",
		Description:   "Une plante puissante et résistante",
		Vie:           60,
		MaxVie:        60,
		ChanceCapture: 90,
		Attaques: []Attaque{{
			Nom:         "Fouetard",
			Description: "Fouete l'adversaire avec une liane",
			Classe:      "plante",
			DegatMin:    10,
			DegatMax:    30,
			PrixVie:     0,
		}},
	},
	{
		Nom:           "Salamèche",
		Classe:        "Feu",
		Description:   "Un Pokémon de feu fougueux et déterminé",
		Vie:           30,
		MaxVie:        60,
		ChanceCapture: 90,
		Attaques: []Attaque{{
			Nom:         "Boule de feu",
			Description: "Envoie une boule de feu sur l'adversaire",
			Classe:      "feu",
			DegatMin:    10,
			DegatMax:    30,
			PrixVie:     0,
		}},
	},
	{
		Nom:           "Carapuce",
		Classe:        "eau",
		Description:   "Un Pokémon d'eau calme et réfléchi",
		Vie:           60,
		MaxVie:        60,
		ChanceCapture: 90,
		Attaques: []Attaque{{
			Nom:         "Jet d'eau",
			Description: "Arrose d'un jet puissant l'adversaire",
			Classe:      "eau",
			DegatMin:    10,
			DegatMax:    30,
			PrixVie:     0,
		}},
	},
	{
		Nom:           "Pikachu",
		Classe:        "Électrique",
		Description:   "Un Pokémon électrique mignon mais puissant",
		Vie:           55,
		MaxVie:        55,
		ChanceCapture: 90,
		Attaques: []Attaque{
			{
				Nom:         "Éclair",
				Description: "Lance une décharge électrique sur l'adversaire",
				Classe:      "Électrique",
				DegatMin:    15,
				DegatMax:    35,
				PrixVie:     0,
			},
		},
	},
	{
		Nom:           "Rondoudou",
		Classe:        "Normal",
		Description:   "Un Pokémon rond et mignon qui aime chanter",
		Vie:           80,
		MaxVie:        80,
		ChanceCapture: 90,
		Attaques: []Attaque{
			{
				Nom:         "Mélodie",
				Description: "Chante une mélodie apaisante",
				Classe:      "Normal",
				DegatMin:    0,
				DegatMax:    0,
				PrixVie:     -10,
			},
		},
	},
	{
		Nom:           "Dracaufeu",
		Classe:        "Feu",
		Description:   "Un dragon de feu impressionnant",
		Vie:           80,
		MaxVie:        80,
		ChanceCapture: 80,
		Attaques: []Attaque{
			{
				Nom:         "Déflagration",
				Description: "Déclenche une déflagration de feu intense",
				Classe:      "Feu",
				DegatMin:    20,
				DegatMax:    40,
				PrixVie:     0,
			},
		},
	},
	{
		Nom:           "Léviator",
		Classe:        "Eau",
		Description:   "Un puissant Pokémon eau avec une peau dure",
		Vie:           90,
		MaxVie:        90,
		ChanceCapture: 70,
		Attaques: []Attaque{
			{
				Nom:         "Hydrocanon",
				Description: "Déchaîne un puissant jet d'eau sur l'adversaire",
				Classe:      "Eau",
				DegatMin:    25,
				DegatMax:    45,
				PrixVie:     0,
			},
		},
	},
	{
		Nom:           "Ptera",
		Classe:        "Vol",
		Description:   "Un Pokémon volant préhistorique",
		Vie:           70,
		MaxVie:        70,
		ChanceCapture: 80,
		Attaques: []Attaque{
			{
				Nom:         "Rafale de vent",
				Description: "Crée une rafale de vent tranchante",
				Classe:      "Vol",
				DegatMin:    15,
				DegatMax:    35,
				PrixVie:     0,
			},
		},
	},
	{
		Nom:           "Goupix",
		Classe:        "Feu",
		Description:   "Un Pokémon renard de feu aux neuf queues",
		Vie:           70,
		MaxVie:        70,
		ChanceCapture: 85,
		Attaques: []Attaque{
			{
				Nom:         "Brasier",
				Description: "Embrase l'adversaire avec des flammes intenses",
				Classe:      "Feu",
				DegatMin:    15,
				DegatMax:    30,
				PrixVie:     0,
			},
		},
	},
	{
		Nom:           "Gyrados",
		Classe:        "Eau",
		Description:   "Un Pokémon eau féroce et redoutable",
		Vie:           85,
		MaxVie:        85,
		ChanceCapture: 75,
		Attaques: []Attaque{
			{
				Nom:         "Ouragan",
				Description: "Crée un ouragan dévastateur",
				Classe:      "Eau",
				DegatMin:    20,
				DegatMax:    40,
				PrixVie:     0,
			},
		},
	},
	{
		Nom:           "Mewtwo",
		Classe:        "Psychique",
		Description:   "Un Pokémon psychique légendaire avec un pouvoir mental immense",
		Vie:           100,
		MaxVie:        100,
		ChanceCapture: 50,
		Attaques: []Attaque{
			{
				Nom:         "Psycho Explosion",
				Description: "Libère une explosion psychique",
				Classe:      "Psychique",
				DegatMin:    30,
				DegatMax:    60,
				PrixVie:     0,
			},
		},
	},
	{
		Nom:           "Ronflex",
		Classe:        "Normal",
		Description:   "Un Pokémon paresseux mais incroyablement résistant",
		Vie:           120,
		MaxVie:        120,
		ChanceCapture: 60,
		Attaques: []Attaque{
			{
				Nom:         "Écrasement",
				Description: "Écrase l'adversaire avec son énorme masse",
				Classe:      "Normal",
				DegatMin:    40,
				DegatMax:    80,
				PrixVie:     0,
			},
		},
	},
}
