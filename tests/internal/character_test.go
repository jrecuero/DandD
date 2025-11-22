package internal

import (
	"testing"

	"github.com/jrecuero/DandD/internal/character"
)

func TestNewCharacter(t *testing.T) {
	attrs := character.NewAttributesMap()
	attrs.Set(character.Str, 10)
	attrs.Set(character.Dex, 12)
	attrs.Set(character.Con, 14)
	attrs.Set(character.Int, 8)
	attrs.Set(character.Wis, 9)
	attrs.Set(character.Cha, 11)

	char := character.NewCharacter("Aragorn", "Ranger", attrs)

	if char.Name != "Aragorn" {
		t.Errorf("expected name 'Aragorn', got %s", char.Name)
	}
	if char.Job != "Ranger" {
		t.Errorf("expected job 'Ranger', got %s", char.Job)
	}
	if char.Attributes.Get(character.Str) != 10 {
		t.Errorf("expected STR=10, got %d", char.Attributes.Get(character.Str))
	}
}

func TestCharacter_String(t *testing.T) {
	attrs := character.NewAttributesMap()
	attrs.Set(character.Str, 15)
	attrs.Set(character.Dex, 12)
	attrs.Set(character.Con, 10)
	attrs.Set(character.Int, 8)
	attrs.Set(character.Wis, 9)
	attrs.Set(character.Cha, 13)

	char := character.NewCharacter("Gandalf", "Wizard", attrs)
	expected := "Gandalf the Wizard [STR: 15, DEX: 12, CON: 10, INT: 8, WIS: 9, CHA: 13]"

	if got := char.String(); got != expected {
		t.Errorf("Character.String() = %q; want %q", got, expected)
	}
}

func TestCharacter_AttributesModification(t *testing.T) {
	attrs := character.NewAttributesMap()
	char := character.NewCharacter("TestChar", "Fighter", attrs)

	// Modify attributes after creation
	char.Attributes.Set(character.Str, 16)
	char.Attributes.Set(character.Con, 14)

	if char.Attributes.Get(character.Str) != 16 {
		t.Errorf("expected STR=16, got %d", char.Attributes.Get(character.Str))
	}
	if char.Attributes.Get(character.Con) != 14 {
		t.Errorf("expected CON=14, got %d", char.Attributes.Get(character.Con))
	}
}
