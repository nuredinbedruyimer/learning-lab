package conditions

import (
	"errors"
)

func Divide(firstNum int, secondNum int) (float32, error) {
	if secondNum != 0 {
		return float32(firstNum / secondNum), nil
	} else {
		return 0, errors.New("Division by zero is not allowed")
	}
}

func WhatIsTheDayOfTheWeek(dayNum int) (string, error) {
	switch dayNum {
	case 1:
		return "Monday", nil
	case 2:
		return "Tuesday", nil
	case 3:
		return "Wednesday", nil
	case 4:
		return "Thursday", nil
	case 5:
		return "Friday", nil
	case 6:
		return "Saturday", nil
	case 7:
		return "Sunday", nil
	default:
		return "", errors.New("Invalid day number. Please provide a number between 1 and 7.")
	}
}
