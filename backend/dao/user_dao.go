package dao

import (
	"db/model"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func SearchUserByEmail(email string) ([]model.Users, error) {

	rows, err := db.Query("SELECT id, name, icon FROM users WHERE email = ?", email)

	users := make([]model.Users, 0)

	for rows.Next() {
		var u model.Users
		if ServerErr := rows.Scan(&u.ID, &u.Name, &u.Icon); ServerErr != nil {
			log.Printf("fail: rows.Scan, %v\n", err)

			if ServerErr := rows.Close(); ServerErr != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
				log.Printf("fail: rows.Close(), %v\n", err)
			}
			return nil, ServerErr
		}
		users = append(users, u)
	}

	return users, err
}

func SearchUserByUserID(UserID string) ([]model.Users, error) {

	rows, err := db.Query("SELECT id, name, icon FROM users WHERE id = ?", UserID)

	users := make([]model.Users, 0)

	for rows.Next() {
		var u model.Users
		if ServerErr := rows.Scan(&u.ID, &u.Name, &u.Icon); ServerErr != nil {
			log.Printf("fail: rows.Scan, %v\n", err)

			if ServerErr := rows.Close(); ServerErr != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
				log.Printf("fail: rows.Close(), %v\n", err)
			}
			return nil, ServerErr
		}
		users = append(users, u)
	}

	return users, err
}

func CreateUser(u model.Users) error {
	//トランザクション開始
	tx, err := db.Begin()
	if err != nil {
		log.Printf("fail: db.Begin, %v\n", err)
		return err
	}

	//INSERTする
	_, err = tx.Exec("INSERT INTO users(id, name, email, icon) values (?,?,?,?)", u.ID, u.Name, u.Email, u.Icon)
	if err != nil {
		log.Printf("fail: tx.Exec, %v\n", err)
		tx.Rollback()
		return err
	}

	//みんなgeneralチャンネルに入れる
	_, err = tx.Exec("INSERT INTO members(user_id, channel_id) values (?,?)", u.ID, "01H280H7SZ3JEHT0QRW8P658BF")
	if err != nil {
		log.Printf("fail: tx.Exec, %v\n", err)
		tx.Rollback()
		return err
	}
	//みんなrandomチャンネルに入れる
	_, err = tx.Exec("INSERT INTO members(user_id, channel_id) values (?,?)", u.ID, "01H28C7STNZH7P46P3XEF07WYD")
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

func UpdateIcon(u model.Users) error {
	//トランザクション開始
	tx, err := db.Begin()
	if err != nil {
		log.Printf("fail: db.Begin, %v\n", err)
		return err
	}

	//INSERTする
	_, err = tx.Exec("UPDATE users SET icon = ? WHERE id = ?", u.Icon, u.ID)
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

func UpdateUserName(u model.Users) error {
	//トランザクション開始
	tx, err := db.Begin()
	if err != nil {
		log.Printf("fail: db.Begin, %v\n", err)
		return err
	}

	//INSERTする
	_, err = tx.Exec("UPDATE users SET name = ? WHERE id = ?", u.Name, u.ID)
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

/*
func Delete(w http.ResponseWriter) {
	//トランザクション開始
	tx, err := db.Begin()
	if err != nil {
		log.Printf("fail: db.Begin, %v\n", err)
		return
	}

	//DELETEする
	_, err1 := tx.Exec("DELETE FROM user WHERE name = ?", "")
	if err1 != nil {
		log.Printf("fail: tx.Exec, %v\n", err1)
		w.WriteHeader(http.StatusInternalServerError)
		tx.Rollback()
		return
	}

	//トランザクション終了
	if err := tx.Commit(); err != nil {
		log.Printf("fail: tx.Commit, %v\n", err)
		return
	}
}
*/
