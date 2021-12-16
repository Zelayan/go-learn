package nglog

import (
	"log/internal/request"
	"os"
)

func Download(service string) {


	logFile, err := os.Create(service + ".log")
	if err != nil {

	}
	defer logFile.Close()
	log := request.GetServiceLog(service)
	_, err = logFile.Write([]byte(log))
	if err != nil {

	}
}

func List() {
	request.Get()
}