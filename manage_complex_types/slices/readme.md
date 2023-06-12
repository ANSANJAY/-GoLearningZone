# Manage Ordered Values in Slices

In Go, a slice is an abstraction layer that sits on top of an array. When you declare a slice, the runtime allocates the required memory and creates the underlying array in the background, but it returns the slice. Slices are similar to arrays in that all items in a slice are of the same type, but they are resizable and can be easily sorted.

Let's explore slices using an example. Assume we have a variable called colors declared as an array with three items: red, green, and blue. If we run the code, the output will be as expected: red, green, blue. However, since we have explicitly set the number of items, this is an array, not a slice.

To convert it to a slice, we can remove the number of items and allow it to be resizable. Here's an example:

 ```go
var colors []string // Declaring a slice instead of an array
```

Now we can add and remove items from the slice. To add an item, we use the built-in append function:

 ```go
colors = append(colors, "purple") // Adding an item to the slice
```
To remove items from a slice, we also use the append function, but with a specified range. Here's an example that removes the first and last items:

 ```go
colors = append(colors[:1], colors[2:]...) // Removing the first and last items
```
You can also declare a slice with a type and an initial size using the make function. The make function takes three arguments: the type of the slice's items, the initial length, and an optional capacity. Here's an example:

 ```go
numbers := make([]int, 5, 5) // Declaring a slice of integers with an initial length and capacity of 5
numbers[0] = 134
```
You can add items to the slice using the append function, similar to what we did earlier with the colors slice.

Slices can also be sorted using the sort package. To sort a slice of integers, you can import the sort package and use the sort.Ints function. Here's an example:

 ```go
import "sort"

// Sorting the slice of integers
sort.Ints(numbers)
fmt.Println(numbers) // Output: [134 x x x x] (sorted in numerical order)
```


Keep in mind that the sort package provides various sorting techniques, including sorting slices of structures by member values. You can explore more sorting techniques in the documentation for the sort package.

That covers the basics of managing ordered values in slices in Go.