package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("Enter Text: ")
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	text := sc.Text()
	bytes := sc.Bytes()

	fmt.Println("Input Text: ", text)
	fmt.Println("Input Bytes: ", bytes)

	for sc.Scan() {
		text := sc.Text()
		fmt.Println("You Entered:", text)
		if text == "exit" {
			fmt.Println("Exiting the scanning ...")
			break
		}
	}
	fmt.Println("Just a message after the for loop.")
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
}