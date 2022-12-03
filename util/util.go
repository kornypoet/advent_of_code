package util

import (
	"bufio"
	"log"
	"os"
)

func ProcessByLine(filename string, process func(string)) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		process(line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}