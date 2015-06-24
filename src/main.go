package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"ladder"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	os.Remove("./ladder.db")

	db, err := sql.Open("sqlite3", "./ladder.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createTablesBytes, err := ioutil.ReadFile("dbsetup.sql")
	createTablesSql := string(createTablesBytes)
	_, err = db.Exec(createTablesSql)
	if err != nil {
		log.Printf("%q: %s\n", err, createTablesSql)
		return
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T\n", tx)
	stmt, err := tx.Prepare("insert into players(name, rank, aga_id) values (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec("Quinten Palmer", 10, 26262626)
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec("Andrew Hall", -5, 5678)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()

	rows, err := db.Query("select * from players")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var rank float64
		var aga_id int
		rows.Scan(&id, &name, &rank, &aga_id)
		player := ladder.Player{
			Id:    id,
			Name:  name,
			Rank:  rank,
			AgaId: aga_id,
		}
		fmt.Println(player)
	}

	/*


		rows, err := db.Query("select id, name from foo")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		stmt, err = db.Prepare("select name from foo where id = ?")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()
		var name string
		err = stmt.QueryRow("3").Scan(&name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(name)

		_, err = db.Exec("delete from foo")
		if err != nil {
			log.Fatal(err)
		}

		_, err = db.Exec("insert into foo(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz')")
		if err != nil {
			log.Fatal(err)
		}

		rows, err = db.Query("select id, name from foo")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			var id int
			var name string
			rows.Scan(&id, &name)
			fmt.Println(id, name)
		}
	*/
}
