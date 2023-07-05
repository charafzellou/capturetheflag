package main

import (
	"fmt"
	"net/http"
)

const (
	apiPort      = ":3610"
	apiSecretKey = "8116fdd3f12b6d7c4b136cbdaa3360a57eb4eb676ae63294450ee1f4f34b36f3"
	// sha256("0x764aeebcf425d56800ef2c84f2578689415a2daa")
	// https://etherscan.io/address/0x764aeebcf425d56800ef2c84f2578689415a2daa
	finalResult = "22bec8b0f5a4318f6f88bd2fb0b0e04081c887e105574838d639943e254e521f"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method == "POST" {
			err := r.ParseForm()
			if err != nil {
				panic(err)
			}
			receivedString := r.FormValue("finalKey")
			if receivedString == apiSecretKey {
				fmt.Println("Received correct string!")
				fmt.Fprintf(w, "This is what you are looking for : %s", finalResult)
			} else {
				fmt.Printf("This string is not the correct one: %s", receivedString)
				fmt.Fprintf(w, "This string is not the correct one: %s", receivedString)
			}
		} else {
			fmt.Println("r.Method: ", r.Method)
			panic("Only POST requests are accepted")
		}
	})

	err := http.ListenAndServe(apiPort, nil)
	if err != nil {
		panic(err)
	}
}
