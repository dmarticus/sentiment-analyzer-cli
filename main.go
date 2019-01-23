package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// dataset returns a map of sentences to their classes from a file
func dataset(file string) map[string]string {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	dataset := make(map[string]string)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		l := scanner.Text()
		data := strings.Split(l, "\t")
		if len(data) != 2 {
			continue
		}
		sentence := data[0]
		if data[1] == "0" {
			dataset[sentence] = negative
		} else if data[1] == "1" {
			dataset[sentence] = positive
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return dataset
}

func main() {
	// Initialize a new classifier
	nb := newClassifier()
	// Get dataset from a text file
	// Dataset can be downloaded from https://archive.ics.uci.edu/ml/datasets/Sentiment+Labelled+Sentences
	dataset := dataset("./sentiment_labelled_sentences/yelp_labelled.txt")
	// Train the classifier with dataset
	nb.train(dataset)

	// Prompt for inputs from console
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter your review: ")
		sentence, _ := reader.ReadString('\n')
		// Classify input sentence
		result := nb.classify(sentence)
		class := ""
		if result[positive] > result[negative] {
			class = positive
		} else {
			class = negative
		}
		fmt.Printf("> Your review is %s\n\n", class)
	}
}
