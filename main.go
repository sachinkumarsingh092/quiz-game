package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/sachinkumarsingh092/quiz-game/quiz"
)

func main() {
	// Flags
	filename := flag.String("filename", "./problems/problems.csv", "To use a given CSV file for quiz.")
	timeLimit := flag.Int("time", 30, "Time limit for the quiz.")
	flag.Parse()

	records, err := quiz.ReadCsvFile(*filename)
	if err != nil {
		log.Fatal("no such file.")
	}

	correct, err := quiz.EvalQuiz(records, *timeLimit)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("\n<<<< You scored %v/%v. >>>>\n", correct, len(records))
}
