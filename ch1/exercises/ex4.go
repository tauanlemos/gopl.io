// Modify dup2 to print the names of all files in which each duplicated line occurs.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, os.Args[0])
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, arg)
			f.Close()
		}
	}
	for line, n := range counts {
		fmt.Printf("%s\t%s\n", line, n)
	}
}

func countLines(f *os.File, counts map[string]string, s string) {
	// fmt.Println(s)
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()] = s
	}
	// NOTE: ignoring potential errors from input.Err()
}

