package utils

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"
	"sync"
)

var Output []Utterance

func isCompleteSentence(s string) bool {
	return strings.HasSuffix(s, ".") || strings.HasSuffix(s, "?") || strings.HasSuffix(s, "!")
}
func isCapitalized(s byte) bool {
	res := 'A' <= s && s <= 'Z'

	return res
}

func midSpeakerChange(toAdd *Utterance, current *Utterance) {

	breakIndexToAdd := -1
	txt := ""
	// Iterate the string from the end
	for i := len(toAdd.Text) - 1; i >= 0; i-- {
		if toAdd.Text[i] == '.' || toAdd.Text[i] == '!' || toAdd.Text[i] == '?' {
			// Found a full stop
			breakIndexToAdd = i
			break
		}
		txt = string(toAdd.Text[i]) + txt
	}
	//construct two utterances
	toAdd.Text = toAdd.Text[0 : breakIndexToAdd+1]
	current.Text = txt + current.Text

}

func ReadUtterances(DirToRead string, fileNames []os.DirEntry, uChan chan<- Utterance, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(uChan)
	for _, file := range fileNames {

		fileName := file.Name()

		f, err := os.Open(DirToRead + "/" + fileName)

		if err != nil {
			panic(err)
		}

		dt, _ := io.ReadAll(f)

		var u Utterance
		if err := json.Unmarshal(dt, &u); err != nil {
			panic(err)
		}

		uChan <- u

		f.Close()

	}

}

func GenerateUtterances(uChan <-chan Utterance, wg *sync.WaitGroup) {
	defer wg.Done()
	var toAdd Utterance

	for u := range uChan {
		if u.Speaker == toAdd.Speaker || toAdd.Speaker == "" {
			statement := toAdd.Text + u.Text + " "
			toAdd = u
			toAdd.Text = statement
		} else {
			if !isCapitalized(u.Text[0]) && isCompleteSentence(u.Text) {
				// mid speaker change event condition
				midSpeakerChange(&toAdd, &u)
				Output = append(Output, toAdd)
				toAdd = u

			} else if u.Partial {
				statement := toAdd.Text + " " + u.Text + " "
				toAdd.Text = statement

			} else {
				// if speakers NOT equal but current complete sentence end-to-end
				Output = append(Output, toAdd)
				toAdd = u
			}

		}
	}

	Output = append(Output, toAdd)
}

func CreateCompleteUtterances(filename string) error {

	completeUtterancesJSON, err := json.MarshalIndent(Output, "", "    ")
	if err != nil {
		return err
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write data to the file
	_, err = file.Write(completeUtterancesJSON)
	if err != nil {

		return err
	}

	log.Println("Data has been written to complete.json")
	return nil

}
