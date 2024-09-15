package main

import "fmt"

func main() {

	var marks [5]int
	fmt.Println("Enter the marks for 5 subj")
	for i := 0; i < 5; i++ {
		fmt.Scan(&marks[i])
	}
	sum := 0
	for i := 0; i < 5; i++ {
		sum += marks[i]
	}
	fmt.Printf("Total marks is %d", sum)

}
