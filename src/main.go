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
	err = ladder.InsertPlayer(tx, ladder.Player{Name: "Quinten Palmer", AgaId: 10, Rank: 234})
	err = ladder.InsertPlayer(tx, ladder.Player{Name: "Andrew Hall", AgaId: -5, Rank: 34})
	tx.Commit()

	tx, err = db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	players, err := ladder.GetAllPlayers(tx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(players)
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
