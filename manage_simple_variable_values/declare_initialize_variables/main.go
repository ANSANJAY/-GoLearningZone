package main

import(
	"fmt"
)

//constatnts can be declared outside the function
const(
	aconst int =64
)


func main(){

	var aString string = "This is a string"
	fmt.Println(aString)
	fmt.Printf("The variabe type is %T\n",aString)

	var sInt int =22
	fmt.Println(sInt)
	fmt.Printf("This is %T\n",sInt)

	var defaultInt int
	fmt.Println(defaultInt)

	//another way :=

	bString := "This is a stringusing :="
	fmt.Println(bString)
	fmt.Printf("The variabe type is %T\n",bString)

	fmt.Println(aconst)
	fmt.Printf("The variabe type is %T\n",aconst)
}
