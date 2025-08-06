package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var config Config
var logFile *os.File
var getTime = func() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
var randomNum = func() int {
	return rand.Intn(100)
}

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

func (r *StatusRecorder) WriteHeader(code int) {
	r.statusCode = code
	r.ResponseWriter.WriteHeader(code)
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

func openLogFile() error {
	if logFile != nil {
		logFile.Close()
	}

	file, err := os.OpenFile(config.LogFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	logFile = file
	return nil
}

func writeLog(message string) {
	if logFile == nil {
		if err := openLogFile(); err != nil {
			log.Printf("Error opening log file: %v", err)
			return
		}
	}

	_, err := logFile.WriteString(message + "\n")
	if err != nil {
		log.Printf("Error writing to log file: %v", err)
	}
}

func logRequest(level int, message string) {
	if level <= config.LogLevel {
		writeLog(message)
	}
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

func main() {
	loadConfig()
	getPID()

	srv := &http.Server{
		Addr: ":8080",
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1, syscall.SIGUSR2)

	http.HandleFunc("/time", loggingMiddleware(getTimeHandler))
	http.HandleFunc("/error", loggingMiddleware(errorHandler))
	http.HandleFunc("/random", loggingMiddleware(getRandomHandler))

	go srv.ListenAndServe()
	for sig := range sigChan {
		switch sig {
		case syscall.SIGINT, syscall.SIGTERM:
			ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
			defer stop()
			log.Printf("got interruption signal")
			if err := srv.Shutdown(ctx); err != nil {
				log.Println("graceful shutdown failed:", err)
			}
			return
		case syscall.SIGUSR1:
			log.Println("Received SIGUSR1: Reopening log file...")
			if err := openLogFile(); err != nil {
				log.Printf("Failed to reopen log file: %v", err)
			}
		case syscall.SIGUSR2:
			log.Println("got SIGUSR2 signal")
		}
	}

	// <-ctx.Done()
	// log.Println("got interruption signal")
	// if err := srv.Shutdown(context.TODO()); err != nil {
	// 	log.Println("graceful shutdown failed:", err)
	// }
}
