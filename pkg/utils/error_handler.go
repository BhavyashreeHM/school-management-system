package utils

import (
	"log"
	"os"
)

func ErrorHandler(err error, message string) error {
	errorLogger :=  log.New(os.Stderr,"ERROR:",log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger.Printf(message,err)
	return err
}