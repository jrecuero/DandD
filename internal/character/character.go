package character

import (
	"encoding/json"
	"fmt"
	"os"
)

// CharacterData represents the structure of the character data JSON file.
// It includes starting attributes and a list of questions.
type CharacterData struct {
	StartingAttributes map[string]int `json:"starting_attributes"`
	Questions          []Question     `json:"questions"`
}

// Question represents a single question in the character creation process.
// It includes the year, the prompt (yest), and a pool of possible answers.
type Question struct {
	Year     int      `json:"year"`
	Question string   `json:"question"`
	Answers  []Answer `json:"answers_pool"`
}

// Answer represents a possible answer to a question.
// It includes the answer ID, description, attribute increases, test details,
// and fail effects.
type Answer struct {
	AnswerID    string         `json:"id"`
	Description string         `json:"description"`
	Increases   map[string]int `json:"attribute_rewards"`
	Test        string         `json:"test_attribute"`
	DC          int            `json:"dc"`
	FailEffect  map[string]int `json:"fail_penalty"`
}

// LoadCharacterData reads the character data from a JSON file and unmarshals
// it into a CharacterData struct.
// It returns the CharacterData and any error encountered during the process.
func LoadCharacterData(filename string) (*CharacterData, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read JSON file: %w", err)
	}
	var data CharacterData
	if err := json.Unmarshal(file, &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %w", err)
	}
	return &data, nil
}

// GetAttributeIncreases converts the attribute increases from an Answer
// into an AttributesMap. It maps attribute short names to their corresponding
// Attribute values and includes only non-zero increases.
func GetAttributeIncreases(answer Answer) map[Attribute]int {
	attributes := map[Attribute]int{}
	for attrName, value := range answer.Increases {
		if attr, ok := GetAttributeFromShortName(attrName); ok {
			if value != 0 {
				attributes[attr] = value
			}
		}
	}
	return attributes
}

// GetAttributeFailEffects converts the fail effects from an Answer
// into an AttributesMap. It maps attribute short names to their corresponding
// Attribute values and includes only non-zero effects.
func GetAttributeFailEffects(answer Answer) map[Attribute]int {
	attributes := map[Attribute]int{}
	for attrName, value := range answer.FailEffect {
		if attr, ok := GetAttributeFromShortName(attrName); ok {
			if value != 0 {
				attributes[attr] = value
			}
		}
	}
	return attributes
}
