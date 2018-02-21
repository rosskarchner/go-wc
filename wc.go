package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func count_file(file io.Reader) (lines, words, bytes int) {
	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	for scanner.Scan() {
		line := scanner.Text()
		words += len(strings.Fields(line))
		// scanner strips out newlines, but wc counts them
		bytes += len(line + "\n")
		lines += 1

	}
	return

}
func main() {

	// TODO: how do I support expressing the flags together?
	// -wlc vs -w -l -c
	// TODO: BSD wc differentiates between bytes and charachters (-c vs -m)
	do_wordsPtr := flag.Bool("w", false, "count words")
	do_linesPtr := flag.Bool("l", false, "count lines")
	do_bytesPtr := flag.Bool("c", false, "count bytes")

	flag.Parse()

	// if no arguments are set, then do words, lines, and bytes
	if !(*do_wordsPtr || *do_linesPtr || *do_bytesPtr) {
		*do_wordsPtr = true
		*do_linesPtr = true
		*do_bytesPtr = true
	}

	format_counts := func(lines, words, bytes int, label string) {

		if *do_linesPtr {
			fmt.Printf(" %d", lines)
		}
		if *do_wordsPtr {
			fmt.Printf(" %d", words)
		}
		if *do_bytesPtr {
			fmt.Printf(" %d", bytes)
		}

		fmt.Printf(" %s\n", label)
	}
	// set the total counters
	total_word_count := 0
	total_line_count := 0
	total_byte_count := 0

	// if there are positional arguments
	if flag.NArg() != 0 {
		for _, filename := range flag.Args() {
			file, err := os.Open(filename)
			check(err)
			lines, words, bytes := count_file(file)
			format_counts(lines, words, bytes, filename)
			total_line_count += lines
			total_word_count += words
			total_byte_count += bytes

		}

		if flag.NArg() > 1 {
			format_counts(total_line_count, total_word_count, total_byte_count, "total")
		}

	} else {
		// no positional arguments, use stdin
		words, lines, bytes := count_file(os.Stdin)
		format_counts(lines, words, bytes, "")
	}
}
