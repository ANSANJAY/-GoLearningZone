# Converting String Inputs to Other Types
In Go, when you receive input from the console using the bufio package and os.Stdin, the values are always in the form of strings. However, you may need to convert these strings to other types. Let's see how you can convert a string to another type in Go.

To perform string conversion, you'll need the strconv package from the Go standard library, specifically designed for string conversion. You'll also need the strings package, which provides functions for string manipulation.

In your code, follow these steps:


1. Import the required packages at the beginning of your code:


```go
 
import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)
```

2. Prompt the user to enter a number:

```go
 
fmt.Print("Enter a number: ")
```

3. Read the input string from the user:

```go
 
numInput, _ := reader.ReadString('\n')
```

4. Convert the string to a number:

```go
 
aFloat, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)
```

Note: The strings.TrimSpace function is used to remove any leading or trailing whitespace from the input string before conversion.

5. Handle the error if the conversion fails:

```go
 
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println("Value of number:", aFloat)
}
```

Here's an example of the complete code:

```go
 
package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter a number: ")
    numInput, _ := reader.ReadString('\n')

    aFloat, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)

    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("Value of number:", aFloat)
    }
}
```

To run the program, open a terminal and navigate to the directory containing the Go file. Then use the `go run .` command to execute the program. You will see the prompt asking you to enter a number. Type a numeric value, and the program will convert and display the entered number. If the input cannot be parsed as a number, an error message will be shown.

The `strconv` package provides various functions for converting strings to different Go types. Refer to the package documentation for more details on the available conversion functions.

Remember to handle errors appropriately when working with string conversions. This example demonstrates how you can convert user input from the console into the desired Go types.