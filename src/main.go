package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/xonturis/stemp/src/logger"
	"github.com/xonturis/stemp/src/temperature"
)

// RAPL ENERGY
// func main() {
// 	raplDirs, err := rapl.GetRAPLDirs()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	file, err := logger.CreateCSVFile()
// 	if err != nil {
// 		fmt.Println("Error creating CSV file:", err)
// 		return
// 	}
// 	defer file.Close()

// 	currentTime := time.Now()

// 	for {
// 		powerConsumption := rapl.GetRAPLMeasurement(raplDirs)

// 		newTime := time.Now()
// 		if currentTime.Day() != newTime.Day() {
// 			file.Close()

// 			file, err = logger.CreateCSVFile()
// 			if err != nil {
// 				fmt.Println("Error creating CSV file:", err)
// 				return
// 			}
// 		}
// 		currentTime := time.Now()

//
// 		data := []string{strconv.FormatInt(currentTime.Unix(), 10), strconv.FormatFloat(powerConsumption, 'f', -1, 64)}

// 		// Insert data into the CSV file
// 		err = logger.InsertDataToCSV(file, data)
// 		if err != nil {
// 			fmt.Println("Error inserting data into CSV file:", err)
// 			return
// 		}
// 	}
// }

func main() {
	file, err := logger.CreateCSVFile()
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer file.Close()

	zones, err := temperature.GetThermalZones()
	if err != nil {
		fmt.Println("Error getting thermal zones:", err)
		return
	}

	currentTime := time.Now()

	for {
		newTime := time.Now()
		if currentTime.Day() != newTime.Day() {
			file.Close()
			file, err = logger.CreateCSVFile()
			if err != nil {
				fmt.Println("Error creating CSV file:", err)
				return
			}
		}
		currentTime := time.Now()

		meanTemp, err := temperature.GetMeanTemperature(zones)
		if err != nil {
			fmt.Println("Error calculating mean temperature:", err)
			return
		}

		data := []string{strconv.FormatInt(currentTime.Unix(), 10), strconv.FormatFloat(meanTemp, 'f', -1, 64)}

		// Insert data into the CSV file
		err = logger.InsertDataToCSV(file, data)
		if err != nil {
			fmt.Println("Error inserting data into CSV file:", err)
			return
		}

		time.Sleep(10 * time.Second)
	}
}
