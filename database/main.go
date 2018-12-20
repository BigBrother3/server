package main

import (
	"encoding/json"

	"log"
	"strconv"

	"github.com/BigBrother3/server/database/database"
	"github.com/BigBrother3/server/database/swapi"
)

var dbName = "./database/test.db"

func main() {
	
	reSetDB := true
	database.Start()
	if reSetDB {
		//database.DeleteAllTable();
		database.Init()
		reSetDatabase()
	}else{
		//example
		//1-7
		//log.Println( database.GetValue(("films"), ("1")) )
		//87 some is invalid
		//log.Println( database.GetValue(("people") , ("1")) )
		//61
		//log.Println( database.GetValue([]byte("planets") , []byte("1")) )
		//37
		//log.Println( database.GetValue([]byte("species") , []byte("1")) )
		//39 some is invalid
		//log.Println(database.GetValue([]byte("starships") , []byte("10")) )
		//37 some is invalid
		//log.Println(database.GetValue([]byte("vehicles") , []byte("10")) )
		//log.Println(database.CheckKeyExist([]byte("vehicles") , []byte("8")) )
		//log.Println( database.GetValue([]byte("users") , []byte("vip") ) )

	}
}

func reSetDatabase() {
	findFilm()
	findPerson()
	findPlanet()
	findSpecies()
	findVehicle()
	findStarship()
}

func findFilm() {
	c := swapi.DefaultClient
	invalidTime := 0
	for index := 1; ; index++ {
		indexStr := strconv.Itoa(index)
		if !database.CheckKeyExist("films" , indexStr) {
			jsonStr := dump(c.Film(index))
			if !putIntoDb("films", indexStr, jsonStr) {
				invalidTime ++
				if invalidTime == 10{
					break
				}
			}else{
				invalidTime = 0
			}
		}else{
			invalidTime = 0
			log.Printf("films/%d is already exit", index)
		}
	}
}

func findPerson() {
	c := swapi.DefaultClient
	invalidTime := 0
	for index := 1; ; index++ {
		
		indexStr := strconv.Itoa(index)
		if !database.CheckKeyExist("people" , indexStr) {
			jsonStr := dump(c.Film(index))
			if !putIntoDb("people", indexStr, jsonStr) {
				invalidTime ++
				if invalidTime == 10{
					break
				}
			}else{
				invalidTime = 0
			}
		}else{
			invalidTime = 0
			log.Printf("person/%d is already exit", index)
		}
	}
}

func findPlanet() {
	c := swapi.DefaultClient
	invalidTime := 0
	for index := 1; ; index++ {
		
		indexStr := strconv.Itoa(index)
		if !database.CheckKeyExist("planets" , indexStr) {
			jsonStr := dump(c.Film(index))
			if !putIntoDb("planets", indexStr, jsonStr) {
				invalidTime ++
				if invalidTime == 10{
					break
				}
			}else{
				invalidTime = 0
			}
		}else{
			invalidTime = 0
			log.Printf("planets/%d is already exit", index)
		}
	}
}

func findSpecies() {
	c := swapi.DefaultClient
	invalidTime := 0
	for index := 1; ; index++ {
		
		indexStr := strconv.Itoa(index)
		if !database.CheckKeyExist("species" , indexStr) {
			jsonStr := dump(c.Film(index))
			if !putIntoDb("planets", indexStr, jsonStr) {
				invalidTime ++
				if invalidTime == 10{
					break
				}
			}else{
				invalidTime = 0
			}
		}else{
			invalidTime = 0
			log.Printf("species/%d is already exit", index)
		}
	}
}

func findStarship() {
	c := swapi.DefaultClient
	invalidTime := 0
	for index := 1; ; index++ {
		
		indexStr := strconv.Itoa(index)
		if !database.CheckKeyExist("starships" , indexStr) {
			jsonStr := dump(c.Film(index))
			if !putIntoDb("starships", indexStr, jsonStr) {
				invalidTime ++
				if invalidTime == 10{
					break
				}
			}else{
				invalidTime = 0
			}
		}else{
			invalidTime = 0
			log.Printf("starships/%d is already exit", index)
		}
	}
}

func findVehicle() {
	c := swapi.DefaultClient
	invalidTime := 0
	for index := 1; ; index++ {
		indexStr := strconv.Itoa(index)
		if !database.CheckKeyExist("vehicles" , indexStr) {
			jsonStr := dump(c.Film(index))
			if !putIntoDb("vehicles", indexStr, jsonStr) {
				invalidTime ++
				if invalidTime == 10{
					break
				}
			}else{
				invalidTime = 0
			}
		}else{
			invalidTime = 0
			log.Printf("vehicles/%d is already exit", index)
		}
	}
}

func dump(data interface{}, err error) []byte {
	jsonStr, err := json.MarshalIndent(data, "", "  ")
	return jsonStr
}

func putIntoDb(bucketName string, index string, jsonStr []byte) bool {
	stb := &swapi.Film{}
	err := json.Unmarshal(jsonStr, &stb)

	if err != nil {
		log.Fatal(err)
		return false
	} else if len(stb.URL) == 0 {
		log.Printf("%s/%s is invalid", bucketName, index)
		return false
	}
	log.Printf("solve %s/%s", bucketName, index)
	database.Update(bucketName, index, string(jsonStr))
	return true

}
