package quiz

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
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

// ParseQuiz : Run the quiz and evaluate.
func ParseQuiz(records [][]string) {
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
