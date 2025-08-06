package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

var baseURL = "http://localhost:8080"
var getTime = func() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
var randomNum = func() int {
	return rand.Intn(100)
}

func main() {
	fmt.Println("Apa")
	getPID()

	http.HandleFunc("/time", getTimeHandler)
	http.HandleFunc("/error", errorHandler)
	http.HandleFunc("/random", getRandomHandler)
	http.HandleFunc("/ip", getIPHandler)
	http.ListenAndServe(":8080", nil)
}

func getPID() {
	cmd := exec.Command("pidof", "-s", "http-server")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func getRandomHandler(w http.ResponseWriter, r *http.Request) {
	randomNum := randomNum()
	response := map[string]int{"number": randomNum}
	json.NewEncoder(w).Encode(response)
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Something went wrong", http.StatusInternalServerError)
	fmt.Fprintln(os.Stderr, "Internal Server Error occurred at /error")
}

func getTimeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Current Time: %s", getTime())
}

func getClientIP(r *http.Request) string {
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {

		ips := strings.Split(xff, ",")
		return strings.TrimSpace(ips[0])
	}

	ipPort := r.RemoteAddr
	if colon := strings.LastIndexByte(ipPort, ':'); colon != -1 {
		return ipPort[:colon]
	}
	return ipPort
}

func getIPHandler(w http.ResponseWriter, r *http.Request) {
	ip := getClientIP(r)
	fmt.Fprintf(w, "Client IP: %s", ip)
}
