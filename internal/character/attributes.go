package character

import "strings"

// Attribute represents a character attribute type.
type Attribute int

// Enumeration of character attributes.
// The attributes include Strength, Dexterity, Constitution, Intelligence, Wisdom, and Charisma.
const (
	Str Attribute = iota
	Dex
	Con
	Int
	Wis
	Cha
)

// attributeNames maps each Attribute to its full name.
var attributeNames = map[Attribute]string{
	Str: "strength",
	Dex: "dexterity",
	Con: "constitution",
	Int: "intelligence",
	Wis: "wisdom",
	Cha: "charisma",
}

// attributeShortNames maps each Attribute to its short name.
var attributeShortNames = map[Attribute]string{
	Str: "STR",
	Dex: "DEX",
	Con: "CON",
	Int: "INT",
	Wis: "WIS",
	Cha: "CHA",
}

// GetAttributeShortName returns the short name of the given attribute.
func GetAttributeShortName(attr Attribute) string {
	return attributeShortNames[attr]
}

// GetAttributeName returns the full name of the given attribute.
func GetAttributeName(attr Attribute) string {
	return attributeNames[attr]
}

// GetAttributeFromName returns the Attribute corresponding to the given name.
// It returns the Attribute and a boolean indicating whether the name was found.
// If the name does not correspond to any Attribute, the boolean will be false.
func GetAttributeFromName(name string) (Attribute, bool) {
	name = strings.ToLower(name)
	for attr, attrName := range attributeNames {
		if attrName == name {
			return attr, true
		}
	}
	return 0, false
}

// GetAttributeFromShortName returns the Attribute corresponding to the given short name.
// It returns the Attribute and a boolean indicating whether the short name was found.
// If the short name does not correspond to any Attribute, the boolean will be false.
func GetAttributeFromShortName(shortName string) (Attribute, bool) {
	shortName = strings.ToUpper(shortName)
	for attr, attrShortName := range attributeShortNames {
		if attrShortName == shortName {
			return attr, true
		}
	}
	return 0, false
}

// AbilityModifier calculates the ability modifier for a given ability score.
// The formula used is (score - 10) / 2, rounded down.
// For example, a score of 15 yields a modifier of +2, while a score of 8 yields -1.
// This function is commonly used in role-playing games to determine bonuses or penalties
// associated with character attributes.
// It takes an integer score as input and returns the corresponding modifier as an integer.
func AbilityModifier(score int) int {
	return (score - 10) / 2
}
