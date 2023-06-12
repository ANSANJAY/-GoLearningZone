package main

import(
	"fmt"
	"bufio"
	"os"
)



func main(){

reader := bufio.NewReader(os.Stdin) // new variable
fmt.Print("Enter text: ") // promt to user
input , _ := reader.ReadString('\n') //if you have to ignore vairable , name _
fmt.Println("given value:",input)
}
