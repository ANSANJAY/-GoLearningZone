package main

import(
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)


//read strings and convert to integer
// - := error object

func main(){

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the text :")
	input , _ :=reader.ReadString('\n')
	fmt.Println("Entered line :",input)

	fmt.Print("Enter a number :");
	numInput,_:=reader.ReadString('\n')
	aFloat,err := strconv.ParseFloat(strings.TrimSpace(numInput),64)
	if err != nil{
		fmt.Println(err)
	} else{
		fmt.Print(aFloat)
		}


}
