package statistics

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCalculateStatistics(t *testing.T) {
	t.Run("success - calculate maximum", func(t *testing.T) {
		// arrange - given
		grades := []int{1, 2, 3, 4, 5}
		operation := "maximum"
		expectedResult := float64(5)
		expectedError := ""
		// act - when
		result, err := CalculateStatistics(operation, grades...)
		// assert then
		require.Equal(t, expectedError, err)
		require.Equal(t, expectedResult, result)
	})

	t.Run("success - calculate minimum", func(t *testing.T) {
		// arrange - given
		grades := []int{1, 2, 3, 4, 5}
		operation := "minimum"
		expectedResult := float64(1)
		expectedError := ""
		// act - when
		result, err := CalculateStatistics(operation, grades...)
		// assert then
		require.Equal(t, expectedError, err)
		require.Equal(t, expectedResult, result)
	})

	t.Run("success - calculate average", func(t *testing.T) {
		// arrange - given
		grades := []int{1, 2, 3, 4, 5}
		operation := "average"
		sum := 0.0
		for _, grade := range grades {
			sum += float64(grade)
		}
		expectedResult := sum / float64(len(grades))
		expectedError := ""
		// act - when
		result, err := CalculateStatistics(operation, grades...)
		// assert then
		require.Equal(t, expectedError, err)
		require.Equal(t, expectedResult, result)
	})

	t.Run("failure - invalid operation", func(t *testing.T) {
		// arrange - given
		grades := []int{1, 2, 3, 4, 5}
		operation := "division"
		expectedResult := float64(0)
		expectedError := "Invalid operation"
		// act - when
		result, err := CalculateStatistics(operation, grades...)
		// assert then
		require.Equal(t, expectedError, err)
		require.Equal(t, expectedResult, result)
	})

}
