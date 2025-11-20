package internal

import (
	"testing"

	"github.com/jrecuero/DandD/internal/character"
)

func TestGetAttributeName(t *testing.T) {
	tests := []struct {
		attr     character.Attribute
		expected string
	}{
		{character.Str, "strength"},
		{character.Dex, "dexterity"},
		{character.Con, "constitution"},
		{character.Int, "intelligence"},
		{character.Wis, "wisdom"},
		{character.Cha, "charisma"},
	}

	for _, tt := range tests {
		if got := character.GetAttributeName(tt.attr); got != tt.expected {
			t.Errorf("GetAttributeName(%v) = %q; want %q", tt.attr, got, tt.expected)
		}
	}
}

func TestGetAttributeShortName(t *testing.T) {
	tests := []struct {
		attr     character.Attribute
		expected string
	}{
		{character.Str, "STR"},
		{character.Dex, "DEX"},
		{character.Con, "CON"},
		{character.Int, "INT"},
		{character.Wis, "WIS"},
		{character.Cha, "CHA"},
	}

	for _, tt := range tests {
		if got := character.GetAttributeShortName(tt.attr); got != tt.expected {
			t.Errorf("GetAttributeShortName(%v) = %q; want %q", tt.attr, got, tt.expected)
		}
	}
}

func TestAbilityModifier(t *testing.T) {
	tests := []struct {
		score    int
		expected int
	}{
		{10, 0},
		{9, 0},
		{8, -1},
		{15, 2},
		{18, 4},
		{3, -3},
		{12, 1},
		{7, -1},
		{20, 5},
	}

	for _, tt := range tests {
		if got := character.AbilityModifier(tt.score); got != tt.expected {
			t.Errorf("AbilityModifier(%d) = %d; want %d", tt.score, got, tt.expected)
		}
	}
}

func TestGetAttributeFromName(t *testing.T) {
	tests := []struct {
		name     string
		expected character.Attribute
		found    bool
	}{
		{"strength", character.Str, true},
		{"dexterity", character.Dex, true},
		{"constitution", character.Con, true},
		{"intelligence", character.Int, true},
		{"wisdom", character.Wis, true},
		{"charisma", character.Cha, true},
		{"invalid", 0, false},
	}
	for _, tt := range tests {
		attr, found := character.GetAttributeFromName(tt.name)
		if found != tt.found || attr != tt.expected {
			t.Errorf("GetAttributeFromName(%q) = (%v, %v); want (%v, %v)", tt.name, attr, found, tt.expected, tt.found)
		}
	}
}

func TestGetAttributeFromShortName(t *testing.T) {
	tests := []struct {
		shortName string
		expected  character.Attribute
		found     bool
	}{
		{"STR", character.Str, true},
		{"DEX", character.Dex, true},
		{"CON", character.Con, true},
		{"INT", character.Int, true},
		{"WIS", character.Wis, true},
		{"CHA", character.Cha, true},
		{"BAD", 0, false},
	}
	for _, tt := range tests {
		attr, found := character.GetAttributeFromShortName(tt.shortName)
		if found != tt.found || attr != tt.expected {
			t.Errorf("GetAttributeFromShortName(%q) = (%v, %v); want (%v, %v)", tt.shortName, attr, found, tt.expected, tt.found)
		}
	}
}

func TestGetAttributeFromName_CaseInsensitive(t *testing.T) {
	cases := []struct {
		input    string
		expected character.Attribute
		found    bool
	}{
		{"Strength", character.Str, true},
		{"STRENGTH", character.Str, true},
		{"sTrEnGtH", character.Str, true},
		{"dexTERity", character.Dex, true},
		{"CONSTITUTION", character.Con, true},
		{"intelligence", character.Int, true},
		{"WISDOM", character.Wis, true},
		{"cHaRiSmA", character.Cha, true},
	}
	for _, c := range cases {
		attr, found := character.GetAttributeFromName(c.input)
		if !found || attr != c.expected {
			t.Errorf("GetAttributeFromName(%q) = (%v, %v); want (%v, true)", c.input, attr, found, c.expected)
		}
	}
}

func TestGetAttributeFromShortName_CaseInsensitive(t *testing.T) {
	cases := []struct {
		input    string
		expected character.Attribute
		found    bool
	}{
		{"str", character.Str, true},
		{"STR", character.Str, true},
		{"sTr", character.Str, true},
		{"dex", character.Dex, true},
		{"CON", character.Con, true},
		{"int", character.Int, true},
		{"wis", character.Wis, true},
		{"cHa", character.Cha, true},
	}
	for _, c := range cases {
		attr, found := character.GetAttributeFromShortName(c.input)
		if !found || attr != c.expected {
			t.Errorf("GetAttributeFromShortName(%q) = (%v, %v); want (%v, true)", c.input, attr, found, c.expected)
		}
	}
}
