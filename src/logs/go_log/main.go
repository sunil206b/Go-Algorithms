package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	setupLogger()
	// simpleHttpGet("www.google.com")
	simpleHttpGet("http://www.google.com")
}

func setupLogger() {
	path, _ := os.Getwd()
	fmt.Println(path)
	logFile, _ := os.OpenFile(path+"/log/test.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
	defer logFile.Close()
	log.SetOutput(logFile)
	fmt.Println("Log setup done")
}

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching url %s : %s", url, err.Error())
	} else {
		log.Printf("Status Code for %s : %s", url, resp.Status)
	}
	defer resp.Body.Close()
}