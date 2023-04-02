package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DbConnection *sql.DB

type Person struct {
	Name string
	Age  int
}

func main() {
	DbConnection, _ := sql.Open("sqlite3", "./example.sql")
	defer DbConnection.Close()
	//cmd := `CREATE TABLE IF NOT EXISTS persons(
	//            name STRING,
	//			age  INT)`
	//_, err := DbConnection.Exec(cmd)
	//if err != nil {
	//	log.Fatalln(err)
	//}

	//cmd := "INSERT INTO persons (name, age) VALUES (?, ?)"
	//_, err := DbConnection.Exec(cmd, "Nancy", 20)
	//if err != nil {
	//	log.Fatalln(err)
	//}

	//cmd := "UPDATE persons SET age = ? WHERE name = ?"
	//_, err := DbConnection.Exec(cmd, 25, "Nancy")
	//if err != nil {
	//	log.Fatalln(err)
	//}

	//cmd := "SELECT * FROM persons where age = ?"
	//row := DbConnection.QueryRow(cmd, 25)
	//var p Person
	//err := row.Scan(&p.Name, &p.Age)
	//if err != nil {
	//	if err == sql.ErrNoRows {
	//		log.Println("No row")
	//	} else {
	//		log.Println(err)
	//	}
	//}
	//fmt.Println(p.Name, p.Age)

	cmd := "DELETE FROM persons WHERE name = ?"
	_, err := DbConnection.Exec(cmd, "Nancy")
	if err != nil {
		log.Fatalln(err)
	}
}
