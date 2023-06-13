# Define and Call Functions
Go is organized with packages, and packages have functions. Your own application has a package always named main, and it has a function also named main that's automatically called by the runtime when the application starts up. You can create your own custom functions and organize them into your own custom packages.

Let's start with a fresh copy of main.go in the practice directory and create a new custom function called forLoop. The function name is followed by parentheses, and if you're not returning a value, you don't need anything between the parentheses and the braces. Within the function, you can output something to the console by calling Println and passing in the desired message.

 ```go
func forLoop() {
    fmt.Println("Doing something")
}
```

To call the function, simply write its name:

 ```go
forLoop()
```
You can also pass arguments into a function by declaring each argument with a name and a type. If a function is going to return a value, declare the return type after the function's closing parentheses.

 ```go
func addValues(value1 int, value2 int) int {
    return value1 + value2
}

func main() {
    sum := addValues(5, 8)
    fmt.Println(sum)
}
```


If a function is going to receive multiple parameters of the same type, you don't need to declare the type more than once:

 ```go
func addValues(value1, value2 int) int {
    return value1 + value2
}
```


A function can also accept an arbitrary number of values of the same type. To do this, declare the parameter name, followed by three dots (...), and then the type:

 ```go
func addAllValues(values ...int) int {
    total := 0
    for _, v := range values {
        total += v
    }
    return total
}
```


To call the function with multiple values, you can pass in any number of integers:

 ```go
multiSum := addAllValues(2, 4, 6, 8)
fmt.Println(multiSum)
```


Go also allows you to return more than one value from a function. To accomplish this, change the function signature to return a comma-separated list of types wrapped in parentheses. You can then return the values using the return statement.

 ```go
func addAllValues(values ...int) (int, int) {
    total := 0
    for _, v := range values {
        total += v
    }
    return total, len(values)
}
```


When calling the function, you can assign the returned values to separate variables:

 ```go
multiSum, multiCount := addAllValues(2, 4, 6, 8)
fmt.Println(multiSum)
fmt.Println(multiCount)
```


There are other things you can do with functions, such as naming the returned values or using different syntax styles. For more variations in function syntax, refer to the Go documentation on functions.

Please note that when copying and pasting the Markdown text, you may need to adjust the formatting or indentation as needed.