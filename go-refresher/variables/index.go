package variables

import (
	"fmt"
)

func TestVariables() {
	/*
		-> in golang we can declare variables in two ways
		1. using var keyword(with explicit and implicit type declaration)
		2. using short variable declaration
	*/
	// using var keyword with explicit type declaration
	var userName string = "John Doe"
	var userAge int = 32
	var isUserActive bool = true
	var userHeight float32 = 1.56

	fmt.Println("UserName = ", userName, "UserAge = ", userAge, "UserHeight = ", userHeight)
	if isUserActive {
		println("User is active")
	} else {
		println("User is not active")
	}
	//  using the shortest and form of creating varibale

	otherUserName := "Jane Doe"
	messageCount := 10
	isMessageRead := false

	messageCount += 1

	fmt.Println("OtherUserName = ", otherUserName, "MessageCount = ", messageCount)
	if isMessageRead {
		println("Message is read")
	} else {
		println("Message is not read")
	}

	//  The last way is the version of the first and we can declare variable using type inference
	var anotherUserName = "Jack Doe"
	var anotherUserAge = 28
	var isAnotherUserActive = false

	fmt.Println("AnotherUserName = ", anotherUserName, "AnotherUserAge = ", anotherUserAge)
	if isAnotherUserActive {
		println("Another user is active")
	} else {
		println("Another user is not active")
	}

}
