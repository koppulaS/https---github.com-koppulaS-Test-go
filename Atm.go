package main

import (
	"fmt"
)

type AccountDetails struct {
	UserId   int32
	Password int32
	Balance  float64
}

func WelcomeScreen() {
	a := AccountDetails{
		UserId:   3456,
		Password: 4556,
		Balance:  679.09,
	}

	var choice int
	fmt.Println("Hi , welcome")
	fmt.Println("please select an option from below")
	fmt.Println("1. login")
	fmt.Println("2. create new account")
	fmt.Println("3. quit")
	fmt.Scan(&choice)
	var userid int32
	var password int32
	switch choice {
	case 1:

		fmt.Println("please enter userid  : ")
		fmt.Scan(&userid)
		fmt.Println("please enter password :")
		fmt.Scan(&password)

		if userid == a.UserId && password == a.Password {

			fmt.Println("\nEnter any option to be served!\n\n")
			fmt.Println("1. Withdraw Money\n")
			fmt.Println("2. Deposit Money\n")
			fmt.Println("3. Request Balance\n\n")
			fmt.Println("4. Quit\n")
		} else {
			fmt.Println("userid & Password did not match, please try again")
		}
		//	WelcomeScreen()

	}
}
func main() {
	WelcomeScreen()
}
