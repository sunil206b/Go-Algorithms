package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("my_file.txt", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buffWriter := bufio.NewWriter(file)
	bs := []byte{97, 98, 99}
	bytsWritten, err := buffWriter.Write(bs)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes written to buffer (not file) %d\n", bytsWritten)

	bytesAvailable := buffWriter.Available()
	log.Printf("Bytes available in buffer: %d\n", bytesAvailable)

	bytsWritten, err = buffWriter.WriteString("\nJust a random string")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes written to buffer (not file) %d\n", bytsWritten)
	unflushedBufferSize := buffWriter.Buffered()
	log.Printf("Bytes buffered: %d\n", unflushedBufferSize)
	buffWriter.Flush()

	bytsWritten, err = buffWriter.WriteString("\nAnother random string")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes written to buffer (not file) %d\n", bytsWritten)
	buffWriter.Reset(buffWriter)
}