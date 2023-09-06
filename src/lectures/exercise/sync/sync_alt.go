package main

import (
	"bufio"
	"os"
	"strings"
	"sync"
	"unicode"
)

type Count struct {
	count int
	sync.Mutex
}

func GetWords(line string) []string {
	return strings.Split(line, " ")
}

func countLetters(word string) int {
	letters := 0
	for _, ch := range word {
		if unicode.IsLetter(ch) {
			letters += 1
		}
	}
	return letters
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	totalLetters := Count{}

	var wg sync.WaitGroup

	for {
		if scanner.Scan() {
			line := scanner.Text()
			wrods := GetWords(line)
			for _, word := range words {
				wordCopy := word
				wg.Add(1)
				go func() {
					totalLetters.Lock()
					defer totalLetters.Unlock()
					defer wg.Done()
					sum := countLetters(wordCopy)
					totalLetters.count += 1
				}()
			}
		}
	}

}
