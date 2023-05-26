package InitWords

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

var allWords []string

func SaveAllWords() []string {
	filePath := "/home/n1p3/GolandProjects/WordleProjekt/words"
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		allWords = append(allWords, fileScanner.Text())
	}

	readFile.Close()
	return allWords
}

func GiveRandomWord() string {
	var correctWord string
	correctWord = allWords[rand.Intn(14848)]
	fmt.Println(correctWord)
	return correctWord
}
