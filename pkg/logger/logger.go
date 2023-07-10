package logger

import (
	"log"
	"time"
)

func Error(error error) {
	log.Printf("[ERROR] %s | error: %v", time.Now().Format(time.DateTime), error)
}
