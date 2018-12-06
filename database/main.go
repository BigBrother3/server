package main

import (
	"encoding/json"

	"log"
	"strconv"

	"github.com/BigBrother3/server/database/database"
	"github.com/BigBrother3/server/database/swapi"
)

func main() {
	reSetDB := true
	database.SetDbName("./database/test.db")
	if reSetDB {
		reSetDatabase()
	}

	//example
	//1-7
	value := database.GetValue([]byte("film"), []byte("1"))
	//1-17
	//value := database.GetValue([]byte("person") , []byte("1"))
	//1-61
	//value := database.GetValue([]byte("planet") , []byte("1"))
	//1-37
	//value := database.GetValue([]byte("species") , []byte("1"))
	//??
	//value := database.GetValue([]byte("starship") , []byte("1"))
	//value := database.GetValue([]byte("vehicle") , []byte("1"))
	log.Println(value)
}

func reSetDatabase() {
	//database.DeleteDB()
	//database.Init()
	/*findFilm()
	findPerson()
	findPlanet()
	findSpecies()*/

	findVehicle()
	findStarship()
}

func findFilm() {
	c := swapi.DefaultClient
	for index := 1; ; index++ {
		jsonStr := dump(c.Film(index))
		indexStr := strconv.Itoa(index)
		if len(database.GetValue([]byte("film"), []byte(indexStr))) == 0 {
			if !putIntoDb([]byte("film"), []byte(indexStr), jsonStr) {
				break
			}
		}
	}
}

func findPerson() {
	c := swapi.DefaultClient
	for index := 1; ; index++ {
		jsonStr := dump(c.Person(index))
		indexStr := strconv.Itoa(index)
		if len(database.GetValue([]byte("person"), []byte(indexStr))) == 0 {
			if !putIntoDb([]byte("person"), []byte(indexStr), jsonStr) {
				break
			}
		}
	}
}

func findPlanet() {
	c := swapi.DefaultClient
	for index := 1; ; index++ {
		jsonStr := dump(c.Planet(index))
		indexStr := strconv.Itoa(index)
		if len(database.GetValue([]byte("planet"), []byte(indexStr))) == 0 {
			if !putIntoDb([]byte("planet"), []byte(indexStr), jsonStr) {
				break
			}
		}
	}
}

func findSpecies() {
	c := swapi.DefaultClient
	for index := 1; ; index++ {
		jsonStr := dump(c.Species(index))
		indexStr := strconv.Itoa(index)
		if len(database.GetValue([]byte("species"), []byte(indexStr))) == 0 {
			if !putIntoDb([]byte("species"), []byte(indexStr), jsonStr) {
				break
			}
		}
	}
}

func findStarship() {
	c := swapi.DefaultClient
	for index := 1; ; index++ {
		jsonStr := dump(c.Starship(index))
		indexStr := strconv.Itoa(index)
		if len(database.GetValue([]byte("starship"), []byte(indexStr))) == 0 {
			if !putIntoDb([]byte("starship"), []byte(indexStr), jsonStr) {
				break
			}
		}
	}
}

func findVehicle() {
	c := swapi.DefaultClient
	for index := 1; ; index++ {
		jsonStr := dump(c.Vehicle(index))
		indexStr := strconv.Itoa(index)
		if !putIntoDb([]byte("vehicle"), []byte(indexStr), jsonStr) {
			break
		}
	}
}

func dump(data interface{}, err error) []byte {
	jsonStr, err := json.MarshalIndent(data, "", "  ")
	return jsonStr
}

func putIntoDb(bucketName []byte, index []byte, jsonStr []byte) bool {
	stb := &swapi.Film{}
	err := json.Unmarshal(jsonStr, &stb)
	if err != nil {
		log.Fatal(err)
		return false
	} else if len(stb.URL) == 0 {
		return false
	}
	log.Printf("solve %s/%s", bucketName, index)
	database.Update(bucketName, index, jsonStr)
	return true

}
