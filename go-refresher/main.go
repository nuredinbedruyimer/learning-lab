package main

import (
	"fmt"
	"go_refresher/conditions"
)

func main() {

	if result, err := conditions.Divide(12, 0); err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Result: ", result)
	}

	if today, errMessage := conditions.WhatIsTheDayOfTheWeek(12); errMessage != nil {
		fmt.Println("Error: ", errMessage)
	} else {
		fmt.Println("Today is: ", today)
	}

}
