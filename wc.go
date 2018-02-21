package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if err := scanner.Err(); err != nil {
		log.Println(err)
		os.Exit(1)
	}


	// these will become command line arguments
	do_words := true
	do_lines := true
	do_bytes := true

	// set the counters
	word_count := 0
	line_count := 0
	byte_count := 0

	for scanner.Scan() {
		line := scanner.Text()
		word_count += len(strings.Fields(line))
		byte_count += len(line)
		line_count += 1

	}

	if do_lines{
		fmt.Printf("      %v", line_count)
	}
	if do_words{
		fmt.Printf("      %v",word_count)
	}
	if do_bytes{
		fmt.Printf("      %v",byte_count)
	}
	fmt.Printf("\n")
}
