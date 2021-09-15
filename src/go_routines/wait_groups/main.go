package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"sync"
)


var (
	matches []string
	wg = sync.WaitGroup{}
	lock = sync.Mutex{}
)

func fileSearch(root string, filename string) {
	fmt.Println("Searching in", root)
	files, _ := ioutil.ReadDir(root)
	for _, file := range files {
		if strings.Contains(file.Name(), filename) {
			lock.Lock()
			matches = append(matches, filepath.Join(root, file.Name()))
			lock.Unlock()
		}
		if file.IsDir() {
			wg.Add(1)
			fileSearch(filepath.Join(root, file.Name()), filename)
		}
	}
	wg.Done()
}

func main() {
	wg.Add(1)
	go fileSearch("C:/React_Workspace", "README.md")
	wg.Wait()
	for _, file := range matches {
		fmt.Println("Matched: ", file)
	}
}