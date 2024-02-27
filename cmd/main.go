package main

import (
	"github.com/alexunjm/money-game/core/game/gameState"
)

func main() {
	state := gameState.NewState()
	for {

		state.ShowMenu()
		ns := state.ExecSelectedOption()
		if ns.GameOver {
			state.Finish()
			return
		}
	}
}
