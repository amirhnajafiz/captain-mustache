package logger

import (
	"log"
	"time"
)

type Logger struct {
	Mode Mode
}

func (l Logger) Error(error error) {
	log.Printf("[ERROR] %s\n\terror: %v\n", time.Now().Format(time.DateTime), error)
}

func (l Logger) Info(message string) {
	if l.Mode == ProductMode || l.Mode == DebugMode {
		log.Printf("[INFO] %s\n\tmessage: %s\n", time.Now().Format(time.DateTime), message)
	}
}

func (l Logger) Debug(data string) {
	if l.Mode == DebugMode {
		log.Printf("[DEBUG] %s\n\tdata: %s\n", time.Now().Format(time.DateTime), data)
	}
}
