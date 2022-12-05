package util

import (
	"bufio"
	"log"
	"os"
)

func ProcessByLine(filename string, process func(string, int)) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 1
	for scanner.Scan() {
		line := scanner.Text()
		process(line, lineNumber)
		lineNumber++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
