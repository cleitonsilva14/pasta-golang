package main 

import (
	"os"
	"log"
	"time"
	"fmt"
)

var (
	InfoLogger *log.Logger
	WarningLogger *log.Logger
	ErrorLogger *log.Logger
)

const (
	TYPE = "sqlserver"
	SERVER = "localhost"
	PORT= 1433
	USER = "sa"
	PASSWORD = "sa@1234"
)

func init(){

	file, err := os.OpenFile("logMain3.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	
	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}


func main(){

	connection := fmt.Sprintf("%s://%s:%d", TYPE, SERVER,PORT)

	InfoLogger.Println("Starting the application...")
	time.Sleep(2 * time.Second)
	InfoLogger.Println("Tentando conectar em " + connection)
	time.Sleep(2 * time.Second)

	WarningLogger.Println("Database not available...")
	time.Sleep(2 * time.Second)
	ErrorLogger.Println("Error on start program..")
	time.Sleep(2 * time.Second)
	InfoLogger.Println("Saindo...")

}


