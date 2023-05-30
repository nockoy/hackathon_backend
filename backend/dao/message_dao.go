package dao

import (
	"db/model"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func GetMessage(RoomID string) ([]model.Messages, error) {

	rows, err := db.Query("SELECT * FROM messages WHERE room_id = ?", RoomID)

	messages := make([]model.Messages, 0)

	for rows.Next() {
		var m model.Messages
		if ServerErr := rows.Scan(&m.ID, &m.ReplyToID, &m.RoomID, &m.UserID, &m.Text, &m.CreatedAt, &m.UpdatedAt); ServerErr != nil {
			log.Printf("fail: rows.Scan, %v\n", err)

			if ServerErr := rows.Close(); ServerErr != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
				log.Printf("fail: rows.Close(), %v\n", err)
			}
			return nil, ServerErr
		}
		messages = append(messages, m)
	}

	return messages, err
}

func CreateMSG(m model.Messages) error {
	//トランザクション開始
	tx, err := db.Begin()
	if err != nil {
		log.Printf("fail: db.Begin, %v\n", err)
		return err
	}

	//INSERTする
	_, err = tx.Exec("INSERT INTO messages(id, room_id, user_id, text, created_at, updated_at) values (?,?,?,?,?,?)", m.ID, m.RoomID, m.UserID, m.Text, m.CreatedAt, m.UpdatedAt)
	if err != nil {
		log.Printf("fail: tx.Exec, %v\n", err)
		tx.Rollback()
		return err
	}

	//トランザクション終了
	if err := tx.Commit(); err != nil {
		log.Printf("fail: tx.Commit, %v\n", err)
		return err
	}

	return nil
}