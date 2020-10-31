package quiz

import (
	"encoding/csv"
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
