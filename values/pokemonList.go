package values

var Pokemons []Pokemon = []Pokemon{
	{
		Nom:           "Bulbizarre",
		Classe:        "Plante",
		Description:   "Une plante puissante et résistante",
		Vie:           60,
		MaxVie:        60,
		ChanceCapture: 40,
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
		Vie:           60,
		MaxVie:        60,
		ChanceCapture: 40,
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
		Classe:        "Eau",
		Description:   "Un Pokémon d'eau calme et réfléchi",
		Vie:           60,
		MaxVie:        60,
		ChanceCapture: 40,
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
		ChanceCapture: 30,
		Attaques: []Attaque{
			{
				Nom:         "Éclair",
				Description: "Lance une décharge électrique sur l'adversaire",
				Classe:      "Électrique",
				DegatMin:    15,
				DegatMax:    35,
				PrixVie:     0,
			},
			{
				Nom:         "Coup de Foudre",
				Description: "Déclenche un coup de foudre puissant",
				Classe:      "Électrique",
				DegatMin:    20,
				DegatMax:    40,
				PrixVie:     0,
			},
			{
				Nom:         "Paralysie",
				Description: "Paralyse temporairement l'adversaire avec une charge électrique",
				Classe:      "Électrique",
				DegatMin:    5,
				DegatMax:    20,
				PrixVie:     -10,
			},
		},
	},
	{
		Nom:           "Rondoudou",
		Classe:        "Normal",
		Description:   "Un Pokémon rond et mignon qui aime chanter",
		Vie:           80,
		MaxVie:        80,
		ChanceCapture: 40,
		Attaques: []Attaque{
			{
				Nom:         "Mélodie",
				Description: "Chante une mélodie apaisante (ptdr non)",
				Classe:      "Normal",
				DegatMin:    5,
				DegatMax:    15,
				PrixVie:     -10,
			},
			{
				Nom:         "Choc Vocal",
				Description: "Émet un puissant choc vocal pour désorienter l'adversaire",
				Classe:      "Normal",
				DegatMin:    10,
				DegatMax:    30,
				PrixVie:     0,
			},
			{
				Nom:         "Câlin",
				Description: "Donne un câlin réconfortant pour restaurer la vie",
				Classe:      "Normal",
				DegatMin:    10,
				DegatMax:    20,
				PrixVie:     20,
			},
		},
	},
	{
		Nom:           "Dracaufeu",
		Classe:        "Feu",
		Description:   "Un dragon de feu impressionnant",
		Vie:           80,
		MaxVie:        80,
		ChanceCapture: 10,
		Attaques: []Attaque{
			{
				Nom:         "Déflagration",
				Description: "Déclenche une déflagration de feu intense",
				Classe:      "Feu",
				DegatMin:    20,
				DegatMax:    40,
				PrixVie:     0,
			},
			{
				Nom:         "Brûlure",
				Description: "Brûle l'adversaire avec une flamme continue",
				Classe:      "Feu",
				DegatMin:    10,
				DegatMax:    20,
				PrixVie:     0,
			},
			{
				Nom:         "Volcan",
				Description: "Crée une éruption volcanique dévastatrice",
				Classe:      "Feu",
				DegatMin:    30,
				DegatMax:    60,
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
		ChanceCapture: 5,
		Attaques: []Attaque{
			{
				Nom:         "Hydrocanon",
				Description: "Déchaîne un puissant jet d'eau sur l'adversaire",
				Classe:      "Eau",
				DegatMin:    25,
				DegatMax:    45,
				PrixVie:     0,
			},
			{
				Nom:         "Tsunami",
				Description: "Crée un tsunami dévastateur",
				Classe:      "Eau",
				DegatMin:    40,
				DegatMax:    80,
				PrixVie:     0,
			},
			{
				Nom:         "Blizzard",
				Description: "Lance un blizzard glacé pour geler l'adversaire",
				Classe:      "Eau",
				DegatMin:    15,
				DegatMax:    30,
				PrixVie:     0,
			},
		},
	},
	/*{
		Nom:           "Goupix",
		Classe:        "Feu",
		Description:   "Un Pokémon renard de feu aux neuf queues",
		Vie:           70,
		MaxVie:        70,
		ChanceCapture: 20,
		Attaques: []Attaque{
			{
				Nom:         "Brasier",
				Description: "Embrase l'adversaire avec des flammes intenses",
				Classe:      "Feu",
				DegatMin:    15,
				DegatMax:    30,
				PrixVie:     0,
			},
			{
				Nom:         "Queue de Feu",
				Description: "Frappe l'adversaire avec une queue enflammée",
				Classe:      "Feu",
				DegatMin:    10,
				DegatMax:    25,
				PrixVie:     0,
			},
			{
				Nom:         "Danse des Flammes",
				Description: "Danse avec les flammes pour augmenter la puissance et brûle à distance",
				Classe:      "Feu",
				DegatMin:    5,
				DegatMax:    20,
				PrixVie:     -10,
			},
		},
	},*/
	{
		Nom:           "Gyrados",
		Classe:        "Eau",
		Description:   "Un Pokémon eau féroce et redoutable",
		Vie:           85,
		MaxVie:        85,
		ChanceCapture: 15,
		Attaques: []Attaque{
			{
				Nom:         "Ouragan",
				Description: "Crée un ouragan dévastateur",
				Classe:      "Eau",
				DegatMin:    20,
				DegatMax:    40,
				PrixVie:     0,
			},
			{
				Nom:         "Morsure Féroce",
				Description: "Mord l'adversaire avec une force redoutable",
				Classe:      "Eau",
				DegatMin:    15,
				DegatMax:    35,
				PrixVie:     0,
			},
			{
				Nom:         "Séisme Marin",
				Description: "Déclenche un séisme sous-marin pour engloutir l'adversaire",
				Classe:      "Eau",
				DegatMin:    30,
				DegatMax:    60,
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
		ChanceCapture: 2,
		Attaques: []Attaque{
			{
				Nom:         "Psycho Explosion",
				Description: "Libère une explosion psychique",
				Classe:      "Psychique",
				DegatMin:    30,
				DegatMax:    60,
				PrixVie:     0,
			},
			{
				Nom:         "Éclair Mental",
				Description: "Frappe l'adversaire avec un éclair mental dévastateur",
				Classe:      "Psychique",
				DegatMin:    20,
				DegatMax:    40,
				PrixVie:     0,
			},
			{
				Nom:         "Barrière Mentale",
				Description: "Crée une barrière mentale pour se protéger mais attaque par télépatie",
				Classe:      "Psychique",
				DegatMin:    30,
				DegatMax:    50,
				PrixVie:     -15,
			},
		},
	},
	{
		Nom:           "Ronflex",
		Classe:        "Normal",
		Description:   "Un Pokémon paresseux mais incroyablement résistant",
		Vie:           120,
		MaxVie:        120,
		ChanceCapture: 5,
		Attaques: []Attaque{
			{
				Nom:         "Écrasement",
				Description: "Écrase l'adversaire avec son énorme masse",
				Classe:      "Normal",
				DegatMin:    40,
				DegatMax:    80,
				PrixVie:     0,
			},
			{
				Nom:         "Riposte",
				Description: "Contre-attaque avec une puissante riposte",
				Classe:      "Normal",
				DegatMin:    20,
				DegatMax:    40,
				PrixVie:     0,
			},
			{
				Nom:         "Ronflement",
				Description: "Émet un ronflement assourdissant pour endormir l'adversaire",
				Classe:      "Normal",
				DegatMin:    10,
				DegatMax:    25,
				PrixVie:     -10,
			},
		},
	},
}
