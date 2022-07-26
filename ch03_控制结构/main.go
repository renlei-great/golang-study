package main

import "fmt"

func main() {
	switch j := 1; j {

	case 1:
		fmt.Println("1")
		fallthrough

	case 2:

		fmt.Println("2")

	case 3:

		fmt.Println("3")

	default:

		fmt.Println("没有匹配")

	}

}
