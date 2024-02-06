package main

import (
	"log"
	"os"
)

func main(){
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	logger.Println("Mensagem de info em log STDOUT.")
}





