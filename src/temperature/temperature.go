package temperature

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// GetThermalZones returns a list of thermal zone directories.
func GetThermalZones() ([]string, error) {
	return filepath.Glob("/sys/class/thermal/thermal_zone*")
}

// convertTempReading converts the read data from []byte to uint64.
func convertTempReading(reading []byte) (uint64, error) {
	return strconv.ParseUint(strings.TrimSpace(string(reading)), 10, 64)
}

// ReadTemperature reads the temperature from a thermal zone file.
func ReadTemperature(zonePath string) (float64, error) {
	tempFile := filepath.Join(zonePath, "temp")
	tempBytes, err := os.ReadFile(tempFile)
	if err != nil {
		return 0, err
	}

	tempInt, err := convertTempReading(tempBytes)
	fmt.Println(tempInt)
	if err != nil {
		return 0, err
	}

	return float64(tempInt) / 1000.0, nil
}

// GetMeanTemperature calculates the average temperature across all zones.
func GetMeanTemperature(zones []string) (float64, error) {
	var totalTemp float64
	for _, zone := range zones {
		temp, err := ReadTemperature(zone)
		if err != nil {
			return 0, err
		}
		totalTemp += temp
	}

	meanTemp := totalTemp / float64(len(zones))
	return meanTemp, nil
}
