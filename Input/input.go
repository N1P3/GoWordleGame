package Input

import rl "github.com/gen2brain/raylib-go/raylib"

func ReadInput() rune {
	var letter rune
	if rl.IsKeyPressed(rl.KeyA) {
		letter = 'A'
	}
	if rl.IsKeyPressed(rl.KeyB) {
		letter = 'B'
	}
	if rl.IsKeyPressed(rl.KeyC) {
		letter = 'C'
	}
	if rl.IsKeyPressed(rl.KeyD) {
		letter = 'D'
	}
	if rl.IsKeyPressed(rl.KeyE) {
		letter = 'E'
	}
	if rl.IsKeyPressed(rl.KeyF) {
		letter = 'F'
	}
	if rl.IsKeyPressed(rl.KeyG) {
		letter = 'G'
	}
	if rl.IsKeyPressed(rl.KeyH) {
		letter = 'H'
	}
	if rl.IsKeyPressed(rl.KeyI) {
		letter = 'I'
	}
	if rl.IsKeyPressed(rl.KeyJ) {
		letter = 'J'
	}
	if rl.IsKeyPressed(rl.KeyK) {
		letter = 'K'
	}
	if rl.IsKeyPressed(rl.KeyL) {
		letter = 'L'
	}
	if rl.IsKeyPressed(rl.KeyM) {
		letter = 'M'
	}
	if rl.IsKeyPressed(rl.KeyN) {
		letter = 'N'
	}
	if rl.IsKeyPressed(rl.KeyO) {
		letter = 'O'
	}
	if rl.IsKeyPressed(rl.KeyP) {
		letter = 'P'
	}
	if rl.IsKeyPressed(rl.KeyQ) {
		letter = 'Q'
	}
	if rl.IsKeyPressed(rl.KeyR) {
		letter = 'R'
	}
	if rl.IsKeyPressed(rl.KeyS) {
		letter = 'S'
	}
	if rl.IsKeyPressed(rl.KeyT) {
		letter = 'T'
	}
	if rl.IsKeyPressed(rl.KeyU) {
		letter = 'U'
	}
	if rl.IsKeyPressed(rl.KeyV) {
		letter = 'V'
	}
	if rl.IsKeyPressed(rl.KeyW) {
		letter = 'W'
	}
	if rl.IsKeyPressed(rl.KeyX) {
		letter = 'X'
	}
	if rl.IsKeyPressed(rl.KeyY) {
		letter = 'Y'
	}
	if rl.IsKeyPressed(rl.KeyZ) {
		letter = 'Z'
	}
	if rl.IsKeyPressed(rl.KeyBackspace) {
		letter = '-'
	}
	if rl.IsKeyDown(rl.KeyLeftControl) && rl.IsKeyPressed(rl.KeyR) {
		letter = '*'
	}
	return letter
}
