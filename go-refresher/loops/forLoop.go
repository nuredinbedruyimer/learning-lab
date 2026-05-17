package loops

import "fmt"

func TestForLoop() {
	/*
		for intialization; condition; update{
			}
	*/

	for index := 0; index <= 5; index++ {
		println("Index = ", index)
	}

	//  we can also use for loop as a while loop
	//  by moving the intialization outside the loop and the update inside the loop body
	usersCount := 0
	const USER_COUNT_LIMIT = 5
	for usersCount <= USER_COUNT_LIMIT {
		println("Users count = ", usersCount)
		usersCount++
	}

	for value := 0; value <= 12; value++ {
		if value%2 == 1 {
			continue
		}
		if value > 10 {
			break

		}
		fmt.Println("Value", value)
	}

	//
}
