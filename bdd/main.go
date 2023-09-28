package bdd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

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
