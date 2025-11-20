package character

import "fmt"

// AttributesMap defines a mapping from string keys to Attribute values.
type AttributesMap map[Attribute]int

// NewAttributesMap creates and returns a new AttributesMap with all attributes initialized to zero.
// This function is useful for initializing a character's attributes before setting specific values.
// It returns a pointer to the newly created AttributesMap.
func NewAttributesMap() AttributesMap {
	return AttributesMap{
		Str: 0,
		Dex: 0,
		Con: 0,
		Int: 0,
		Wis: 0,
		Cha: 0,
	}
}

// Get retrieves the value of the specified attribute from the AttributesMap.
// It takes an Attribute as input and returns the corresponding integer value.
// If the attribute does not exist in the map, it returns zero.
func (am AttributesMap) Get(attr Attribute) int {
	return am[attr]
}

// Set assigns a new value to the specified attribute in the AttributesMap.
// It takes an Attribute and an integer value as input and updates the map accordingly.
// If the attribute does not exist in the map, it will be added.
func (am AttributesMap) Set(attr Attribute, value int) {
	am[attr] = value
}

// Increase increments the value of the specified attribute by the given amount.
// It takes an Attribute and an integer value as input and updates the map accordingly.
// If the attribute does not exist in the map, it will be added with the incremented value.
func (am AttributesMap) Increase(attr Attribute, value int) {
	am[attr] += value
}

// Decrease decrements the value of the specified attribute by the given amount.
// It takes an Attribute and an integer value as input and updates the map accordingly.
// If the attribute does not exist in the map, it will be added with the decremented value.
func (am AttributesMap) Decrease(attr Attribute, value int) {
	am[attr] -= value
}

// String returns a string representation of the AttributesMap.
// It formats the attributes and their values in a readable manner.
func (am AttributesMap) String() string {
	return fmt.Sprintf("STR: %d, DEX: %d, CON: %d, INT: %d, WIS: %d, CHA: %d",
		am[Str], am[Dex], am[Con], am[Int], am[Wis], am[Cha])
}

// AttributeMapToString converts an AttributesMap to a formatted string.
// It lists each attribute and its corresponding value in a readable format.
func AttributeMapToString(m map[Attribute]int) string {
	result := ""
	for _, attr := range []Attribute{Str, Dex, Con, Int, Wis, Cha} {
		if m[attr] != 0 {
			result += fmt.Sprintf("%s: %d, ", GetAttributeShortName(attr), m[attr])
		}
	}
	if len(result) > 2 {
		result = result[:len(result)-2] // Remove trailing comma and space
	}
	return result
}
