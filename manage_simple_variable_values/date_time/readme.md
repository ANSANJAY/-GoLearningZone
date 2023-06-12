# Working with Dates and Times in Go
In Go, you can manage dates and times using the time package. Variables declared with the time type encapsulate all the necessary functionalities for dates, times, mathematical operations, time zone management, and more. Let's explore some examples.

First, we need to import the time package:

```go
 import "time"
 ```

To obtain the current date and time, you can use the time.Now() function:

```go 
n := time.Now()
fmt.Println("I recorded this video at", n)
```

The above code retrieves the current timestamp using time.Now() and outputs it as a string.

You can also create your own explicit date/time values using the time.Date() function. Here's an example:

```go
 t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
fmt.Println("Go launched at", t)
```

In the above code, we create a specific date/time value representing November 10, 2009, at 11:00 PM UTC using time.Date(). We then output this value.

To format the date/time value in a specific layout, you can use the Format() function:

```go
fmt.Println(t.Format(time.ANSIC))
```

The Format() function formats the t value according to the layout string provided. In this case, we use the time.ANSIC constant, which represents a standard format.

If you have a string representing a date/time value and want to parse it into a time object, you can use the time.Parse() function:

```go
 parsedTime, _ := time.Parse(time.ANSIC, "Tue Nov 10 23:00:00 2009")
fmt.Printf("The type of parsedTime is %T\n", parsedTime)
```

The `Parse() `function converts the string into a time object based on the provided layout. In this example, we use the time.ANSIC layout again. The resulting parsedTime object can then be used accordingly.

The `time` package offers many other useful functions and constants. You can perform operations such as adding dates together and utilize various formatting patterns to customize the output of date/time values. To explore more about working with dates and times in Go, refer to the documentation on the time package available on the Go language website.

That covers the basics of working with dates and times in Go.