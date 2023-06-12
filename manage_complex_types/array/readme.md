# Store Ordered Values in Arrays
In Go, an array is an ordered collection of values or references. While it's better to use slices instead of arrays for representing ordered collections of values, it's important to understand arrays before diving into slices.

To declare an array in Go, you specify the number of items in the array wrapped in brackets, followed by the type of the values in the array. Here's an example:

 ```go
var colors [3]string // Declaring an array of three strings
```

You can set individual values in the array explicitly using the array index:

 ```go
colors[0] = "red"
colors[1] = "green"
colors[2] = "blue"
```
To output the values of the array, you can pass the array to the fmt.Println function:

 ```go
fmt.Println(colors) // Output: [red green blue]
```

Remember that array indices start at zero, so colors[0] represents the first item in the array.

Alternatively, you can declare and initialize an array in a single statement using braces {} and a comma-delimited list of values:

 ```go
var numbers = [5]int{1, 2, 3, 4, 5} // Declaring and initializing an array of five integers
```

To find the number of items in an array, you can use the built-in len function:

 ```go
fmt.Println("Number of colors:", len(colors))
fmt.Println("Number of numbers:", len(numbers))
```

Arrays in Go are objects, and if you pass an array to a function, a copy of the array will be made. However, arrays have limited functionality. You cannot easily sort them or add/remove items at runtime. For such operations and more flexibility, it's recommended to use slices instead of arrays.

That covers the basics of storing ordered values in arrays in Go.