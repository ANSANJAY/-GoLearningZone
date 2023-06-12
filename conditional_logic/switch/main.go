package main

	import(
		"fmt"
		"math/rand"
		"time"
	)
	
	func main() {
	
		rand.Seed(time.Now().Unix())
		//dow := rand.Intn(3)+1
		//fmt.Println("Day",dow) // random no. between 1 to 7
	
		var res string
		switch dow := rand.Intn(3)+1; dow {
		case 1:
			res = "Sun"
			//fallthrough
		case 2:
			res = "Mon"
			//fallthrough
		case 3:
			res = "Tues"
			//fallthrough
		default:
			res = "othr day"
		}
		fmt.Println(res)
	
	}

