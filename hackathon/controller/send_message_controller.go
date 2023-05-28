package controller

import (
	"db/model"
	"db/usecase"
	"encoding/json"
	"log"
	"net/http"
	"unicode/utf8"

	_ "github.com/go-sql-driver/mysql"
)

type MSGResForHTTPPost struct { //使ってない
	RoomID string `json:"room_id"`
	From   int    `json:"from"`
	Text   int    `json:"text"`
}

func SendMessage(w http.ResponseWriter, r *http.Request) {

	var m model.Messages

	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		log.Println("fail: Error2")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if isOk := SendMessageCheck(m.Text); isOk != true {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bytes, err := usecase.SendMessage(m)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)

}

func SendMessageCheck(text string) bool {

	if text == "" {
		log.Println("fail: text is empty")
		return false
	}

	if utf8.RuneCountInString(text) > 500 {
		log.Println("fail: name length is over 500")
		return false
	}

	return true
}