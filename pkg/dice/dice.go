package dice

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RollDie simulates rolling a die with the given number of sides.
// It returns a random integer between 1 and the number of sides, inclusive.
// For example, RollDie(6) simulates rolling a standard six-sided die.
// sides must be greater than 0.
func RollDie(sides int) int {
	return rand.Intn(sides) + 1
}

// RollDice simulates rolling a specified number of dice with the given number of sides.
// It returns a slice of integers representing the result of each die roll.
// numDice must be greater than 0.
// sides must be greater than 0.
func RollDice(numDice, sides int) []int {
	rolls := make([]int, numDice)
	for i := 0; i < numDice; i++ {
		rolls[i] = RollDie(sides)
	}
	return rolls
}

// SumRolls returns the sum of a slice of dice rolls.
// It returns the total sum of the rolls.
// rolls is a slice of integers representing individual die rolls.
func SumRolls(rolls []int) int {
	sum := 0
	for _, roll := range rolls {
		sum += roll
	}
	return sum
}

// Roll simulates rolling a specified number of dice with the given number of sides
// and returns the total sum of the rolls.
// numDice must be greater than 0.
// sides must be greater than 0.
func Roll(numDice, sides int) int {
	rolls := RollDice(numDice, sides)
	return SumRolls(rolls)
}
