package main

import(
	"fmt"
)

func main(){
	aInt := 42
	var p = &aInt
	fmt.Println("value of p",*p )

	val1 := 44.3
	p1 := &val1
	fmt.Println("value of p1",*p1 )

	*p1 = *p1/2
	fmt.Println("value of p1",*p1 )
	fmt.Println("value of p1",val1 )




}