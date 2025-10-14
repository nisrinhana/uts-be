package config

import (
	"log"
	"os"
)

func InitLogger() *log.Logger {
	file, _ := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	return log.New(file, "APP_LOG: ", log.Ldate|log.Ltime|log.Lshortfile)
}
