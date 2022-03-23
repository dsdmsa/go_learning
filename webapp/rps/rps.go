package rps

import (
	"math/rand"
	"time"
)

const (
	R = 0
	P = 1
	S = 2

	P_WINS = 0
	C_WINS = 1
	DROW   = 3
)

type Round struct {
	Winner          int    `json:"winner"`
	ComputerChoiche string `json:"computer_choiche"`
	ReundResult     string `json:"result"`
}

func PlayRound(playerValue int) Round {
	rand.Seed(time.Now().UnixNano())
	computerValue := rand.Intn(3)
	computerChoiche := ""
	roundResult := ""
	winner := -1

	switch computerValue {
	case R:
		computerChoiche = "Rock"
	case P:
		computerChoiche = "Paper"
	case S:
		computerChoiche = "Sciccor"
	default:
	}

	if playerValue == computerValue {
		roundResult = "It's a dorow"
		winner = DROW
	} else if playerValue == (computerValue+1)%3 {
		roundResult = "Player wins!"
		winner = P_WINS
	} else {
		roundResult = "Computer wins!"
		winner = C_WINS
	}

	var result Round
	result.ComputerChoiche = computerChoiche
	result.ReundResult = roundResult
	result.Winner = winner

	return result
}
