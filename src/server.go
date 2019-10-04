package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

type IncomingReq struct {
	Email   string `json:"emailID"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type Response struct {
	Message string `json:"message"`
}

var USERNAME, PASSWORD, HOST, PORT, FROM string
var IPORT int

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

	sendMail(req)

	response := Response{}
	response.Message = "Email has been successfully send to " + req.Email

	output, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

func sendMail(req IncomingReq) {
	msg := gomail.NewMessage()
	msg.SetHeader("From", FROM)
	msg.SetHeader("To", req.Email)
	msg.SetHeader("Subject", req.Subject)
	msg.SetBody("text/html", "<b>"+req.Body+"</b>!")

	daemon := gomail.NewDialer(HOST, IPORT, USERNAME, PASSWORD)

	// Send the email
	if err := daemon.DialAndSend(msg); err != nil {
		panic(err)
	}
}

func validateAndSetMailCredentials() {
	USERNAME = os.Getenv("USERNAME")
	PASSWORD = os.Getenv("PASSWORD")
	HOST = os.Getenv("HOST")
	PORT = os.Getenv("PORT")
	FROM = os.Getenv("FROM")

	if USERNAME == "" || PASSWORD == "" || HOST == "" || PORT == "" || FROM == "" {
		panic("Mail server USERNAME/PASSWORD/HOST/PORT/FROM cannot be empty")
	}

	intValue, err := strconv.Atoi(PORT)
	if err != nil {
		panic("Mail server PORT value is invalid")
	} else {
		IPORT = intValue
	}
}

func main() {
	validateAndSetMailCredentials()

	http.HandleFunc("/mail", handleMail)
	address := ":5252"
	log.Println("Starting server on address", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		panic(err)
	}
}
