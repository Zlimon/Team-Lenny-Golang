package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"unicode"
)

func main() {
	// The Complete Works of William Shakespeare
	// http://www.gutenberg.org/files/100/100-0.txt
	f, err := os.Open(`C:/Users/magnu/GolandProjects/TeamLenny/obligatorisk-oppgave-2/src/oppgave2/text.txt`)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer f.Close()

	// Count letters
	mLetters := make(map[rune]int64)
	s := bufio.NewScanner(f)
	for s.Scan() {
		for _, r := range s.Text() {
			if unicode.IsLetter(r) {
				mLetters[unicode.ToLower(r)]++
			}
		}
	}
	if err := s.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	// Sort letters by descending count
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

	// Print highest 5 letter counts
	n := len(sLetters)
	if len(sLetters) > 5 {
		n = 5
	}
	for _, letter := range sLetters[:n] {
		fmt.Println(string(letter.letter), letter.count)
	}
}
