package cmd

import (
	"fmt"

	"github.com/ocobobyte/temperature-classifier/temp"
)

// Fixed array for the days of the week
var days = [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

// Runs the command to calculate the average temperature and classify each day
func Execute() {
	// Create a slice to store temperatures
	temps := make([]float64, len(days))

	// Read the temperature for each day of the week
	for i, day := range days {
		fmt.Printf("Enter the temperature for %s: ", day)
		fmt.Scanf("%f", &temps[i])
	}

	// Calculate average temperature
	avg := temp.Average(temps)
	fmt.Printf("Average temperature: %.2f°C\n", avg)

	// Classify each day
	hotDays, coldDays := temp.ClassifyDays(temps, days[:])
	fmt.Println("Hot Days:", hotDays)
	fmt.Println("Cold Days:", coldDays)
}