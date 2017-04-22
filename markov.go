package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
)

type Markov struct{}

/*type NextState struct {
	Word  string
	Count int
	Prob  float64
}*/
type State struct {
	Word       string
	Count      int
	Prob       float64
	NextStates []State
}

var markov Markov

func addWordToStates(states []State, word string) ([]State, int) {
	iState := -1
	for i := 0; i < len(states); i++ {
		if states[i].Word == word {
			iState = i
		}
	}
	if iState >= 0 {
		states[iState].Count++
	} else {
		var tempState State
		tempState.Word = word
		tempState.Count = 1

		states = append(states, tempState)
		iState = len(states) - 1

	}
	return states, iState
}

func calcMarkovStates(words []string) []State {
	var states []State
	//count words
	for i := 0; i < len(words)-1; i++ {
		var iState int
		states, iState = addWordToStates(states, words[i])
		if iState < len(words) {
			states[iState].NextStates, _ = addWordToStates(states[iState].NextStates, words[i+1])
		}
	}

	//count prob
	for i := 0; i < len(states); i++ {
		states[i].Prob = (float64(states[i].Count) / float64(len(words)) * 100)
		for j := 0; j < len(states[i].NextStates); j++ {
			states[i].NextStates[j].Prob = (float64(states[i].NextStates[j].Count) / float64(len(words)) * 100)
		}
	}
	fmt.Println("total words: " + strconv.Itoa(len(words)))
	//fmt.Println(states)
	return states
}

func textToWords(text string) []string {
	s := strings.Split(text, " ")
	words := s
	return words
}

func readText(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		//Do something
	}
	content := string(data)
	return content, err
}

func (markov Markov) train(firstWord string, path string) []State {
	text, _ := readText(path)
	words := textToWords(text)
	states := calcMarkovStates(words)
	//fmt.Println(states)

	return states
}

func getNextMarkovState(states []State, word string) string {
	iState := -1
	for i := 0; i < len(states); i++ {
		if states[i].Word == word {
			iState = i
		}
	}
	if iState < 0 {
		return "word no exist on the memory"
	}
	fmt.Println(rand.Float64())
	return word
}
func (markov Markov) generateText(states []State, initWord string, count int) string {
	var generatedText []string
	word := initWord
	for i := 0; i < count; i++ {
		word = getNextMarkovState(states, word)
		generatedText = append(generatedText, word)
	}
	return "a"
}
