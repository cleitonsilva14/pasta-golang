package main

import (
	"log"
	"os"
	"time"
	_ "strings"
	"fmt"
)

func main(){

	now := time.Now()

	year := now.Format("2006")
	month := now.Format("01")
	day := now.Format("02")

	path := fmt.Sprintf("logs/%s/%s/%s/", year, month, day)

	// 20220701.log
	
	fileName := fmt.Sprintf("%s%s%s.log", year, month, day)

	fmt.Println(path)
	fmt.Println(fileName)

	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	fileLog := fmt.Sprintf("%s/%s", path, fileName)
	file, errF := os.Create(fileLog) 

	if errF != nil {
		log.Fatal(errF)
	}
	defer file.Close()
	

}
