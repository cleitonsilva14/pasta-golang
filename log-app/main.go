package main 

import (
	"log"
	"os"
	"time"
)

const (
	TYPE= "sqlserver"
	SERVER = "localhost"
	PORT = 1433
)

var currentTime = time.Now()

var (
	YEAR = currentTime.Year()
	MONTH = currentTime.Month()
	DAY = currentTime.Day()
)

func main(){

	file, err := openLogFile("./log1.log")
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("log file created")
	log.Printf("Connected on SERVER %s://%s:%d\n", TYPE, SERVER, PORT)
	for i := 0; i < 12; i++ {
		log.Println(currentTime.Month()-9)
	}
}


func openLogFile(path string) (*os.File, error) {
	logFile, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY , 0644)
	if err != nil {
		return nil, err
	}
	return logFile, nil
}


