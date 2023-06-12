package main

import(
	"fmt"

)

func main() {
	colors := []string{"r","b","g"}
	fmt.Println(colors)

	for i:=0;i<len(colors);i++{
		fmt.Println(colors[i])
	}

	//using range
	for i:= range colors{
		fmt.Println(colors[i])
	}

	for _,color := range colors{
		fmt.Println(color)
	}

	value := 1
	for value <10{
		fmt.Println("Value:",value)
		value++
	}

	sum :=1
	for sum <100{
		sum +=sum
		fmt.Println(sum)
		if sum > 50 {
			goto theEnd
		}
	}
theEnd:
 fmt.Println("End")
}

