package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func genRand(low int, high int) int {
	// Set the seed for the random number generator
	rand.New(rand.NewSource(time.Now().UnixNano()))
	// Generate a random number between 1024 and 65535
	randomNumber := rand.Intn(high-low) + low
	// Print the random number
	log.Printf("Time.now(): %s\n", time.Now().String())
	// log.Println("Random Number Generated: ", randomNumber)
	// Convert the random number to a string
	// and return it
	return randomNumber
}

func genRandPort(low int, high int) string {
	return fmt.Sprintf(":%d", genRand(low, high))
}

func printTraffic(r *http.Request) {
	log.Printf("r.RemoteAddr: %s\n", r.RemoteAddr)
	log.Printf("r.Method: %s\n", r.Method)
}

func writeLogFile(content string) {
	logFile, err := os.OpenFile("logs/api.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	fmt.Fprint(logFile, content)
	defer logFile.Close()
}

func writeResultFile(content string) {
	resultFile, err := os.OpenFile("logs/result.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	fmt.Fprint(resultFile, content)
	defer resultFile.Close()
}

func setHeaders(w http.ResponseWriter, allowedMethod string) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", allowedMethod)
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
}
