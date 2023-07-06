package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func getHint(w http.ResponseWriter, r *http.Request) {
	level++
	log.Printf("_______________\n")
	log.Printf("getHint called:\n")
	setHeaders(w, "POST")

	if r.Method == "POST" {
		printTraffic(r)

		var data AuthRequest
		// Parse the JSON body
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, "Failed to parse JSON body", http.StatusBadRequest)
			return
		}

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
						usersPoints[data.User]--
						usersPoints[data.User]--
						usersPoints[data.User]--
						var hint string
						rand := genRand(0, 4)
						switch rand {
						case 0:
							// hint := wikidata.org/wiki/Q13414952
							// hint := "Q 13 41 49 52"
							hint = "Q DC3 ) 1 4"
						case 1:
							// hint := https://store.steampowered.com/app/724490/Protocol/
							// hint := "App ID : 724490"
							hint = "App ID : 724490"
						case 2:
							// hint := "https://www.linternaute.fr/proverbe/cgi/recherche/recherche.php?f_motcle=imb%E9cile&lance_recherche.x=0&lance_recherche.y=0&f_action_recherche=0"
							// hint := "https://tinyurl.com/ctf-esgi-5A"
							hint = "Tiny ctf esgi 5A"
						case 3:
							// hint := "https://pastebin.com/cSfXHMvB"
							// hint := T75f91DQ2C
							// hint := "Struct FullRequest"
							hint = "CopyTrash cSfXHMvB T75f91DQ2C"
						}
						// convert the hash to a string
						fmt.Fprintf(w, "Coward over here asking for hints...\nnHere you go, your random hint:\n%s", hint)
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

func getChallenge(w http.ResponseWriter, r *http.Request) {
	level++
	log.Printf("_______________\n")
	log.Printf("getChallenge called:\n")
	setHeaders(w, "POST")

	if r.Method == "POST" {
		printTraffic(r)

		var data AuthRequest
		// Parse the JSON body
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, "Failed to parse JSON body", http.StatusBadRequest)
			return
		}

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
						usersPoints[data.User]--
						// calculate SHA1 of "data.User + usersSecrets[data.User] + usersPoints[data.User]"
						hash := sha1.New()
						hash.Write([]byte(data.User + usersSecrets[data.User] + fmt.Sprint(usersPoints[data.User])))
						// convert the hash to a string
						challenge := fmt.Sprintf("%x", hash.Sum(nil))
						log.Printf("Welcome to the challenge !\nHere is your first Challenge:\n%s\nDon't forget that:%s\n", challenge, secretKey)
						fmt.Fprintf(w, "Welcome to the challenge !\nHere is your first Challenge:\n%s\nDon't forget that:%s\n", challenge, secretKey)
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

func submitChallenge(w http.ResponseWriter, r *http.Request) {
	level++
	log.Printf("_______________\n")
	log.Printf("submitChallenge called:\n")
	setHeaders(w, "POST")

	if r.Method == "POST" {
		printTraffic(r)

		var data FullRequest
		// Parse the JSON body
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, "Failed to parse JSON body", http.StatusBadRequest)
			return
		}

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
						usersPoints[data.User]--
						// Generating hash
						if data.Content.Level == level {
							if data.Content.Challenge.Username == data.User {
								if data.Content.Challenge.Secret == usersSecrets[data.User] {
									if data.Content.Challenge.Points == usersPoints[data.User] {
										if data.Content.Protocol == "SHA-1" {
											if data.Content.SecretKey == originKey {
												writeLogFile(fmt.Sprintf("User: %s\nPoints: %d\nLevel: %d\nProtocol: %s\nSecretKey: %s\n", data.User, usersPoints[data.User], data.Content.Level, data.Content.Protocol, data.Content.SecretKey))
												writeResultFile(fmt.Sprintf("User: %s\nPoints: %d\nLevel: %d\nSecretKey: %s\n", data.User, usersPoints[data.User], data.Content.Level, data.Content.SecretKey))
												fmt.Fprintf(w, "Congrats!\nYou have completed the challenge!\nSend this to your instructor:%s\n", "flag{w3lc0m3_t0_th3_h4all_0f_f4m3}")
												return
											} else {
												log.Printf("Secret key does not match: %s\n", data.Content.SecretKey)
												fmt.Fprintf(w, "Wrong secret key: %s\n", data.Content.SecretKey)
												return
											}
										} else {
											log.Printf("Protocol does not match: %s\n", data.Content.Protocol)
											fmt.Fprintf(w, "Wrong Protocol: %s\n", data.Content.Protocol)
											return
										}
									} else {
										log.Printf("Challenge Points does not match: %d\n", data.Content.Challenge.Points)
										fmt.Fprintf(w, "Wrong challenge points: %d\n", data.Content.Challenge.Points)
										return
									}
								} else {
									log.Printf("Challenge Secret does not match: %s\n", data.Content.Challenge.Secret)
									fmt.Fprintf(w, "Wrong challenge secret: %s\n", data.Content.Challenge.Secret)
									return
								}
							} else {
								log.Printf("Challenge Username does not match: %s\n", data.Content.Challenge.Username)
								fmt.Fprintf(w, "Wrong challenge username: %s\n", data.Content.Challenge.Username)
								return
							}
						} else {
							log.Printf("Level does not match: %d\n", data.Content.Level)
							fmt.Fprintf(w, "Wrong level: %d\n", data.Content.Level)
							return
						}
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
