package WordChecker

func IsWordCorrect(word string, allWords []string, wordCounter int, wordsGuessed [6]string) int {

	var value int

	for i := 0; i < wordCounter; i++ {
		if word == wordsGuessed[i] {
			return 2
		}
	}

	wordsForSearch := allWords
	c := make(chan int)
	d := make(chan int)
	e := make(chan int)
	go searchWordsList(word, wordsForSearch[:5000], c)
	go searchWordsList(word, wordsForSearch[5000:10000], d)
	go searchWordsList(word, wordsForSearch[10000:], e)

	value1 := <-c
	value2 := <-d
	value3 := <-e
	value = value1 + value2 + value3
	return value
}

func searchWordsList(word string, wordsForSearch []string, c chan int) {
	wordChecked := 0
	for i := 0; i < len(wordsForSearch); i++ {
		if word == wordsForSearch[i] {
			c <- 1
		}
		wordChecked++
		if wordChecked == len(wordsForSearch) {
			c <- 0
		}

	}
}
