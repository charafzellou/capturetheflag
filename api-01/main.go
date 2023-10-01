package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	originKey = "Pasting code from the Internet into production code is like chewing gum found in the street."
	// md5("Il n'y a que les imbéciles qui ne changent pas d'avis")
	// secretKey = "5bc2fb8cff6b14d9c62ea6447da62a4c"
	secretKey = "Das Einfügen von Code aus dem Internet in Produktionscode ist ..."
)

var (
	level        uint
	users        []string
	usersSecrets = make(map[string]string)
	usersPoints  = make(map[string]uint)
)

func main() {
	apiPort := genRandPort(1024, 8192)

	// write apiPort in log file
	writeLogFile(fmt.Sprintf("API listening on port: %s\n", apiPort))

	// expose route /ping to check if the API is up
	http.HandleFunc("/ping", handlerPing)

	// expose user-related routes
	http.HandleFunc("/signup", handlerSignup)
	http.HandleFunc("/check", handlerCheck)

	// expose level and point-related routes
	http.HandleFunc("/getUserSecret", handlerSecret)
	http.HandleFunc("/getUserLevel", handlerLevel)
	http.HandleFunc("/getUserPoints", handlerUserPoints)

	// expose challenge-related routes
	http.HandleFunc("/iNeedAHint", getHint)
	http.HandleFunc("/enterChallenge", getChallenge)
	http.HandleFunc("/submitSolution", submitChallenge)

	// expose the flag
	err := http.ListenAndServe(apiPort, nil)
	if err != nil {
		log.Fatalf("A big mess has occured : ", err)
	}
}
