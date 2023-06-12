package main

import(
	"fmt"
	"sort"
)


func main(){
	states := make(map[string]string)
	fmt.Println(states)
	states["A"]="Ana" 
	states["B"]="Bana"
	states["C"]= "Cana"

	fmt.Println(states)
	Cana := states["C"]
	fmt.Println(Cana)
	
	//delete(states,"C")
	//fmt.Println(states)

	for k , v := range states{
		fmt.Printf("%v:%v\n",k ,v)
	}

	keys := make([]string,len(states))
	i :=0
	for k:= range states{
		keys[i]=k
		i++
	}

	fmt.Println(keys)
	sort.Strings(keys)
	fmt.Println(keys)


	}

