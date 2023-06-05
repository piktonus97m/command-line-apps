package main

import (
	"bufio" // Read text
	"fmt"   // Print formatted output
	"io"    //io.reader interface
	"os"    // use os resources
)

func main() {
	// Calling the count function to count the number of words/
	// Received from the standard input and printing it out
	fmt.Println(count(os.Stdin))
}

func count(r io.Reader) int {
	// A scanner is used to read text from a reader (such as files)
	scanner := bufio.NewScanner(r)

	// Define the scanner split type to words (default is split lines)
	scanner.Split(bufio.ScanWords)

	// Defining a counter
	wordcounter := 0

	// For every world scanned, increment the counter
	for scanner.Scan() {
		wordcounter++
	}

	// Return the total
	return wordcounter
}

// In order to make this program works, you need to do the following:
//  echo "my first command line tool with go" | ./01-basic-word-counter
