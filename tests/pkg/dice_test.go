package pkg

import (
	"testing"

	"github.com/jrecuero/DandD/pkg/dice"
)

func TestRollDie(t *testing.T) {
	tests := []struct {
		name  string
		sides int
	}{
		{"d6", 6},
		{"d20", 20},
		{"d4", 4},
		{"d8", 8},
		{"d10", 10},
		{"d12", 12},
		{"d100", 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Roll the die multiple times to verify the range
			for i := 0; i < 100; i++ {
				result := dice.RollDie(tt.sides)
				if result < 1 || result > tt.sides {
					t.Errorf("RollDie(%d) = %d; want value between 1 and %d", tt.sides, result, tt.sides)
				}
			}
		})
	}
}

func TestRollDice(t *testing.T) {
	tests := []struct {
		name    string
		numDice int
		sides   int
	}{
		{"2d6", 2, 6},
		{"3d6", 3, 6},
		{"1d20", 1, 20},
		{"4d6", 4, 6},
		{"5d8", 5, 8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rolls := dice.RollDice(tt.numDice, tt.sides)

			// Check correct number of dice rolled
			if len(rolls) != tt.numDice {
				t.Errorf("RollDice(%d, %d) returned %d rolls; want %d",
					tt.numDice, tt.sides, len(rolls), tt.numDice)
			}

			// Check each roll is within valid range
			for i, roll := range rolls {
				if roll < 1 || roll > tt.sides {
					t.Errorf("RollDice(%d, %d)[%d] = %d; want value between 1 and %d",
						tt.numDice, tt.sides, i, roll, tt.sides)
				}
			}
		})
	}
}

func TestSumRolls(t *testing.T) {
	tests := []struct {
		name     string
		rolls    []int
		expected int
	}{
		{"empty", []int{}, 0},
		{"single roll", []int{5}, 5},
		{"multiple rolls", []int{1, 2, 3, 4, 5}, 15},
		{"all ones", []int{1, 1, 1, 1}, 4},
		{"all sixes", []int{6, 6, 6, 6}, 24},
		{"mixed values", []int{3, 5, 2, 6, 1}, 17},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := dice.SumRolls(tt.rolls)
			if result != tt.expected {
				t.Errorf("SumRolls(%v) = %d; want %d", tt.rolls, result, tt.expected)
			}
		})
	}
}

func TestRoll(t *testing.T) {
	tests := []struct {
		name    string
		numDice int
		sides   int
		minSum  int
		maxSum  int
	}{
		{"1d6", 1, 6, 1, 6},
		{"2d6", 2, 6, 2, 12},
		{"3d6", 3, 6, 3, 18},
		{"1d20", 1, 20, 1, 20},
		{"4d6", 4, 6, 4, 24},
		{"2d10", 2, 10, 2, 20},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Roll multiple times to verify the range
			for i := 0; i < 50; i++ {
				result := dice.Roll(tt.numDice, tt.sides)
				if result < tt.minSum || result > tt.maxSum {
					t.Errorf("Roll(%d, %d) = %d; want value between %d and %d",
						tt.numDice, tt.sides, result, tt.minSum, tt.maxSum)
				}
			}
		})
	}
}

func BenchmarkRollDie(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dice.RollDie(6)
	}
}

func BenchmarkRollDice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dice.RollDice(4, 6)
	}
}

func BenchmarkRoll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dice.Roll(3, 6)
	}
}
