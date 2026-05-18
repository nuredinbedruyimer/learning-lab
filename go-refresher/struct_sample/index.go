package struct_sample

import "fmt"

type User struct {
	userName   string
	email      string
	role       int
	isLoggedIn bool
}

func fullData(u User) string {
	return u.userName + " " + u.email
}

func StructSample() {

	user1 := User{
		userName:   "John Doe",
		email:      "john.doe@example.com",
		role:       1,
		isLoggedIn: true,
	}
	user2 := User{
		userName:   "Jane Doe",
		email:      "jane.doe@example.com",
		role:       2,
		isLoggedIn: false,
	}
	users := []User{
		user1,
		user2,
		User{
			userName:   "Bob Smith",
			email:      "bob.smith@example.com",
			role:       3,
			isLoggedIn: false,
		},
	}

	for _, user := range users {
		fmt.Println(fullData(user))
	}

}
