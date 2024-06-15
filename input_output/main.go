package main

import "fmt"

type calculatorChoice int

const (
	add calculatorChoice = iota + 1
	sub
	mul
	div
	logout
)

func main() {
	fmt.Println("Wecome to calculator world!")

	for {
		fmt.Print("1. Add\n2. Subtract\n3. Multiply\n4. Divison \n5. Logout\nEnter Choice:")

		var choice calculatorChoice
		fmt.Scanf("%d\n", &choice)
		if choice == logout {
			fmt.Println("Thank you, you successfully got logged out")
			break
		}

		a, b := takeInput()
		switch choice {
		case add:
			fmt.Printf("Addition of a:%v and b:%v is %v\n", a, b, a+b)
		case sub:
			fmt.Printf("Addition of a:%v and b:%v is %v\n", a, b, a-b)
		case mul:
			fmt.Printf("Addition of a:%v and b:%v is %v\n", a, b, a*b)
		case div:
			if b != 0 {
				fmt.Printf("Division of a: %v and b: %v is %v\n", a, b, a/b)
			} else {
				fmt.Println("Error: Division by zero is not allowed")
			}
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func takeInput() (int, int) {
	var (
		a, b int
	)
	fmt.Print("Enter value for a:")
	fmt.Scanf("%v", &a)
	fmt.Print("Enter value for b:")
	fmt.Scanf("%v", &b)
	return a, b
}
