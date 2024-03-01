package statistics

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func CalculateStatistics(operationType string, grades ...int) (float64, string) {
	operation, err := chooseOperation(operationType)
	if err != "" {
		return 0, err
	}
	result := operation(grades...)
	return result, ""
}

func chooseOperation(operationType string) (func(values ...int) float64, string) {
	switch operationType {
	case maximum:

		return findMaxValue, ""
	case average:

		return calculateAverage, ""
	case minimum:

		return findMinValue, ""
	default:
		return nil, "Invalid operation"
	}
}

func findMaxValue(values ...int) float64 {
	if len(values) == 0 {
		return 0
	}
	maxValue := values[0]
	for _, value := range values {
		if value > maxValue {
			maxValue = value
		}
	}
	return float64(maxValue)
}

func findMinValue(values ...int) float64 {
	if len(values) == 0 {
		return 0
	}
	minValue := values[0]
	for _, value := range values {
		if value < minValue {
			minValue = value
		}
	}
	return float64(minValue)
}

func calculateAverage(values ...int) float64 {
	if len(values) == 0 {
		return 0.0
	}
	sum := 0
	for _, value := range values {
		sum += value
	}
	averageValue := float64(sum) / float64(len(values))
	return averageValue
}
