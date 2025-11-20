package internal

import (
	"testing"

	"github.com/jrecuero/DandD/internal/character"
)

func TestNewAttributesMap(t *testing.T) {
	am := character.NewAttributesMap()
	for _, attr := range []character.Attribute{
		character.Str, character.Dex, character.Con,
		character.Int, character.Wis, character.Cha,
	} {
		if got := am.Get(attr); got != 0 {
			t.Errorf("NewAttributesMap: expected %v to be 0, got %d", attr, got)
		}
	}
}

func TestAttributesMap_SetAndGet(t *testing.T) {
	am := character.NewAttributesMap()
	am.Set(character.Str, 15)
	am.Set(character.Dex, 12)
	if got := am.Get(character.Str); got != 15 {
		t.Errorf("Set/Get: expected Str=15, got %d", got)
	}
	if got := am.Get(character.Dex); got != 12 {
		t.Errorf("Set/Get: expected Dex=12, got %d", got)
	}
}

func TestAttributesMap_Increase(t *testing.T) {
	am := character.NewAttributesMap()
	am.Set(character.Con, 5)
	am.Increase(character.Con, 3)
	if got := am.Get(character.Con); got != 8 {
		t.Errorf("Increase: expected Con=8, got %d", got)
	}
	am.Increase(character.Wis, 2)
	if got := am.Get(character.Wis); got != 2 {
		t.Errorf("Increase: expected Wis=2, got %d", got)
	}
}

func TestAttributesMap_Decrease(t *testing.T) {
	am := character.NewAttributesMap()
	am.Set(character.Int, 10)
	am.Decrease(character.Int, 4)
	if got := am.Get(character.Int); got != 6 {
		t.Errorf("Decrease: expected Int=6, got %d", got)
	}
	am.Decrease(character.Cha, 3)
	if got := am.Get(character.Cha); got != -3 {
		t.Errorf("Decrease: expected Cha=-3, got %d", got)
	}
}

func TestAttributesMap_String(t *testing.T) {
	am := character.NewAttributesMap()
	am.Set(character.Str, 8)
	am.Set(character.Dex, 14)
	am.Set(character.Con, 12)
	am.Set(character.Int, 10)
	am.Set(character.Wis, 13)
	am.Set(character.Cha, 7)
	expected := "STR: 8, DEX: 14, CON: 12, INT: 10, WIS: 13, CHA: 7"
	if got := am.String(); got != expected {
		t.Errorf("AttributesMap.String() = %q; want %q", got, expected)
	}
}

func TestAttributeMapToString(t *testing.T) {
	m := map[character.Attribute]int{
		character.Str: 5,
		character.Dex: 7,
		character.Con: 0,
		character.Int: 11,
		character.Wis: 13,
		character.Cha: 15,
	}
	expected := "STR: 5, DEX: 7, INT: 11, WIS: 13, CHA: 15"
	if got := character.AttributeMapToString(m); got != expected {
		t.Errorf("AttributeMapToString() = %q; want %q", got, expected)
	}
}
