package database

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var dbName = "test.db"
var db *sql.DB
var tableNames = []string{"users","films","people","planets","species","starships","vehicles"}
//you must start the database , or you can not use the method
func Start(){
	var err error
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/testdb")
	checkErr(err)
}

func Stop(){
	db.Close()
}

func Init(){
	for i := 0 ; i < len(tableNames) ; i++{
		prepareSentence := fmt.Sprintf("create table if not exists %s(id varchar(10), content varchar(10000) )",tableNames[i])
		_ , err := db.Exec(prepareSentence)
		checkErr(err)

	}
}

func DeleteAllTable(){
	for i := 0 ; i < len(tableNames) ; i++{
		prepareSentence := fmt.Sprintf("drop table if exists %s",tableNames[i])
		_ , err := db.Exec(prepareSentence)
		checkErr(err)

	}
}

func Update(tableName string, key string, value string) {
	prepareSentence := fmt.Sprintf("replace into %s set id=?,content=?",tableName)

	stmt, err := db.Prepare(prepareSentence)
	checkErr(err)
	defer stmt.Close()
	_, err = stmt.Exec(key , value)
	
	checkErr(err)
}

func GetValue( tableName string, key string) string {
	prepareSentence := fmt.Sprintf("select content from %s where id=\"%s\"",tableName,key)

	var content string
	rows, err := db.Query(prepareSentence)
	defer rows.Close()
	checkErr(err)
	for rows.Next() {
		
		err = rows.Scan(&content)
	}

	return content
}


func CheckKeyExist(tableName string , key string) bool {
	prepareSentence := fmt.Sprintf("select count(content) from %s where id=\"%s\"",tableName,key)

	var count int
	rows, err := db.Query(prepareSentence)
	defer rows.Close()
	checkErr(err)
	for rows.Next() {
		err = rows.Scan(&count)
	}
	return count != 0
}

func GetBucketCount(tableName string) (int) {
	prepareSentence := fmt.Sprintf("select count(*) from %s",tableName)

	var count int
	rows, err := db.Query(prepareSentence)
	defer rows.Close()
	checkErr(err)
	for rows.Next() {
		err = rows.Scan(&count)
		checkErr(err)
	}
	return count
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
