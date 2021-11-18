package cars

const CapacityPerHour int = 221

// SuccessRate is used to calculate the ratio of an item being created without
// error for a given speed
func SuccessRate(speed int) float64 {
	var rate = 0.0

	if speed == 0 {
		rate = 0
	} else if speed <= 4 {
		rate = 1
	} else if speed <= 8 {
		rate = 0.9
	} else if speed <= 10 {
		rate = 0.77
	}
	return rate
}

// CalculateProductionRatePerHour for the assembly line, taking into account
// its success rate
func CalculateProductionRatePerHour(speed int) float64 {
	return float64(CapacityPerHour*speed) * SuccessRate(speed)
}

// CalculateProductionRatePerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateProductionRatePerMinute(speed int) int {
	return int(CalculateProductionRatePerHour(speed) / 60)
}

// CalculateLimitedProductionRatePerHour describes how many working items are
// produced per hour with an upper limit on how many can be produced per hour
func CalculateLimitedProductionRatePerHour(speed int, limit float64) float64 {
	var production = CalculateProductionRatePerHour(speed)
	if production > limit {
		return limit
	} else {
		return production
	}
}
