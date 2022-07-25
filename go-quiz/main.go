package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"gophercises/go-quiz/quiz"
	"log"
	"os"
	"time"
)

var (
	csvf     string
	duration int
	shuffle  bool
)

func init() {
	flag.StringVar(&csvf, "quiz", "", "A .csv file with questions and answers.")
	flag.IntVar(&duration, "time", 30, "The duration of the entire quiz(in seconds).")
	flag.BoolVar(&shuffle, "shuffle", false, "Shuffle the order of the quiz")
	flag.Parse()
}

func main() {
	if csvf == "" {
		csvf = "problems.csv"
	}
	file, err := os.Open(csvf)
	if err != nil {
		log.Fatalf("Failed to open file: %v\n", err)
	}
	defer file.Close()

	csvr := csv.NewReader(file)
	content, err := csvr.ReadAll()
	if err != nil {
		log.Fatalf("Failed to read file: %v\n", err)
	}

	quizzes := quiz.ParseContent(content, shuffle)

	timeoutCh := time.After(time.Duration(duration) * time.Second)
	scoreCh := make(chan int)
	score := 0

	go func() {
		for i, q := range quizzes {
			fmt.Printf("Question %d:\t%s = ", i+1, q.Question)
			var ans string
			fmt.Scanln(&ans)
			if ans == q.Answer {
				score++
			}
		}
		scoreCh <- score
	}()

	select {
	case <-scoreCh:
		fmt.Printf("\nYou scored %d out of %d questions\n", score, len(quizzes))
	case <-timeoutCh:
		fmt.Printf("\n\nTimeout!\nYou scored %d out of %d questions\n", score, len(quizzes))
	}
}
