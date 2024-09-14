package utils

import (
	"log"
	"time"
)

var logChannel = make(chan string)

func init() {
	go processLog()
}

func processLog() {
	for logMessage := range logChannel {
		log.Println(logMessage)
	}
}

// LogRequest logs the duration of a request.
func LogRequest(start time.Time, endpoint string) {
	duration := time.Since(start)
	logChannel <- "Request to " + endpoint + " took " + duration.String()
}

// LogInfo logs an informational message.
func LogInfo(message string) {
	logChannel <- "INFO: " + message
}

// LogError logs an error message.
func LogError(message string) {
	logChannel <- "ERROR: " + message
}

// LogSuccess logs a success message.
func LogSuccess(message string) {
	logChannel <- "SUCCESS: " + message
}
