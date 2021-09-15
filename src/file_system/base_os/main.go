package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var newFile *os.File
	fmt.Printf("%T\n", newFile)

	newFile, err := os.Create("a.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()
	n, err := newFile.WriteString("something")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)

	err = os.Truncate("a.txt", 0)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile("a.txt", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println("File Name:", file.Name())
	
	fileInfo, err := os.Stat("a.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File Name:", fileInfo.Name())
	fmt.Println("File Size:", fileInfo.Size())
	fmt.Println("Last Modified:", fileInfo.ModTime())
	fmt.Println("is Directory?", fileInfo.IsDir())
	fmt.Println("Permissions", fileInfo.Mode())

	fileInf, err := os.Stat("b.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("File does not exsit")
		}
	}
	fmt.Println(fileInf)

	oldPath := "a.txt"
	newPath := "aaa.txt"
	err = os.Rename(oldPath, newPath)
	if err != nil {
		log.Fatal(err)
	}
}