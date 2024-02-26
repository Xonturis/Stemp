/**
* This file contains the code for reading RAPL data.
 */

package rapl

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// readRAPL reads the content of the "energy_uj" file under the RAPL tree.
func readRAPL(dir string) ([]byte, error) {
	return os.ReadFile(filepath.Join(dir, "energy_uj"))
}

// convertRAPLReading converts the read data from []byte to uint64.
func convertRAPLReading(reading []byte) (uint64, error) {
	return strconv.ParseUint(strings.TrimSpace(string(reading)), 10, 64)
}

func GetRAPLDirs() ([]string, error) {
	// Get all RAPL domain directories under /sys/class/powercap/intel-rapl
	raplDirs, err := filepath.Glob("/sys/class/powercap/intel-rapl/intel-rapl:*")
	if err != nil {
		return nil, err
	}

	if len(raplDirs) == 0 {
		return nil, errors.New("no RAPL domains found")
	}
	return raplDirs, nil
}

// GetRAPLMeasurement reads RAPL measurements and returns the energy consumption of the CPU during
// the elapsed time.
func GetRAPLMeasurement(raplDirs []string) float64 {
	// Start time for elapsed calculation
	startTime := time.Now()

	// Initialize total energy
	totalEnergy := uint64(0)

	// Loop through each RAPL domain directory
	for _, raplDir := range raplDirs {
		// Open the energy file for this domain
		var content []byte
		var err error

		content, err = readRAPL(raplDir)
		if err != nil {
			fmt.Printf("Error opening %s: %v\n", raplDir, err)
			continue
		}

		// Read initial energy value
		initialEnergy, err := convertRAPLReading(content)
		if err != nil {
			fmt.Printf("Error reading energy from %s: %v\n", raplDir, err)
			continue
		}

		// Read final energy value after some time
		time.Sleep(time.Second * 1) // Adjust as needed

		content, err = readRAPL(raplDir)
		if err != nil {
			fmt.Printf("Error opening %s: %v\n", raplDir, err)
			continue
		}
		finalEnergy, err := convertRAPLReading(content)
		if err != nil {
			fmt.Printf("Error reading energy from %s: %v\n", raplDir, err)
			continue
		}

		// Add domain's energy to total
		totalEnergy += finalEnergy - initialEnergy
	}

	// Calculate elapsed time in seconds
	elapsedTime := time.Since(startTime).Seconds()

	// Calculate power consumption in watts
	powerConsumption := (float64(totalEnergy) / elapsedTime) / 1e6

	return powerConsumption
}
