package main

import(
	"fmt"
)

func main(){
	a := Flowers{"orang","marigold"}
	fmt.Println(a)
	fmt.Printf("%+v\n",a)
	a.Color = "pink"
	a.Name = "lily"
	fmt.Println(a)

	}

type Flowers struct{
	Color string
	Name string
}
