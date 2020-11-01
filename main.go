package main

import (
	"flag"
	"fmt"
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

	/// TODO : make flag for time
	correct, err := quiz.EvalQuiz(records, 10)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("\n<<<< You scored %v/%v. >>>>\n", correct, len(records))
}
