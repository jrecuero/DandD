package main

import (
	"fmt"
	"math/rand/v2"
	"path/filepath"

	"github.com/jrecuero/DandD/internal/character"
)

// Constants for file paths
// and character creation JSON file.
const (
	character_creation_file = "character_creation.json"
	data_assets_path        = "./assets/data/"
)

// loadCharacterData loads the character data from the JSON file.
// It panics if there is an error.
// Returns the loaded CharacterData.
func loadCharacterData() *character.CharacterData {
	fpath := filepath.Join(data_assets_path, character_creation_file)
	data, err := character.LoadCharacterData(fpath)
	if err != nil {
		panic(err)
	}
	return data
}

// loadAttributes initializes the AttributesMap
// with the starting attributes from CharacterData.
// Returns the initialized AttributesMap.
func loadAttributes(characterData *character.CharacterData) character.AttributesMap {
	attributesMap := character.NewAttributesMap()
	for attrName, value := range characterData.StartingAttributes {
		if attr, ok := character.GetAttributeFromShortName(attrName); ok {
			attributesMap.Set(attr, value)
		}
	}
	return attributesMap
}

// chooseAnswer selects an answer from the provided answer pool.
// It shuffles the answers and picks the first one.
// Returns the selected Answer.
func chooseAnswer(anwser_pool []character.Answer) character.Answer {
	rand.Shuffle(len(anwser_pool), func(i, j int) { anwser_pool[i], anwser_pool[j] = anwser_pool[j], anwser_pool[i] })
	// for i, answer := range anwser_pool[:3] {
	// 	fmt.Printf("\t%d. %s\n", i+1, answer.Description)
	// }
	anwser := anwser_pool[0]
	return anwser
}

// displayAnswer displays the selected answer details,
// including attribute tests, increases, and fail effects.
// It takes the selected Answer and the current AttributesMap as parameters.
// It prints the details to the console.
func displayAnswer(anwser character.Answer, attributes character.AttributesMap) {
	fmt.Printf("Selected answer: %s\n", anwser.Description)
	attrName := anwser.Test
	attr, ok := character.GetAttributeFromShortName(attrName)
	if !ok {
		panic("invalid attribute name in test: " + attrName)
	}
	attrValue := attributes.Get(attr)
	fmt.Printf("- Attribute to test %s[%d] DC: %d\n", attrName, attrValue, anwser.DC)
	increases := character.GetAttributeIncreases(anwser)
	failEffects := character.GetAttributeFailEffects(anwser)
	fmt.Println("- Attribute increases:", character.AttributeMapToString(increases))
	fmt.Println("- Attribute fail effects:", character.AttributeMapToString(failEffects))
	fmt.Println()
}

func main() {
	characterData := loadCharacterData()
	attributes := loadAttributes(characterData)

	fmt.Println("Initial attributes:", attributes)

	for _, question := range characterData.Questions {
		fmt.Printf("Year %d: %s\n", question.Year, question.Question)
		anwser := chooseAnswer(question.Answers)
		displayAnswer(anwser, attributes)
	}
}
