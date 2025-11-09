package utils

import "log"

func Info(msg string) {
	log.Printf("ℹ️  %s\n", msg)
}

func Error(err error) {
	if err != nil {
		log.Printf("❌ %v\n", err)
	}
}
