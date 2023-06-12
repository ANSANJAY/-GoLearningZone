package main

import(
	"fmt"
	"time"
)

func main(){
	n:=time.Now()
	fmt.Println("Time nonw : ",n)

fmt.Println(n.Format(time.ANSIC))


}

