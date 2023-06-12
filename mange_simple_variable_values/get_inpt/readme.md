# Getting User Input in Go
In Go, you can interact with the user and receive input from the command line. To achieve this, you need to use the "bufio" and "os" packages. Here's how you can get user input in a Go program:

1. Import the required packages at the beginning of your code:

```go

import (
    "bufio"
    "fmt"
    "os"
)
```

2. Create a reader object to read input from the command line:

```go

reader := bufio.NewReader(os.Stdin)
```

3. Display a prompt to the user using fmt.Print:
```go
fmt.Print("Enter text: ")
```

4. Read the input provided by the user using reader.ReadString:
```go
input, _ := reader.ReadString('\n')
```

Note: In the example above, the second variable returned by reader.ReadString is an error object. For now, we're ignoring it by assigning it to the underscore character.

5. Print out the input received from the user:

```go
fmt.Println("You entered:", input)
```

Here's an example of the complete code:

```go

package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter text: ")
    input, _ := reader.ReadString('\n')

    fmt.Println("You entered:", input)
}
```

To run the program, open a terminal and navigate to the directory containing the Go file. Then use the `go run .` command to execute the program. You will see the prompt asking you to enter text. Type your input, and the program will display the entered text.

Remember to handle errors properly when working with user input. This example shows the basic process of obtaining input from a command line application in Go.