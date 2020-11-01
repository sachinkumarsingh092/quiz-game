package quiz

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// ReadCsvFile : Read a given CSV file.
func ReadCsvFile(filePath string) ([][]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(filePath, err)
	}

	return records, nil
}

// getInput : Send answers to the channels.
func getInput(ch chan string) {
	for {
		in := bufio.NewReader(os.Stdin)
		input, err := in.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		ch <- input
	}
}

// EvalQuiz : Evaluate and score the quiz.
func EvalQuiz(records [][]string, t int) (int, error) {
	total := 0
	timer := time.NewTimer(time.Duration(t) * time.Second)
	answer := make(chan string)

	go getInput(answer)

	for i := 0; i < len(records); i++ {
		score, err := runQuiz(records[i][0], records[i][1], answer, timer.C)
		if err != nil && score == -1 {
			return total, err
		}

		total += score
	}

	return total, nil
}

// runQuiz : Run the quiz.
func runQuiz(question, answer string, ansChan <-chan string, quit <-chan time.Time) (int, error) {
	fmt.Println(question)

	for {
		select {
		case <-quit:
			return -1, fmt.Errorf("Time out")

		case input := <-ansChan:
			score := 0
			if strings.Compare(strings.Trim(strings.ToLower(input), "\n"), answer) == 0 {
				score = 1
			} else {
				return 0, fmt.Errorf("Wrong Answer")
			}
			return score, nil
		}
	}
}
