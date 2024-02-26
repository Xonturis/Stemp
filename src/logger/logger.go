/**
* This file handles the csv output of the probe.
 */

package logger

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

// openFileIfExists checks if a file exists at the given path.
// If the file exists, it returns a pointer to the file and nil error.
// If the file does not exist, it returns nil pointer and an error.
func openFileIfExists(filename string) (*os.File, error) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if os.IsNotExist(err) {
		return nil, err
	}
	return file, nil
}

// CreateCSVFile creates a CSV file with the current date in name, and (timestamp, temperature) as header.
func CreateCSVFile() (*os.File, error) {
	// Get current date
	currentTime := time.Now()
	date := currentTime.Format("02012006") // "ddMMYYYY"

	// Create the file name
	filename := fmt.Sprintf("%s-Stemp.csv", date)

	// Create the CSV file
	header := []string{"timestamp", "temperature"}

	file, _ := openFileIfExists(filename)

	if file != nil {
		return file, nil
	}

	file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(header)
	if err != nil {
		return nil, err
	}

	return file, nil
}

// InsertDataToCSV inserts data into the CSV file.
func InsertDataToCSV(file *os.File, data []string) error {
	writer := csv.NewWriter(file)
	defer writer.Flush()

	err := writer.Write(data)
	if err != nil {
		return err
	}

	return nil
}
