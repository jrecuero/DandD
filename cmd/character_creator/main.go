package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/jrecuero/DandD/internal/character"
	"github.com/jrecuero/DandD/pkg/dice"
)

type rollData struct {
	attribute   character.Attribute
	attrValue   int
	dc          int
	increases   map[character.Attribute]int
	failEffects map[character.Attribute]int
}

// Constants for file paths
// and character creation JSON file.
const (
	character_creation_file = "character_creation.json"
	data_assets_path        = "./assets/data/"
)

// loadCharacterData loads the character data from the JSON file.
// It panics if there is an error.
// Returns the loaded CharacterData.
func loadCharacterData() *character.CharacterCreationData {
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
func loadAttributes(characterData *character.CharacterCreationData) character.AttributesMap {
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
// It returns a rollData struct with relevant information.
func displayAnswer(anwser character.Answer, attributes character.AttributesMap) *rollData {
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
	return &rollData{
		attribute:   attr,
		attrValue:   attrValue,
		dc:          anwser.DC,
		increases:   increases,
		failEffects: failEffects,
	}
}

// displayDots displays dots in the console for the specified duration.
// It is used to simulate waiting time.
// It takes the duration as a parameter.
func displayDots(duration time.Duration) {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()
	done := time.After(duration)
	for {
		select {
		case <-ticker.C:
			fmt.Print(".")
		case <-done:
			fmt.Println()
			return
		}
	}
}

// rollDice performs the attribute test roll.
// It takes the rollData and current AttributesMap as parameters.
// It rolls a d20, adds the attribute value, and compares it to the DC.
// Depending on the result, it applies increases or fail effects to the AttributesMap.
// It prints the results to the console.
func rollDice(rollData *rollData, attributes character.AttributesMap) {
	fmt.Println("Rolling for attribute test...")
	fmt.Printf("Press Enter to roll the dice")

	// Wait for user to press Enter
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	// Start goroutine to display dots while "rolling"
	dotsChan := make(chan bool)
	go func() {
		displayDots(3 * time.Second)
		dotsChan <- true
	}()

	// Wait for dots to finish
	<-dotsChan

	rollDice := dice.Roll(1, 20)
	mod := character.AbilityModifier(rollData.attrValue)
	total := rollDice + mod
	fmt.Printf("You rolled a d20 + modifier (%d): %d\n", mod, rollDice-mod)
	fmt.Printf("Rolled: %d + %d = %d vs DC %d\n", rollDice, mod, total, rollData.dc)
	if total >= rollData.dc {
		fmt.Println("Test passed! Applying increases.")
		for attr, inc := range rollData.increases {
			attributes.Increase(attr, inc)
			fmt.Printf("- Increased %s by %d\n", character.GetAttributeName(attr), inc)
		}
	} else {
		fmt.Println("Test failed! Applying fail effects.")
		for attr, dec := range rollData.failEffects {
			attributes.Decrease(attr, dec)
			fmt.Printf("- Decreased %s by %d\n", character.GetAttributeName(attr), dec)
		}
	}
	fmt.Println("Updated attributes:", attributes.ColorString())
	fmt.Println()
}

func main() {
	var character_name string
	var character_job string
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter character name: ")
	character_name, _ = reader.ReadString('\n')
	character_name = strings.TrimSpace(character_name)

	fmt.Printf("Enter character job: ")
	character_job, _ = reader.ReadString('\n')
	character_job = strings.TrimSpace(character_job)
	character := character.NewCharacter(character_name, character_job, character.AttributesMap{})

	characterData := loadCharacterData()
	attributes := loadAttributes(characterData)

	fmt.Println("Initial attributes:", attributes)

	for _, question := range characterData.Questions {
		fmt.Printf("Year %d: %s\n", question.Year, question.Question)
		anwser := chooseAnswer(question.Answers)
		rollData := displayAnswer(anwser, attributes)
		rollDice(rollData, attributes)
	}
	character.Attributes = attributes
	fmt.Println("Final character:")
	fmt.Printf("Name: %s\n", character.Name)
	fmt.Printf("Job: %s\n", character.Job)
	fmt.Printf("Attributes: %s\n", character.Attributes.ColorString())
}
