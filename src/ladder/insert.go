package ladder

import "database/sql"

func InsertPlayer(tx *sql.Tx, player Player) error {
	stmt, err := tx.Prepare("insert into players(name, rank, aga_id) values (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(player.Name, player.AgaId, player.Rank)
	if err != nil {
		return err
	}
	return nil
}
