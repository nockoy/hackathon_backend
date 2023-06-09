package controller

import (
	"db/model"
	"db/usecase"
	"encoding/json"
	"log"
	"net/http"
	"unicode/utf8"
)

func GetChannelByChannelID(w http.ResponseWriter, r *http.Request) {

	channelID := r.URL.Query().Get("channel_id")

	if channelID == "" {
		log.Println("fail: channelID is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bytes, err := usecase.GetChannelByChannelID(channelID)

	if err != nil {

		log.Printf("fail: , %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)

}

func GetUserChannelsByUserID(w http.ResponseWriter, r *http.Request) {

	userID := r.URL.Query().Get("user_id")

	if userID == "" {
		log.Println("fail: userID is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bytes, err := usecase.GetUserChannelsByUserID(userID)

	if err != nil {

		log.Printf("fail: , %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)

}

func GetOtherChannelsByUserID(w http.ResponseWriter, r *http.Request) {

	userID := r.URL.Query().Get("user_id")

	if userID == "" {
		log.Println("fail: userID is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bytes, err := usecase.GetOtherChannelsByUserID(userID)

	if err != nil {

		log.Printf("fail: , %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)

}

func RegisterChannel(w http.ResponseWriter, r *http.Request) {

	var channel model.Channels

	if err := json.NewDecoder(r.Body).Decode(&channel); err != nil {
		log.Println("fail: Error3")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if isOk := RegisterChannelCheck(channel.Name); isOk != true {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bytes, err := usecase.RegisterChannel(channel)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)

}

func RegisterChannelCheck(name string) bool {

	if name == "" {
		log.Println("fail: name is empty")
		return false
	}

	if utf8.RuneCountInString(name) > 16 {
		log.Println("fail: name length is over 16")
		return false
	}

	return true
}
