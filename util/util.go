package util

import (
	"bufio"
	"io"
	"log"
	"os"
)

func InputFromFile(filename string) io.Reader {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	// defer file.Close()

	return file
}

func ProcessByLine(input io.Reader, process func(string, int)) {
	scanner := bufio.NewScanner(input)
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
