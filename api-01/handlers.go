package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func handlerPing(w http.ResponseWriter, r *http.Request) {
	level++
	log.Printf("_______________\n")
	log.Printf("handlerPing called:\n")
	setHeaders(w, "GET")

	if r.Method == "GET" {
		printTraffic(r)
		fmt.Fprintf(w, "pong")
		return
	} else {
		printTraffic(r)
		fmt.Fprint(w, "Only GET requests are accepted")
		return
	}
}

func handlerSignup(w http.ResponseWriter, r *http.Request) {
	level++
	log.Printf("_______________\n")
	log.Printf("handlerSignup called:\n")
	setHeaders(w, "POST")

	if r.Method == "POST" {
		printTraffic(r)

		var data Request
		// Parse the JSON body
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, "Failed to parse JSON body", http.StatusBadRequest)
			return
		}
		log.Printf("Username used : %s\n", data.User)

		if data.User == "" {
			log.Printf("User name is empty\n")
			http.Error(w, "User name is empty", http.StatusBadRequest)
			return
		} else if len(data.User) < 4 {
			log.Printf("User name is too short: %s\n", data.User)
			http.Error(w, "User name is too short", http.StatusBadRequest)
			return
		} else if len(data.User) > 12 {
			log.Printf("User name is too long: %s\n", data.User)
			http.Error(w, "User name is too long", http.StatusBadRequest)
			return
		} else {
			// Add the user to the array
			users = append(users, data.User)
			// set usersSecrets[data.User] as the key and hash of data.User as the value
			hashUser := sha256.Sum256([]byte(data.User))
			usersSecrets[data.User] = fmt.Sprintf("%x", hashUser)
			usersPoints[data.User] = 100
			// Respond with success message
			log.Printf("New user has been saved: %s\n", data.User)
			fmt.Fprintf(w, "All is good, your user has been saved: %s\n", data.User)
			return
		}
	} else {
		printTraffic(r)
		fmt.Fprint(w, "Only POST requests are accepted")
		return
	}
}

func handlerCheck(w http.ResponseWriter, r *http.Request) {
	level++
	log.Printf("_______________\n")
	log.Printf("handlerCheck called:\n")
	setHeaders(w, "POST")

	if r.Method == "POST" {
		printTraffic(r)

		var data Request
		// Parse the JSON body
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, "Failed to parse JSON body", http.StatusBadRequest)
			return
		}
		log.Printf("Username used : %s\n", data.User)

		// Check if the user is in the array
		if data.User == "" {
			log.Printf("User name is empty\n")
			http.Error(w, "User name is empty", http.StatusBadRequest)
			return
		} else if len(data.User) < 4 || len(data.User) > 12 {
			log.Printf("User name is too short or too long: %s\n", data.User)
			http.Error(w, "User name is too short or too long", http.StatusBadRequest)
			return
		} else {
			for _, user := range users {
				if user == data.User {
					log.Printf("User found: %s\n", data.User)
					fmt.Fprintf(w, "User found: %s\n", data.User)
					return
				}
			}
		}
	} else {
		printTraffic(r)
		fmt.Fprint(w, "Only POST requests are accepted")
		return
	}
}

func handlerSecret(w http.ResponseWriter, r *http.Request) {
	level++
	log.Printf("_______________\n")
	log.Printf("handlerSecret called:\n")
	setHeaders(w, "POST")

	if r.Method == "POST" {
		printTraffic(r)

		var data Request
		// Parse the JSON body
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, "Failed to parse JSON body", http.StatusBadRequest)
			return
		}
		log.Printf("Username used : %s\n", data.User)

		// Check if the user is in the array
		for _, user := range users {
			if user == data.User {
				log.Printf("User found: %s\n", data.User)
				// check if the user is in the map
				if _, ok := usersSecrets[data.User]; ok {
					log.Printf("User secret found: %s\n", usersSecrets[data.User])
					fmt.Fprintf(w, "User secret: %s\n", usersSecrets[data.User])
					return
				} else {
					log.Printf("User not found in the map: %s\n", data.User)
					fmt.Fprintf(w, "No secret is attached to this user : %s\n", data.User)
					return
				}
			}
		}
	} else {
		printTraffic(r)
		fmt.Fprint(w, "Only POST requests are accepted")
		return
	}
}

func handlerLevel(w http.ResponseWriter, r *http.Request) {
	log.Printf("_______________\n")
	log.Printf("handlerUserPoints called:\n")
	setHeaders(w, "POST")

	if r.Method == "POST" {
		printTraffic(r)

		var data AuthRequest
		// Parse the JSON body
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, "Failed to parse JSON body", http.StatusBadRequest)
			return
		}
		log.Printf("Username used : %s\n", data.User)

		// Check if the user is in the array
		for _, user := range users {
			if user == data.User {
				log.Printf("User found: %s\n", data.User)
				// check if the user is in the map
				if _, ok := usersSecrets[data.User]; ok {
					log.Printf("User secret found: %s\n", usersSecrets[data.User])
					// heck if the user secret matches the one in the map
					if usersSecrets[data.User] == data.Secret {
						log.Printf("User secret matches the one in the map: %s\n", data.Secret)
						fmt.Fprintf(w, "Level: %d\n", level)
						return
					} else {
						log.Printf("User secret does not match the one in the map: %s\n", data.Secret)
						fmt.Fprintf(w, "Wrong secret for this user: %s\n", data.User)
						return
					}
				} else {
					log.Printf("User not found in the map: %s\n", data.User)
					fmt.Fprintf(w, "No secret is attached to this user : %s\n", data.User)
					return
				}
			}
		}
	} else {
		printTraffic(r)
		fmt.Fprint(w, "Only POST requests are accepted")
		return
	}
}

func handlerUserPoints(w http.ResponseWriter, r *http.Request) {
	level++
	log.Printf("_______________\n")
	log.Printf("handlerUserPoints called:\n")
	setHeaders(w, "POST")

	if r.Method == "POST" {
		printTraffic(r)

		var data AuthRequest
		// Parse the JSON body
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, "Failed to parse JSON body", http.StatusBadRequest)
			return
		}
		log.Printf("Username used : %s\n", data.User)

		// Check if the user is in the array
		for _, user := range users {
			if user == data.User {
				log.Printf("User found: %s\n", data.User)
				// check if the user is in the map
				if _, ok := usersSecrets[data.User]; ok {
					log.Printf("User secret found: %s\n", usersSecrets[data.User])
					// check if the user secret matches the one in the map
					if usersSecrets[data.User] == data.Secret {
						log.Printf("User secret matches the one in the map: %s\n", data.Secret)
						fmt.Fprintf(w, "User points: %s\n%d\n", data.User, usersPoints[data.User])
						return
					} else {
						log.Printf("User secret does not match the one in the map: %s\n", data.Secret)
						fmt.Fprintf(w, "Wrong secret for this user: %s\n", data.User)
						return
					}
				} else {
					log.Printf("User not found in the map: %s\n", data.User)
					fmt.Fprintf(w, "No secret is attached to this user : %s\n", data.User)
					return
				}
			}
		}
	} else {
		printTraffic(r)
		fmt.Fprint(w, "Only POST requests are accepted")
	}
}
