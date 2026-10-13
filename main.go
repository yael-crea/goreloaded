package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const maxWordSize = 1024 * 1024

func countWords(reader io.Reader) (int, error) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	scanner.Buffer(make([]byte, 64*1024), maxWordSize)

	count := 0
	for scanner.Scan() {
		count++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return count, nil
}

func run(args []string, stdout, stderr io.Writer) int {
	if len(args) != 1 {
		fmt.Fprintln(stderr, "usage: go run . <fichier>")
		return 1
	}

	file, err := os.Open(args[0])
	if err != nil {
		fmt.Fprintf(stderr, "impossible d'ouvrir %q: %v\n", args[0], err)
		return 1
	}
	defer file.Close()

	count, err := countWords(file)
	if err != nil {
		fmt.Fprintf(stderr, "impossible de lire %q: %v\n", args[0], err)
		return 1
	}

	fmt.Fprintln(stdout, count)
	return 0
}

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

