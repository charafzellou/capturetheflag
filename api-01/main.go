package main

import (
	"fmt"
	"net/http"
)

const (
	originKey = "Il n'y a que les imbéciles qui ne changent pas d'avis."
	// sha1("Il n'y a que les imbéciles qui ne changent pas d'avis")
	secretKey = "800ab01c158d4271e7f366203666a6b7eb6e4535"
)

type (
	Request struct {
		User string `json:"User"`
	}
	AuthRequest struct {
		User   string `json:"User"`
		Secret string `json:"Secret"`
	}
	FullRequest struct {
		User    string `json:"User"`
		Secret  string `json:"Secret"`
		Content struct {
			Level     uint `json:"Level"`
			Challenge struct {
				Username string `json:"Username"`
				Secret   string `json:"Secret"`
				Points   uint   `json:"Points"`
			} `json:"Challenge"`
			Protocol  string `json:"Protocol"`
			SecretKey string `json:"SecretKey"`
		} `json:"Content"`
	}
)

var (
	level        uint
	users        []string
	usersSecrets = make(map[string]string)
	usersPoints  = make(map[string]uint)
)

func main() {
	apiPort := genRandPort(2048, 4096)

	// write apiPort in log file
	writeLogFile(fmt.Sprintf("API listening on port: %s\n", apiPort))

	http.HandleFunc("/ping", handlerPing)

	http.HandleFunc("/signup", handlerSignup)
	http.HandleFunc("/check", handlerCheck)
	http.HandleFunc("/secret", handlerSecret)

	http.HandleFunc("/getLevel", handlerLevel)
	http.HandleFunc("/getUserPoints", handlerUserPoints)

	http.HandleFunc("/getHint", getHint)
	http.HandleFunc("/getChallenge", getChallenge)
	http.HandleFunc("/submitChallenge", submitChallenge)

	err := http.ListenAndServe(apiPort, nil)
	if err != nil {
		panic(err)
	}
}
