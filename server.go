package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/gomail.v2"
)

type IncomingReq struct {
	Email string `json:"emailID"`
}

func handleMail(w http.ResponseWriter, r *http.Request) {
	// Read body
	reqBody, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var req IncomingReq
	err = json.Unmarshal(reqBody, &req)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	log.Println(req.Email)

	output, err := json.Marshal(req)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	sendMail(req)

	w.Header().Set("content-type", "application/json")
	w.Write(output)

}

func sendMail(req IncomingReq) {
	msg := gomail.NewMessage()
	msg.SetHeader("From", "rishi@email.com")
	msg.SetHeader("To", req.Email)
	msg.SetHeader("Subject", "Hello!")
	msg.SetBody("text/html", "Hello <b>Receiver</b>!")

	host := "smtp.mailtrap.io"
	port := 2525
	username := "e953cacdb56e85"
	password := "2ac7a9386b91de"

	daemon := gomail.NewDialer(host, port, username, password)

	// Send the email
	if err := daemon.DialAndSend(msg); err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/mail", handleMail)
	address := ":5252"
	log.Println("Starting server on address", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		panic(err)
	}
}
