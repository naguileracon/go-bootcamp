package salaries

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCalculateSalaryTaxes(t *testing.T) {
	t.Run("success - salary lower than 50000", func(t *testing.T) {
		// arrange - given
		salary := 48000.0
		expectedResult := 0.0
		// act - when
		result := CalculateSalaryTaxes(salary)
		// assert then
		require.Equal(t, expectedResult, result)
	})

	t.Run("success - salary between 50000 and 150000", func(t *testing.T) {
		// arrange - given
		salary := 67000.0
		expectedResult := salary * 0.17
		// act - when
		result := CalculateSalaryTaxes(salary)
		// assert then
		require.Equal(t, expectedResult, result)
	})

	t.Run("success - salary greater than 150000", func(t *testing.T) {
		// arrange - given
		salary := 170000.0
		expectedResult := salary * 0.27
		// act - when
		result := CalculateSalaryTaxes(salary)
		// assert then
		require.Equal(t, expectedResult, result)
	})

}

func TestCalculateSalary(t *testing.T) {
	t.Run("success - salary category A", func(t *testing.T) {
		// arrange - given
		minutesWorked := 60
		hoursWorked := float64(minutesWorked / 60)
		category := "A"
		monthlySalary := 3000 * hoursWorked
		expectedResult := monthlySalary + (monthlySalary * 0.5)
		expectedErr := ""
		// act - when
		result, err := CalculateSalary(minutesWorked, category)
		// assert then
		require.Equal(t, expectedErr, err)
		require.Equal(t, expectedResult, result)
	})

	t.Run("success - salary category B", func(t *testing.T) {
		// arrange - given
		minutesWorked := 60
		hoursWorked := float64(minutesWorked / 60)
		category := "B"
		monthlySalary := 1500 * hoursWorked
		expectedResult := monthlySalary + (monthlySalary * 0.2)
		expectedErr := ""
		// act - when
		result, err := CalculateSalary(minutesWorked, category)
		// assert then
		require.Equal(t, expectedErr, err)
		require.Equal(t, expectedResult, result)
	})

	t.Run("success - salary category C", func(t *testing.T) {
		// arrange - given
		minutesWorked := 60
		hoursWorked := float64(minutesWorked / 60)
		category := "C"
		expectedResult := 1000 * hoursWorked
		expectedErr := ""
		// act - when
		result, err := CalculateSalary(minutesWorked, category)
		// assert then
		require.Equal(t, expectedErr, err)
		require.Equal(t, expectedResult, result)
	})

	t.Run("failed - invalid category", func(t *testing.T) {
		// arrange - given
		minutesWorked := 0
		category := "D"
		expectedResult := float64(0)
		expectedErr := "Invalid category"
		// act - when
		result, err := CalculateSalary(minutesWorked, category)
		// assert then
		require.Equal(t, expectedErr, err)
		require.Equal(t, expectedResult, result)
	})

}
