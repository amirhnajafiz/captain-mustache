package logger

import (
	"log"
	"time"
)

func Error(error error) {
	log.Printf("[ERROR] %s\n\terror: %v\n", time.Now().Format(time.DateTime), error)
}

func Info(message string) {
	log.Printf("[INFO] %s\n\tmessage: %s\n", time.Now().Format(time.DateTime), message)
}

func Debug(data string) {
	log.Printf("[DEBUG] %s\n\tdata: %s\n", time.Now().Format(time.DateTime), data)
}
