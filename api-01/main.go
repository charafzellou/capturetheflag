package main

import (
	"fmt"
	"net/http"
)

const (
	apiPort = ":3369"
	// sha256("Il n'y a que les imbéciles qui ne changent pas d'avis")
	secretKey = "ca32652906af8dd747e741cd3e960338138099b0615e62b4f23366cf65f52646"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method == "GET" {
			fmt.Printf("r.Host: %s", r.Host)
			fmt.Fprintf(w, "The secret key is: %s", secretKey)
		} else {
			fmt.Println("r.Method: ", r.Method)
			panic("Only GET requests are accepted")
		}
	})

	err := http.ListenAndServe(apiPort, nil)
	if err != nil {
		panic(err)
	}
}
