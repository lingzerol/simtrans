package service

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
)

func InitLog() {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal("Failed to get the path of exectutable path: ", err)
		os.Exit(1)
	}
	filePattern := filepath.Join(filepath.Dir(exePath), "simtrans.%s.log")
	today := time.Now().Format("2006-01-02")
	logFileName := fmt.Sprintf(filePattern, today)
	file, err := os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
		os.Exit(1)
	}
	defer file.Close()
	log.SetOutput(file)

	logger.SetLevel(logrus.InfoLevel)
	for {
		time.Sleep(time.Hour)

		currentDate := time.Now().Format("2006-01-02")
		if currentDate != today {
			today = currentDate
			logFileName = fmt.Sprintf(filePattern, today)
			file.Close()
			file, err = os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
			if err != nil {
				log.Fatal("Failed to open log file:", err)
				os.Exit(1)
			}
			log.SetOutput(file)
		}
	}
}
