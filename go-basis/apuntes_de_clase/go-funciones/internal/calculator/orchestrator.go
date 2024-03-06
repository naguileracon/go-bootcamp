package calculator

func Orchestrator(operator MathOperator) (mo MathOperation, err string) {
	switch operator {
	case OperatorSum:
		mo = Add
	case OperatorSub:
		mo = Substract
	case OperatorMul:
		mo = Multiply
	case OperatorDiv:
		mo = Divide
	default:
		err = "error: invalid operator"
	}
	return
}
