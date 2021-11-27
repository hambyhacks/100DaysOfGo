package cars

import "fmt"

// SuccessRate is used to calculate the ratio of an item being created without
// error for a given speed
func SuccessRate(speed int) float64 {
	var successRate float64
	switch speed {
	case 0:
		successRate = 0
	case 1, 2, 3, 4:
		successRate = 1.0
	case 5, 6, 7, 8:
		successRate = 0.9
	case 9, 10:
		successRate = 0.77
	default:
		fmt.Println("Invalid Range!")
	}
	return successRate
	panic("SuccessRate not implemented")
}

// CalculateProductionRatePerHour for the assembly line, taking into account
// its success rate
func CalculateProductionRatePerHour(speed int) float64 {
	var prodRate float64 = (221.0 * float64(speed)) * float64(SuccessRate(speed))
	return prodRate
	panic("CalculateProductionRatePerHour not implemented")
}

// CalculateProductionRatePerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateProductionRatePerMinute(speed int) int {
	var prodRateInMins int = int(CalculateProductionRatePerHour(speed)) / 60
	return prodRateInMins
	panic("CalculateProductionRatePerMinute not implemented")
}

// CalculateLimitedProductionRatePerHour describes how many working items are
// produced per hour with an upper limit on how many can be produced per hour
func CalculateLimitedProductionRatePerHour(speed int, limit float64) float64 {
	var prodLimitPerHour float64
	prodLimitPerHour = CalculateProductionRatePerHour(speed)
	if prodLimitPerHour >= limit {
		prodLimitPerHour = limit
	} else {
		prodLimitPerHour = prodLimitPerHour
	}
	return prodLimitPerHour
	panic("CalculateLimitedProductionRatePerHour not implemented")
}
