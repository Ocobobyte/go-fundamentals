package temp

import "fmt"

// Calculates the average of a slice of temperatures
func Average(temps []float64) float64 {
	sum := 0.0
	for _, temp := range temps {
		sum += temp
	}
	return sum / float64(len(temps))
}

// Takes a slice of temperatures and corresponding days and returns two slices: one for hot days (temperature > 25°C) and one for cold days.
func ClassifyDays(temps []float64, days []string) (hotDays []string, coldDays []string) {
	for i, t := range temps {
		entry := fmt.Sprintf("%s (%.2f°C)", days[i], t)
		if t > 25.0 {
			hotDays = append(hotDays, entry)
		} else {
			coldDays = append(coldDays, entry)
		}
	}
	return
}
