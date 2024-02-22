package utils

import (
	"encoding/json"
	"os"
	"sync"
	"testing"
)

func TestIsCompleteSentence(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"This is a complete sentence.", true},
		{"This is an incomplete sentence", false},
		{"What?", true},
	}

	for _, test := range tests {
		if got := isCompleteSentence(test.input); got != test.expected {
			t.Errorf("isCompleteSentence(%q) = %v; expected %v", test.input, got, test.expected)
		}
	}
}

// Write similar test cases for isCapitalized, midSpeakerChange, and other functions...

func TestReadUtterances(t *testing.T) {
	// Create a temporary directory and files for testing
	tmpDir := t.TempDir()

	tmpFile1, err := os.CreateTemp(tmpDir, "*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile1.Name())

	// Write JSON data to the temporary file
	jsonData := `{"speaker": "Alice", "text": "Hello", "timestampMs": 123, "isPartial": false}`
	if _, err := tmpFile1.Write([]byte(jsonData)); err != nil {
		t.Fatal(err)
	}

	// Call ReadUtterances and verify the output
	var wg sync.WaitGroup
	uChan := make(chan Utterance)

	fakeDir, _ := os.ReadDir(tmpDir)
	wg.Add(1)

	go ReadUtterances(tmpDir, fakeDir, uChan, &wg)

	utterance := <-uChan
	wg.Wait()

	expected := Utterance{Speaker: "Alice", Text: "Hello", TimestampMs: 123, Partial: false}
	if utterance != expected {
		t.Errorf("ReadUtterances() produced unexpected output. Got %+v, expected %+v", utterance, expected)
	}
}

// Write similar test cases for other functions...

func TestCreateCompleteUtterances(t *testing.T) {
	// Prepare test data
	Output = []Utterance{
		{Speaker: "Alice", Text: "Hello", TimestampMs: 123, Partial: false},
		{Speaker: "Bob", Text: "How are you?", TimestampMs: 456, Partial: false},
	}

	// Call createCompleteUtterances
	filename := "test_output.json"
	if err := CreateCompleteUtterances(filename); err != nil {
		t.Fatalf("createCompleteUtterances() returned an unexpected error: %v", err)
	}
	defer os.Remove(filename)

	// Read the written file and verify its contents
	content, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to read the output file: %v", err)
	}

	var output []Utterance
	if err := json.Unmarshal(content, &output); err != nil {
		t.Fatalf("Failed to unmarshal JSON from the output file: %v", err)
	}

	if len(output) != len(Output) {
		t.Fatalf("Mismatch in the number of utterances. Expected %d, got %d", len(Output), len(output))
	}

	// Compare each utterance
	for i := range Output {
		if output[i] != Output[i] {
			t.Errorf("Mismatch in utterance at index %d. Expected %+v, got %+v", i, Output[i], output[i])
		}
	}
}
