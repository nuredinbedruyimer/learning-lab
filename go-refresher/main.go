package main

import (
	"fmt"
	"go_refresher/functions"
)

func main() {
	firstNum := 12
	secondNum := 5
	var additionResult int = functions.Add(firstNum, secondNum)
	fmt.Println(" Addtion = ", additionResult)

	subtractionResult := functions.Subtract(firstNum, secondNum, true)
	fmt.Println(" Subtraction = ", subtractionResult)

	multiplicationResult := functions.Multiply(firstNum, secondNum)
	fmt.Println(" Multiplication = ", multiplicationResult)
	divisionResult, err := functions.Divide(firstNum, secondNum, true)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(" Division = ", divisionResult)
	}

	divisionWithRemainderResult, remainder, err := functions.DivideWithRemainder(firstNum, 0, true)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(" Division with Remainder = ", divisionWithRemainderResult, " Remainder = ", remainder)
	}

}
