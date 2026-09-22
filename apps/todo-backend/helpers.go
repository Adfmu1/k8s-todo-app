package main

import (
	"log/slog"
	"os"
)

func checkErr(err error, logger slog.Logger) {
	if err != nil {
		logger.Error(err.Error())
	}
}

func writeToFile(textToAppend string) {
	file, err := os.OpenFile("./todos.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	checkErr(err, *application.logger)
	defer file.Close()

	_, err = file.Write([]byte(textToAppend + "\n"))
	checkErr(err, *application.logger)
}
