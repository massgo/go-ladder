package ladder

import (
	"database/sql"
	"log"
)

func GetAllPlayers(tx *sql.Tx) ([]Player, error) {
	rows, err := tx.Query("select * from players")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	players := []Player{}
	for rows.Next() {
		var id int
		var name string
		var rank float64
		var aga_id int
		rows.Scan(&id, &name, &rank, &aga_id)
		player := Player{
			Id:    id,
			Name:  name,
			Rank:  rank,
			AgaId: aga_id,
		}
		players = append(players, player)
	}
	return players, nil
}
