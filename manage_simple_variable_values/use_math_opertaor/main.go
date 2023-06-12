package main

import(
	"fmt"
	"math"
)

func main(){

	i1, i2,i3 := 12,13,15 // implicit initialization
	intSum := i1+i2+i3
	fmt.Println("Integer sum : ",intSum)

	f1,f2,f3 := 34.6,77.32,88.32
	floatSum := f1+f2+f3
	fmt.Println("Integer sum : ",floatSum)

	// for precision

	floatSum = math.Round(floatSum*100) / 100
		fmt.Println(floatSum)


	circelRadius := 13.3
	circumference := circelRadius*2*math.Pi
	fmt.Printf("Circumference : %.2f\n",circumference)
}

