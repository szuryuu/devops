package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"time"
)

type Config struct {
	LogFileName string `json:"log_filename"`
	LogLevel    int    `json:"log_level"`
}

type HTTPLogger struct {
	Timestamp   time.Time `json:"timestamp"`
	Method      string    `json:"method"`
	Path        string    `json:"path"`
	StatusCode  int       `json:"status_code"`
	Duration    string    `json:"duration"`
	HTTPVersion string    `json:"http_version"`
	ClientIP    string    `json:"client_ip"`
}

type StatusRecorder struct {
	http.ResponseWriter
	statusCode int
}

var config Config
var getTime = func() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
var randomNum = func() int {
	return rand.Intn(100)
}

func main() {
	loadConfig()
	getPID()

	runtime.GOMAXPROCS(2)

	// srv := http.Server{}

	// ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	// defer stop()

	http.HandleFunc("/time", loggingMiddleware(getTimeHandler))
	http.HandleFunc("/error", loggingMiddleware(errorHandler))
	http.HandleFunc("/random", loggingMiddleware(getRandomHandler))

	http.ListenAndServe(":8080", nil)

	// <-ctx.Done()
	// log.Println("got interruption signal")
	// if err := srv.Shutdown(context.TODO()); err != nil {
	// 	log.Println("graceful shutdown failed:", err)
	// }
}

func loadConfig() {
	dirPath := "../app/config/config.json"

	file, err := os.Open(dirPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&config); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(config)
	}

}

func getPID() {
	pid := os.Getpid()
	fmt.Println("PID:", pid)
}

func getTimeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Current Time: %s", getTime())
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

func writeLog(message string) {
	filename := config.LogFileName

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(message + "\n")
	if err != nil {
		log.Fatal(err)
	}
}

func logRequest(level int, message string) {
	if level <= config.LogLevel {
		writeLog(message)
	}
}

func (r *StatusRecorder) WriteHeader(code int) {
	r.statusCode = code
	r.ResponseWriter.WriteHeader(code)
}

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rec := &StatusRecorder{ResponseWriter: w, statusCode: http.StatusOK}

		next(rec, r)

		duration := time.Since(start)
		clientIP := r.RemoteAddr

		logData := HTTPLogger{
			Timestamp:   time.Now(),
			Method:      r.Method,
			Path:        r.URL.Path,
			StatusCode:  rec.statusCode,
			Duration:    duration.String(),
			HTTPVersion: r.Proto,
			ClientIP:    clientIP,
		}

		logBytes, _ := json.Marshal(logData)
		logRequest(2, string(logBytes))
	}
}
