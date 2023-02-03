package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method == "GET" {
			fmt.Printf("r.Host: %s", r.Host)
			fmt.Fprintf(w, "r.Host: %s", r.Host)
		} else {
			fmt.Println("r.Method: ", r.Method)
			panic("Only GET requests are accepted")
		}
	})

	err := http.ListenAndServe(":3200", nil)
	if err != nil {
		panic(err)
	}
}
