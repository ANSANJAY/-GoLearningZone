package main

import(
	"fmt"
)

func main() {

	theAns :=43
	var result string

	if theAns<0{
		result = "less than zero"
	}else if theAns == 0 {
		result = "zero"
	} else{
		result = "greater than zero"
	}

	fmt.Println(result)
	//declaration of x is allowed
	if x := -42;x<0{ 
		result = "less than zero"
	}else if x ==0 {
		result = "zero"
	} else{
		result = "greater than zero"
	}

	fmt.Println(result)

}

