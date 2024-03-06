package calculator

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDivide(t *testing.T) {
	t.Run("success - happy path", func(t *testing.T) {
		// arrange - given
		n1 := 10
		n2 := 5
		expectedResult := 2
		expectedError := ""

		// act - when
		result, err := Divide(n1, n2)

		// assert - then
		if err != expectedError {
			t.Errorf("Expected error %s, got %s", expectedError, err)
			return
		}
		if result != expectedResult {
			t.Errorf("Expected result %d, got %d", expectedResult, result)
			return
		}
	})

	t.Run("error - zero division", func(t *testing.T) {
		// arrange - given

		// act - when

		// assert - then
	})

	t.Run("error - indeterminate division", func(t *testing.T) {
		// arrange
		n1 := 0
		n2 := 0
		expectedError := ErrMsgDivisionIndeterminate
		expectedResult := 0

		//act
		result, err := Divide(n1, n2)

		//assert
		if expectedError != err {
			t.Errorf("error unexpected error")
			return
		}

		if result != expectedResult {
			t.Errorf("error unexpected result")
			return
		}
	})
}

func TestMultiply(t *testing.T) {

	// arrange
	n1 := 10
	n2 := 5
	expectedResult := 50
	expectedError := ""

	//act
	result, err := Multiply(n1, n2)

	//assert
	require.Equal(t, expectedError, err)
	require.Equal(t, expectedResult, result)

}

func TestOrchestrator(t *testing.T) {


	t.Run("success - returns Add function", func(t *testing.T){
		op := OperatorSum
		
	}

}
