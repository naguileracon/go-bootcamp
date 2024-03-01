package average_calculator

func CalculateGradesAverage(grades ...float64) (gradeAverage float64) {
	for _, grade := range grades {
		gradeAverage += grade
	}
	if len(grades) > 0 {
		gradeAverage = gradeAverage / float64(len(grades))
		return
	}
	return
}
