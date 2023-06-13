package main

import(
	"fmt"
)

func main(){
	a := Flowers{"orang","marigold","sweet"}
	fmt.Println(a)
	fmt.Printf("%+v\n",a)
	a.Color = "pink"
	a.Name = "lily"
	fmt.Println(a)

	a.Frag()

	}

type Flowers struct{
	Color string
	Name string
	Smell string
}

//Frag is a function
func (f Flowers) Frag(){
	fmt.Println(f.Smell)
}
