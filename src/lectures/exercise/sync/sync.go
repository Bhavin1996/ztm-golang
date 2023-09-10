//--Summary:
//  Create a program that can read text from standard input and count the
//  number of letters present in the input.
//
//--Requirements:
//* Count the total number of letters in any chosen input
//* The input must be supplied from standard input
//* Input analysis must occur per-word, and each word must be analyzed
//  within a goroutine
//* When the program finishes, display the total number of letters counted
//
//--Notes:
//* Use CTRL+D (Mac/Linux) or CTRL+Z (Windows) to signal EOF, if manually
//  entering data
//* Use `cat FILE | go run ./exercise/sync` to analyze a file
//* Use any synchronization techniques to implement the program:
//  - Channels / mutexes / wait groups

package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

//var wg sync.WaitGroup

func CountLetters(rd bufio.Reader, result chan int) {
	count := 0
	for {
		r, _, err := rd.ReadRune()

		if err != nil {
			break
		}
		if unicode.IsLetter(r) {
			count += 1
		}
	}
	result <- count
}

func main() {
	filepath := "D:/ztm-golang/src/lectures/exercise/sync/test.txt"
	//scanner := bufio.NewScanner(os.Stdin)
	result := make(chan int)
	totalLetters := 0

	//below code creates new file if you need it
	/*file, err := os.Create(filepath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Enter content in the file (type 'exit' to end):")

	for {

		scanner.Scan()
		data := scanner.Text()

		if data == "exit" {
			break // Terminate the loop when the user enters "exit".
		}

		_, err := file.WriteString(data + "\n")
		if err != nil {
			fmt.Println("Error:", err)
			break
		}
	}
	defer file.Close()*/
	//new file creattion done

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error", err)
	}
	defer file.Close()

	rd := bufio.NewReader(file)
	go CountLetters(*rd, result)
	totalLetters = <-result
	fmt.Println("Total letters in a file are ", totalLetters)
}
