# Create Loops with For Statements
In Go, you can iterate through collections of values using the for statement. Unlike in C and Java, there is no while statement, but you can achieve the same types of iterations with extended versions of the for statement syntax.

The most common for statement in Go is similar to other C-style languages, with a three-part declaration. Let's start with an array of strings with three values:

 ```go
colors := []string{"red", "green", "blue"}
fmt.Println(colors)
```

To iterate through the array, we'll use the three-part for statement declaration. We need a counter variable (i), an evaluation condition (i < len(colors)), and a post-loop update (i++):

 ```go
for i := 0; i < len(colors); i++ {
    fmt.Println(colors[i])
}
```

Alternatively, you can assign the counter variable using the range keyword, which is more concise and readable:

 ```go
for i := range colors {
    fmt.Println(colors[i])
}
```

Using range sets i to the index position each time through the loop, and you don't need to explicitly reset the counter variable.

You can also use a foreach loop by returning two variables from range in a comma-delimited list. The first variable will be the index, and the second will be the associated value. If you only need the value, you can ignore the index by naming it with an underscore character:

 ```go
for _, color := range colors {
    fmt.Println(color)
}
```

In Go, there is no dedicated while keyword. Instead, you can use the for keyword with a boolean condition. Declare a variable (value) and set it to a value (e.g., 1). Then use a for loop and specify the condition (value < 10). Inside the loop, output the value, and increment it by one:

 ```go
value := 1
for value < 10 {
    fmt.Println(value)
    value++
}
```

The break and continue statements work similarly to other C-style languages. break jumps to the end of the current code block and can be used with both for and switch statements. continue starts the next iteration of the code block.

Additionally, Go supports the use of the goto statement with labels. You can add labels to your code and jump to the label using goto or break statements:

 ```go
sum := 0
for {
    if sum > 200 {
        goto theEnd
    }
    sum += sum
}
theEnd:
fmt.Println(sum)
```

In this example, the loop continues until the sum exceeds 200, at which point it jumps to the label theEnd. Note that goto statements should be used sparingly and with caution.

In summary, Go provides various coding patterns for looping, including traditional C-style syntax and additional features to make the for statement concise and readable.

Please note that when copying and pasting the Markdown text, you may need to adjust the formatting or indentation as needed.




