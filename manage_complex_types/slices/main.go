package main

import(
	"fmt"
	"sort"
)

func main(){
	var colors = []string{"Red","green","blue"}
	fmt.Println(colors)
	colors = append(colors,"Purple")
	fmt.Println(colors)

	colors = append(colors[1:len(colors)])
	fmt.Println(colors)
	colors = append(colors[0:len(colors)-1])
	fmt.Println(colors)

	//make

	number := make([]int,5,5) // size and cap
	number[0] = 234
	number[1] = 224
	number[2] = 2324
	number[3] = 2342
	number[4] = 2343
	fmt.Println(number)
	sort.Ints(number)


}