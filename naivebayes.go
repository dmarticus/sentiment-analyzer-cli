package main

// The string values of the 2 classes
// They can be "positive" >< "negative" as in this example
// They can also be "ham" >< "spam", i.e.
const (
	positive = "positive"
	negative = "negative"
)

/*
 * Classifier
 */

// wordFrequency stores frequency of words. For example:
// wordFrequency{
//      word: "excellent"
//	counter: map[string]int{
//		"positive": 15
//		"negative": 0
//	}
// }
type wordFrequency struct {
	word    string
	counter map[string]int
}

// classifier can be trained and used to categorize objects
// Attributes:
//	dataset: map each class with a list of  sentences from training data
//		map[string][]string{
//			"positive": []string{
//				"The restaurant is excellent",
//				"I really love this restaurant",
//			},
//			"negative": []string{
//				"Their food is awful",
//			}
//
//		}
//	words: map each word with their frequency
//		map[string]wordFrequency{
//			"restaurant": wordFrequency{
//				word: "restaurant"
//				counter: map[string]int{
//					"positive": 2
//					"negative": 0
//				}
// 			}
//		}
type classifier struct {
	dataset map[string][]string
	words   map[string]wordFrequency
}

// newClassifier returns a new classifier with empty dataset and words
func newClassifier() *classifier {
	c := new(classifier)
	c.dataset = map[string][]string{
		positive: []string{},
		negative: []string{},
	}
	c.words = map[string]wordFrequency{}
	return c
}

// train populates a classifier's dataset and words with input dataset map
// Sample dataset: map[string]string{
//	"The restaurant is excellent": "Positive",
//	"I really love this restaurant": "Positive",
//	"Their food is awful": "Negative",
//}
func (c *classifier) train(dataset map[string]string) {
	for sentence, class := range dataset {
		c.addSentence(sentence, class)
		words := tokenize(sentence)
		for _, w := range words {
			c.addWord(w, class)
		}
	}
}

// classify return the probablitities of a sentence being each class
// Sample @return map[string]float64 {
//	"positive": 0.7,
//	"negative": 0.1,
//}
// Meaning 70% chance the input sentence is positive, 10% it's negative
func (c classifier) classify(sentence string) map[string]float64 {
	words := tokenize(sentence)
	posProb := c.probability(words, positive)
	negProb := c.probability(words, negative)
	return map[string]float64{
		positive: posProb,
		negative: negProb,
	}
}

// addSentence adds a sentence and its class to a classifier's dataset map
func (c *classifier) addSentence(sentence, class string) {
	c.dataset[class] = append(c.dataset[class], sentence)
}

// addSentence adds a word to a classifier's words map and update its frequency
func (c *classifier) addWord(word, class string) {
	wf, ok := c.words[word]
	if !ok {
		wf = wordFrequency{word: word, counter: map[string]int{
			positive: 0,
			negative: 0,
		}}
	}
	wf.counter[class]++
	c.words[word] = wf
}

// priorProb returns the prior probability of each class of the classifier
// This probability is determined purely by the training dataset
func (c classifier) priorProb(class string) float64 {
	return float64(len(c.dataset[class])) / float64(len(c.dataset[positive])+len(c.dataset[negative]))
}

// totalWordCount returns the word count of a class (duplicated also count)
// If class provided is not positive or negative, it returns
// the total word count in dataset.
func (c classifier) totalWordCount(class string) int {
	posCount := 0
	negCount := 0
	for _, wf := range c.words {
		posCount += wf.counter[positive]
		negCount += wf.counter[negative]
	}
	if class == positive {
		return posCount
	} else if class == negative {
		return negCount
	} else {
		return posCount + negCount
	}
}

// totalDistinctWordCount returns the number of distinct words in dataset
func (c classifier) totalDistinctWordCount() int {
	posCount := 0
	negCount := 0
	for _, wf := range c.words {
		posCount += zeroOneTransform(wf.counter[positive])
		negCount += zeroOneTransform(wf.counter[negative])
	}
	return posCount + negCount
}

// probability retuns the probability of a list of words being in a class
func (c classifier) probability(words []string, class string) float64 {
	prob := c.priorProb(class)
	for _, w := range words {
		count := 0
		if wf, ok := c.words[w]; ok {
			count = wf.counter[class]
		}
		prob *= (float64((count + 1)) / float64((c.totalWordCount(class) + c.totalDistinctWordCount())))
	}
	for _, w := range words {
		count := 0
		if wf, ok := c.words[w]; ok {
			count += (wf.counter[positive] + wf.counter[negative])
		}
		prob /= (float64((count + 1)) / float64((c.totalWordCount("") + c.totalDistinctWordCount())))
	}
	return prob
}
