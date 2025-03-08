package utils

import (
	"io"
	"log"
	"os"
)

func newLogger(logFilePath string) *log.Logger {
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("failed to open log file: %v", err)
	}

	multiWriter := io.MultiWriter(os.Stdout, logFile)

	return log.New(multiWriter, "[MAIN		] ", log.Ldate|log.Ltime|log.Lshortfile)
}

var Log = newLogger("main.log")
