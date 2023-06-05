package main

import (
	"db/controller"
	"db/dao"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func init() {
	dao.DBInit()
}

func main() {

	http.HandleFunc("/user", controller.UserHandler)
	http.HandleFunc("/user/id", controller.UserIDHandler)
	http.HandleFunc("/message", controller.MessageHandler)
	http.HandleFunc("/channel", controller.ChannelIDHandler)
	http.HandleFunc("/channel/join", controller.UserChannelHandler)
	http.HandleFunc("/channel/notjoin", controller.OtherChannelHandler)

	// ③ Ctrl+CでHTTPサーバー停止時にDBをクローズする
	dao.DBClose()

	// 8000番ポートでリクエストを待ち受ける
	log.Println("Listening...")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}