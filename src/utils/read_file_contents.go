package utils

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func RF(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading the file!")
	}

	fmt.Println(string(data))
}

func RFWithScanner(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func RFWithManualBuffer(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 4 bytes to simulate low memory in buffer
	buf := make([]byte, 4)
	for {
		// n here is the number of bytes read!
		n, err := file.Read(buf)
		// this is how u check if we're at the end of a file!
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		// now print what we have
		fmt.Println(string(buf[:n]))
	}
}

func Run() {
	RFWithManualBuffer("../notex.txt")
}
