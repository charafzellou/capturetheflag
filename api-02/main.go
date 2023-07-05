package main

import (
	"fmt"
	"net/http"
)

const (
	apiPort      = ":3941"
	apiSecretKey = "ca32652906af8dd747e741cd3e960338138099b0615e62b4f23366cf65f52646"
	// below file contains API Three Port : 3610
	urlFile = "https://fromsmash.com/mrsii82xfO-ct"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method == "POST" {
			err := r.ParseForm()
			if err != nil {
				panic(err)
			}
			receivedString := r.FormValue("secretKey")
			if receivedString == apiSecretKey {
				fmt.Println("Received correct string!")
				fmt.Fprintf(w, "Download this file: %s", urlFile)
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
