package dao

import (
	"db/model"
	"log"
)

func CreateWorkspace(workspace model.Workspaces) error {
	//トランザクション開始
	tx, err := db.Begin()
	if err != nil {
		log.Printf("fail: db.Begin, %v\n", err)
		return err
	}

	//INSERTする
	_, err = tx.Exec("INSERT INTO workspaces(id, name, created_at, updated_at) values (?,?,?,?)", workspace.ID, workspace.Name, workspace.CreatedAt, workspace.UpdatedAt)
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