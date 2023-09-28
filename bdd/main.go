package bdd

import (
	"encoding/json"
	"fmt"
	"game/values"
	"io/ioutil"
	"strings"
)

var Database *QuickDB

type QuickDB struct {
	filePath string
	data     map[string]interface{}
}

func NewQuickDB(filePath string) *QuickDB {
	db := &QuickDB{
		filePath: filePath,
		data:     make(map[string]interface{}),
	}
	db.loadData()
	return db
}

func (db *QuickDB) loadData() {
	file, err := ioutil.ReadFile(db.filePath)
	if err != nil {
		return
	}

	if err := json.Unmarshal(file, &db.data); err != nil {
		fmt.Println("Error loading data:", err)
	}
}

func (db *QuickDB) saveData() {
	dataJSON, err := json.MarshalIndent(db.data, "", "  ")
	if err != nil {
		fmt.Println("Error saving data:", err)
		return
	}

	err = ioutil.WriteFile(db.filePath, dataJSON, 0644)
	if err != nil {
		fmt.Println("Error saving data to file:", err)
	}
}

func (db *QuickDB) Set(key string, value interface{}) {
	db.data[key] = value
	db.saveData()
}

func (db *QuickDB) Get(key string) interface{} {
	return db.data[key]
}

func (db *QuickDB) Delete(key string) {
	delete(db.data, key)
	db.saveData()
}

func (db *QuickDB) GetAll() map[string]interface{} {
	return db.data
}

func (db *QuickDB) AddPokemon(name string) {
	if db.Get("owned_pokemon") != nil {
		old := db.Get("owned_pokemon").(string)
		db.Set("owned_pokemon", old+"_"+name)
	} else {
		db.Set("owned_pokemon", name)
	}
	db.SavePokemons(values.MainCharacter.Pokemons)
}

func (db *QuickDB) RemovePokemon(name string) {
	if db.Get("owned_pokemon") != nil {
		old := db.Get("owned_pokemon").(string)
		old = strings.Replace(old, name, "", -1)
		db.Set("owned_pokemon", old)
	}
	db.SavePokemons(values.MainCharacter.Pokemons)
}

func (db *QuickDB) GetPokemon() []string {
	if db.Get("owned_pokemon") != nil {
		all := db.Get("owned_pokemon").(string)
		splitted := strings.Split(all, "_")
		var result []string
		for _, c := range splitted {
			if c != "" {
				result = append(result, c)
			} else {
				continue
			}
		}
		return result
	} else {
		return []string{}
	}
}

func (db *QuickDB) SavePokemons(pokemons []values.Pokemon) error {
	pokemonsJSON, err := json.Marshal(pokemons)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("pokemons.json", pokemonsJSON, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (db *QuickDB) LoadPokemons() error {
	jsonData, err := ioutil.ReadFile("pokemons.json")
	if err != nil {
		return err
	}

	var pokemons []values.Pokemon

	err = json.Unmarshal(jsonData, &pokemons)
	if err != nil {
		return err
	}

	values.MainCharacter.Pokemons = pokemons

	return nil
}
