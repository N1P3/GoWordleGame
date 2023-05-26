package main

import (
	"WordleProjekt/InitWords"
	"WordleProjekt/Input"
	"WordleProjekt/WordChecker"
	rl "github.com/gen2brain/raylib-go/raylib"
	"strings"
)

const screenWidth = 800
const screenHeight = 800
const squareSide = 70

var colors [6][5]rl.Color
var wordCounter = 0
var letterCounter = 0
var words [6][5]string
var allWords []string
var gameState int
var correctWord string
var wordsGuessed [6]string
var word string

func main() {
	allWords = InitWords.SaveAllWords()
	correctWord = InitWords.GiveRandomWord()
	gameState = 0

	for i := 0; i < 6; i++ {
		for j := 0; j < 5; j++ {
			colors[i][j] = rl.White
		}
	}

	rl.InitWindow(screenWidth, screenHeight, "Wordle")
	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		Draw(correctWord, Input.ReadInput())
	}
	defer rl.CloseWindow()
}

func Draw(correctWord string, letter rune) {
	var isCorrect int
	if gameState == 0 && letter != 0 {
		if letterCounter < 5 && wordCounter < 6 {

			if letter != 0 && letter != '-' {
				words[wordCounter][letterCounter] = string(letter)
				letterCounter++
			}
		}
		if letter == '-' && letterCounter != 0 {
			letterCounter--
			words[wordCounter][letterCounter] = " "

		}

		if letterCounter == 5 {
			var wordWhole []string
			for i := 0; i < letterCounter; i++ {
				wordWhole = append(wordWhole, words[wordCounter][i])
			}

			word = strings.Join(wordWhole, "")
			if WordChecker.IsWordCorrect(word, allWords, wordCounter, wordsGuessed) == 1 && wordCounter < 6 {
				wordsGuessed[wordCounter] = word
				for j := 0; j < 5; j++ {
					for k := 0; k < 5; k++ {
						if word[j] == correctWord[k] {
							colors[wordCounter][j] = rl.Yellow
						}
					}
				}
				for i := 0; i < 5; i++ {
					if word[i] == correctWord[i] {
						colors[wordCounter][i] = rl.Green
					}
				}
				gameState = checkIfGameOver(word, correctWord)
				wordCounter++
				letterCounter = 0
				isCorrect = 1
			} else {
				isCorrect = WordChecker.IsWordCorrect(word, allWords, wordCounter, wordsGuessed)
			}
		}
	}

	rl.BeginDrawing()
	rl.ClearBackground(rl.LightGray)
	var y int32 = 150
	var x int32 = 150
	for i := 0; i < 6; i++ {
		for j := 0; j < 5; j++ {
			rl.DrawRectangle(x, y, squareSide, squareSide, colors[i][j])
			x = x + 100
		}
		y += 100
		x = 150
	}
	y = 165
	x = 165
	for i := 0; i < 6; i++ {
		for j := 0; j < 5; j++ {
			rl.DrawText(words[i][j], x, y, 50, rl.Black)
			x = x + 100
		}
		y += 100
		x = 165

	}

	if isCorrect == 0 && letterCounter == 5 {
		rl.DrawText("INCORRECT\n    WORD", 150, 300, 75, rl.Red)
		rl.DrawText("press backspace to delete", 150, 500, 30, rl.Red)
	}
	if isCorrect == 2 && letterCounter == 5 {
		rl.DrawText("WORD\nUSED", 270, 300, 75, rl.Red)
		rl.DrawText("press backspace to delete", 150, 500, 30, rl.Red)
	}

	if gameState == 1 {
		rl.DrawRectangle(250, 150, 300, 500, rl.White)
		rl.DrawText("CONGRATS\n YOU WON", 265, 200, 45, rl.DarkPurple)
		rl.DrawText("Press:\nctrl+R to restart\nESC to quit", 260, 400, 30, rl.Black)
		if Input.ReadInput() == '*' {
			restartGame()
		}
	}
	if gameState == 2 {
		rl.DrawRectangle(250, 150, 300, 500, rl.White)
		rl.DrawText("Game Over", 265, 200, 50, rl.Red)
		rl.DrawText("Correct word:\n     "+correctWord, 265, 255, 34, rl.Red)
		rl.DrawText("Press:\nctrl+R to restart\nESC to quit", 260, 400, 30, rl.Black)
		if Input.ReadInput() == '*' {
			restartGame()
		}
	}

	rl.EndDrawing()
}

func checkIfGameOver(word string, correctWord string) int {
	if word == correctWord {
		return 1
	} else if wordCounter == 5 {
		return 2
	}
	return 0
}

func restartGame() {
	for i := 0; i < 6; i++ {
		for j := 0; j < 5; j++ {
			colors[i][j] = rl.White
		}
	}
	for i := 0; i < 6; i++ {
		for j := 0; j < 5; j++ {
			words[i][j] = ""
		}
	}
	wordCounter = 0
	letterCounter = 0
	gameState = 0
	correctWord = InitWords.GiveRandomWord()
}
