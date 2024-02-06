package main

import (
    "log"
    "os"
)

func main() {
    // create info log
    fileInfo, err := openLogFile("./myinfo.log")
    if err != nil {
        log.Fatal(err)
    }
    infoLog := log.New(fileInfo, "[info]", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)
    infoLog.Println("this is info")

    // create error log
    fileError, err := openLogFile("./myerror.log")
    if err != nil {
        log.Fatal(err)
    }
    errorLog := log.New(fileError, "[error]", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)
    errorLog.Println("this is error")
}

func openLogFile(path string) (*os.File, error) {
    logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
    if err != nil {
        return nil, err
    }
    return logFile, nil
}