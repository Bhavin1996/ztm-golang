//--Summary:
//  Create a program that can create a report of rune information from
//  lines of text.
//
//--Requirements:
//* Create a single function to iterate over each line of text that is
//  provided in main().
//  - The function must return nothing and must execute a closure
//* Using closures, determine the following information about the text and
//  print a report to the terminal:
//  - Number of letters
//  - Number of digits
//  - Number of spaces
//  - Number of punctuation marks
//
//--Notes:
//* The `unicode` stdlib package provides functionality for rune classification

package main

import (
	"fmt"
	"unicode"
)

type LineIterator func(line string)

func report(ln []string, call LineIterator) {
	for _, line := range ln {
		call(line)
	}
}

func main() {
	lines := []string{
		"There are",
		"68 letters,",
		"five digits,",
		"12 spaces,",
		"and 4 punctuation marks in these lines of text!",
	}
	NumberOfLetters := 0
	NumberOfDigits := 0
	NumberOfSpaces := 0
	NumberOfPunctuation := 0

	res := func(line string) {
		for _, ln := range line {
			if unicode.IsLetter(ln) {
				NumberOfLetters += 1
			} else if unicode.IsDigit(ln) {
				NumberOfDigits += 1
			} else if unicode.IsSpace(ln) {
				NumberOfSpaces += 1
			} else if unicode.IsPunct(ln) {
				NumberOfPunctuation += 1
			}
		}
	}
	report(lines, res)
	fmt.Println("Number of Letters are", NumberOfLetters)
	fmt.Println("Number of Digits are", NumberOfDigits)
	fmt.Println("Number of spaces are", NumberOfSpaces)
	fmt.Println("Number of Punctuations are", NumberOfPunctuation)
}
