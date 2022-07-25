package quiz

import (
	"math/rand"
	"strings"
	"time"
)

type Quiz struct {
	Question string
	Answer   string
}

// Takes in a [][]string and returns a []Quiz
// When shuffle is true the returned []Quiz is in random order
func ParseContent(content [][]string, shuffle bool) []Quiz {
	quizzes := make([]Quiz, len(content))
	for i, v := range content {
		quizzes[i] = Quiz{Question: v[0], Answer: strings.TrimSpace(v[1])}
	}
	if shuffle {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(quizzes), func(i, j int) {
			quizzes[i], quizzes[j] = quizzes[j], quizzes[i]
		})
	}
	return quizzes
}

// Checks a []Quiz for a Quiz and return a bool
func contains(haystack []Quiz, needle Quiz) bool {
	for _, matchValue := range haystack {
		if matchValue == needle {
			return true
		}
	}
	return false
}
