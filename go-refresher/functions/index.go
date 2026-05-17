package functions

import (
	"errors"
)

func Add(firstNum int, secondNum int) int {
	return firstNum + secondNum
}

func Subtract(firstNum int, secondNum int, isSubtractFirstFromSecond bool) int {
	if isSubtractFirstFromSecond {
		return secondNum - firstNum
	} else {
		return firstNum - secondNum
	}
}

func Multiply(firstNum int, secondNum int) int {
	return firstNum * secondNum

}

func Divide(firstNum int, secondNum int, isDivideFirsBySecond bool) (float32, error) {
	if isDivideFirsBySecond {
		if secondNum != 0 {
			return float32(firstNum / secondNum), nil
		} else {
			return 0, errors.New("Division by zero is not allowed")
		}
	} else {
		if firstNum != 0 {
			return float32(secondNum / firstNum), nil
		}
		return 0, errors.New("Division by zero is not allowed")
	}
}

func DivideWithRemainder(firstNum int, secondNum int, isFirstModuleBySecond bool) (int, int, error) {
	if isFirstModuleBySecond {
		if secondNum != 0 {
			return firstNum / secondNum, firstNum % secondNum, nil
		} else {
			return 0, 0, errors.New("Division by zero is not allowed")
		}
	} else {
		if firstNum != 0 {
			return secondNum / firstNum, secondNum % firstNum, nil
		} else {
			return 0, 0, errors.New("Division by zero is not allowed")
		}
	}

}

//  experment the named return values and that of variadic functions
// 1. variadic functions

func Sum(nums ...int) int {

	total := 0
	for _, currNum := range nums {
		total += currNum
	}

	return total

}

func GetTotalMessagesFrorAllUsers(nums ...int) (totalMesssages int) {
	totalMesssages = Sum(nums...)
	return
}
