package main

import (
	"fmt"
)

func main(){
something()
multisum,multicount := addval(5,7,8,9)
fmt.Println("sum and count is ",multisum,multicount)
}

func something(){
	fmt.Println("doing something")
}

func addval(val ...int)(int,int){
total := 0
for _,v:=range val{
	total += v 
}
return total,len(val)
}