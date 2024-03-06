package salaries

func CalculateSalaryTaxes(salary float64) float64 {
	switch {
	case salary <= 50000:
		return 0
	case salary > 50000 && salary <= 150000:
		return salary * 0.17
	default:
		return salary * 0.27
	}
}

func CalculateSalary(minutesWorked int, category string) (salary float64, err string) {
	var hoursWorked = float64(minutesWorked / 60)
	switch category {
	case "A":
		monthlySalary := 3000 * hoursWorked
		salary = monthlySalary + (monthlySalary * 0.5)
		return
	case "B":
		monthlySalary := 1500 * hoursWorked
		salary = monthlySalary + (monthlySalary * 0.2)
		return
	case "C":
		salary = 1000 * hoursWorked
		return
	default:
		err = "Invalid category"
		return
	}

}
