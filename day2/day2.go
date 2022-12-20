package main

import (
	"advent-of-code-2022/file-reader"
	"fmt"
	"strings"
)

const DRAW = 3
const WIN = 6
const LOSS = 0

const OPPONENT_ROCK = "A"
const YOU_ROCK = "X"

const OPPONENT_PAPER = "B"
const YOU_PAPER = "Y"

const OPPONENT_SCISSORS = "C"
const YOU_SCISSORS = "Z"

const OUTCOME_WIN = "Z"
const OUTCOME_LOSE = "X"
const OUTCOME_DRAW = "Y"

func main() {
	var content, _ = file_reader.ReadFile("day2.txt")

	var totalScore int = 0

	for _, row := range content {
		var moves = strings.Split(row, " ")
		var opponent = moves[0]
		var you = moves[1]

		var calculatedScore = calculateScore(opponent, you)
		totalScore += calculatedScore
	}

	var scoreBasedOnOutcome = star2()

	fmt.Println(totalScore)
	fmt.Println(scoreBasedOnOutcome)
}

func star2() int {
	var content, _ = file_reader.ReadFile("day2.txt")

	var totalScore int = 0

	for _, row := range content {
		var moves = strings.Split(row, " ")
		var opponent = moves[0]
		var outcome = moves[1]
		var you = moveOnOutcome(outcome, opponent)

		var calculatedScore = calculateScore(opponent, you)
		totalScore += calculatedScore
	}

	return totalScore
}

func getMoveScore(move string) int {
	switch move {
	case YOU_ROCK:
		return 1
	case YOU_PAPER:
		return 2
	case YOU_SCISSORS:
		return 3
	default:
		panic("Cannot handle this case")

		return 0
	}
}

func moveOnOutcome(outcome string, opponent string) string {
	switch outcome {
	case OUTCOME_WIN:
		if opponent == OPPONENT_SCISSORS {
			return YOU_ROCK
		}

		if opponent == OPPONENT_PAPER {
			return YOU_SCISSORS
		}

		return YOU_PAPER
	case OUTCOME_DRAW:
		if opponent == OPPONENT_SCISSORS {
			return YOU_SCISSORS
		}

		if opponent == OPPONENT_PAPER {
			return YOU_PAPER
		}

		return YOU_ROCK
	case OUTCOME_LOSE:
		if opponent == OPPONENT_SCISSORS {
			return YOU_PAPER
		}

		if opponent == OPPONENT_PAPER {
			return YOU_ROCK
		}

		return YOU_SCISSORS
	default:
		panic("Unhandled case")
	}
}

func calculateScore(opponent string, you string) int {
	switch opponent {
	case OPPONENT_ROCK:
		if you == YOU_PAPER {
			return WIN + getMoveScore(you)
		}

		if you == YOU_ROCK {
			return DRAW + getMoveScore(you)
		}

		return LOSS + getMoveScore(you)
	case OPPONENT_PAPER:
		if you == YOU_PAPER {
			return DRAW + getMoveScore(you)
		}

		if you == YOU_SCISSORS {
			return WIN + getMoveScore(you)
		}

		return LOSS + getMoveScore(you)
	case OPPONENT_SCISSORS:
		if you == YOU_ROCK {
			return WIN + getMoveScore(you)
		}

		if you == YOU_SCISSORS {
			return DRAW + getMoveScore(you)
		}

		return LOSS + getMoveScore(you)
	default:
		panic("Cannot handle this case")
	}
}
