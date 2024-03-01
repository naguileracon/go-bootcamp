package average_calculator

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCalculateGradesAverage(t *testing.T) {
	t.Run("success - calculate grades average", func(t *testing.T) {
		// arrange - given
		grades := []float64{5.8, 5.9, 10.6, 15.6, 18.9}
		sum := 0.0
		for _, grade := range grades {
			sum += grade
		}
		expectedResult := sum / float64(len(grades))
		// act - when
		result := CalculateGradesAverage(grades...)
		// assert then
		require.Equal(t, expectedResult, result)
	})

}
