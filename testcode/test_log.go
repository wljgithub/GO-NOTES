package main

import (
	"log"
	"os"
)

var logFile, err = os.Create("log.txt")

var Log = log.New(logFile, "[INFO]", log.Ldate|log.Ltime|log.Lshortfile)

func main() {

		if err != nil {
			log.Fatal("fail to create file")
		}
		Log.Println("test log ")

}
