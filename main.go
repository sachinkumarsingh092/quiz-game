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
		filename = "problems.csv"
	}

	records, err := quiz.ReadCsvFile(filename)
	if err != nil {
		log.Fatal("no such file.")
	}

	var correct int = 0
	for i := 0; i < len(records); i++ {
		fmt.Printf("What's %s equal to?\n", records[i][0])

		var answer string
		fmt.Scanln(&answer)

		if answer == records[i][1] {
			correct++
		}
	}

	fmt.Printf("\n<<<< You scored %v/%v. >>>>\n", correct, len(records))
}
