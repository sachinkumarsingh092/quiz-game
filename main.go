package main

import (
	"flag"
	"log"

	"github.com/sachinkumarsingh092/quiz-game/quiz"
)

func main() {
	flag.Parse()
	filename := flag.Arg(0)
	if filename == "" {
		filename = "./problems/problems.csv"
	}

	records, err := quiz.ReadCsvFile(filename)
	if err != nil {
		log.Fatal("no such file.")
	}

	quiz.ParseQuiz(records)
}
