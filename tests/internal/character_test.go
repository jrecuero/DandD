package internal

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/jrecuero/DandD/internal/character"
)

func TestLoadCharacterData_Success(t *testing.T) {
	tmpDir := t.TempDir()
	jsonPath := filepath.Join(tmpDir, "character.json")
	jsonContent := `{
		"starting_attributes": {"strength": 10, "dexterity": 8},
		"questions": [
			{
				"year": 1,
				"question": "What do you do?",
				"answers_pool": [
					{
						"id": "Y1A1",
						"description": "Fight",
						"attribute_rewards": {"strength": 2},
						"test_attribute": "strength",
						"dc": 10,
						"fail_penalty": {"dexterity": -1}
					}
				]
			}
		]
	}`
	if err := os.WriteFile(jsonPath, []byte(jsonContent), 0644); err != nil {
		t.Fatalf("failed to write temp JSON: %v", err)
	}

	data, err := character.LoadCharacterData(jsonPath)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if data.StartingAttributes["strength"] != 10 {
		t.Errorf("expected strength 10, got %d", data.StartingAttributes["strength"])
	}
	if len(data.Questions) != 1 {
		t.Errorf("expected 1 question, got %d", len(data.Questions))
	}
	q := data.Questions[0]
	if q.Year != 1 || q.Question != "What do you do?" {
		t.Errorf("unexpected question: %+v", q)
	}
	if len(q.Answers) != 1 {
		t.Errorf("expected 1 answer, got %d", len(q.Answers))
	}
	ans := q.Answers[0]
	if ans.AnswerID != "Y1A1" || ans.Description != "Fight" {
		t.Errorf("unexpected answer: %+v", ans)
	}
	if ans.Increases["strength"] != 2 {
		t.Errorf("expected strength increase 2, got %d", ans.Increases["strength"])
	}
	if ans.Test != "strength" || ans.DC != 10 {
		t.Errorf("unexpected test or dc: %+v", ans)
	}
	if ans.FailEffect["dexterity"] != -1 {
		t.Errorf("expected dexterity fail penalty -1, got %d", ans.FailEffect["dexterity"])
	}
}

func TestLoadCharacterData_FileNotFound(t *testing.T) {
	_, err := character.LoadCharacterData("nonexistent.json")
	if err == nil {
		t.Error("expected error for missing file, got nil")
	}
}

func TestLoadCharacterData_InvalidJSON(t *testing.T) {
	tmpDir := t.TempDir()
	jsonPath := filepath.Join(tmpDir, "bad.json")
	if err := os.WriteFile(jsonPath, []byte("not json"), 0644); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}
	_, err := character.LoadCharacterData(jsonPath)
	if err == nil {
		t.Error("expected error for invalid JSON, got nil")
	}
}

func TestGetAttributeIncreases(t *testing.T) {
	answer := character.Answer{
		Increases: map[string]int{
			"STR": 2,
			"DEX": 0, // Should be ignored
			"CON": 1,
			"BAD": 5, // Should be ignored (invalid short name)
		},
	}
	attrs := character.GetAttributeIncreases(answer)
	if len(attrs) != 2 {
		t.Errorf("expected 2 attributes, got %d", len(attrs))
	}
	if attrs[character.Str] != 2 {
		t.Errorf("expected STR=2, got %d", attrs[character.Str])
	}
	if attrs[character.Con] != 1 {
		t.Errorf("expected CON=1, got %d", attrs[character.Con])
	}
	if _, ok := attrs[character.Dex]; ok {
		t.Errorf("expected DEX to be omitted when value is 0")
	}
}

func TestGetAttributeFailEffects(t *testing.T) {
	answer := character.Answer{
		FailEffect: map[string]int{
			"WIS":  -2,
			"CHA":  0, // Should be ignored
			"INT":  -1,
			"NOPE": -5, // Should be ignored (invalid short name)
		},
	}
	attrs := character.GetAttributeFailEffects(answer)
	if len(attrs) != 2 {
		t.Errorf("expected 2 attributes, got %d", len(attrs))
	}
	if attrs[character.Wis] != -2 {
		t.Errorf("expected WIS=-2, got %d", attrs[character.Wis])
	}
	if attrs[character.Int] != -1 {
		t.Errorf("expected INT=-1, got %d", attrs[character.Int])
	}
	if _, ok := attrs[character.Cha]; ok {
		t.Errorf("expected CHA to be omitted when value is 0")
	}
}
