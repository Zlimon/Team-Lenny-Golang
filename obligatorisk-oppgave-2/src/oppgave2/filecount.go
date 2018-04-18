package main

import (
	"fmt"
	"bufio"
	"os"
	"sort"
	"unicode"
)

func main() {
	filename := os.Args[1]
	fileCount(filename)
}

func fileCount(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	//	Teller antall bokstaver.
	mLetters := make(map[rune]int64)

	lineCounter := 0
	for scanner.Scan() {
		for _, r := range scanner.Text() {
			if unicode.IsLetter(r) {
				mLetters[unicode.ToLower(r)]++
			}
		}
		lineCounter++ // Teller antall linjer.
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	// Sorterer bokstavene i minkende grad.
	type sLetter struct {
		letter rune
		count  int64
	}
	sLetters := make([]sLetter, 0, len(mLetters))
	for letter, count := range mLetters {
		sLetters = append(sLetters, sLetter{letter: letter, count: count})
	}
	sort.Slice(
		sLetters,
		func(i, j int) bool { return sLetters[i].count >= sLetters[j].count },
	)

	// Printer ut all samlet informasjon.
	fmt.Println("Information about file", filename,":\n")
	fmt.Println("Number of lines in file:", lineCounter,"\n")
	n := len(sLetters)
	if len(sLetters) > 5 {
		n = 5
	}
	runeCounter := 1
	for _, letter := range sLetters[:n] {
		fmt.Println(runeCounter, "Rune:", string(letter.letter), ", Counts:", letter.count)
		runeCounter++
	}
}