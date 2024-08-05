package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// reading standard.txt and convert to array of lines
	readFile, err := os.Open("standard.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	readFile.Close()

	// get argument as a string
	if len(os.Args) < 2 {
		fmt.Println("I need something to work with")
		os.Exit(1)
	}
	if len(os.Args) > 2 {
		fmt.Println("Give me just one string please")
		os.Exit(2)
	}
	text := string(os.Args[1])

	// looking for "\n" and turn it into "n3wL1ne" so string.Split can find it
	preLine := []rune(text)
	for m := 0; m < len(preLine); m++ {
		arrayMiddle := "n3wL!Ne"
		if preLine[m] == 92 && preLine[m+1] == 'n' {
			array1 := preLine[0:m]
			array2 := preLine[m+2:]
			s1 := string([]rune(array1))
			s2 := string([]rune(array2))
			text = s1 + arrayMiddle + s2
			preLine = ([]rune(text))
		}
	}

	// split the text into lines if required
	nextStep := string(preLine)
	line := strings.Split(nextStep, "n3wL!Ne")

	// loop to work on lines
	for j := 0; j < len(line); j++ {

		// to make or not make new lines in situations with no other text
		if len(text) < 1 {
			break
		}
		if len(line[j]) < 1 && j == 0 {
			continue
		}
		if len(line[j]) < 1 {
			fmt.Println()
			continue
		}

		word := []rune(line[j])

		// row by row loop
		for k := 1; k < 9; k++ {

			// character by character loop
			for i := 0; i < len(word); i++ {
				m := rune(k)

				// reduce each character value by 32 in ascii table,
				// multiply by the 9 rows each character uses in standard.txt,
				// add the row number
				asciiFetch := ((word[i] - 32) * 9) + m

				fmt.Printf(fileLines[asciiFetch])
			}
			fmt.Println()
		}
	}
}
