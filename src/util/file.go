package util

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

func CSVToSliceStr(path string) []string {
	s := []string{}

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return nil
	}

	for _, record := range records {
		s = append(s, record...)
	}

	return s
}

func WriteSliceStrToCSV(pathFile string, name string, slice []string, suffixTime bool) error {

	fileName := pathFile + name
	if suffixTime {
		fileName += time.Now().Format("_2006-01-02_15-04-05")

	}

	fileName += ".csv"

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Could not create %s", fileName)
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, v := range slice {
		writer.Write([]string{v})
	}

	return nil
}
