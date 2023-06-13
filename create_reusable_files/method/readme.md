# Define Functions as Methods of Custom Types
In Go, you can attach functions to your custom types, which are then referred to as methods. Unlike more strictly object-oriented languages like Java, where methods are members of classes, in Go, methods are members of types.

In this example, let's revisit the Dog struct used in a previous exercise and see how to attach a method to the struct. First, we need to add a new field to the Dog struct called sound, with a type of string. Then, we can initialize the sound field when creating a Poodle object.

 ```go
type Dog struct {
    sound string
}

type Poodle struct {
    Dog
}

func main() {
    poodle := Poodle{Dog{"Woof!"}}
}
```

To create a custom method, we define a function and add a receiver before the function name. The receiver is the type that the method is attached to. In this case, we pass in a reference to a Dog object.

 ```go
func (d *Dog) Speak() {
    fmt.Println(d.sound)
}
```

We can now reference the Dog object within the method using the receiver identifier d. In this example, we use fmt.Println to output the sound field of the Dog object.

To call the method, we use the syntax poodle.Speak() in our main function.

 ```go
func main() {
    poodle := Poodle{Dog{"Woof!"}}
    poodle.Speak()
}
```

The output will be "Woof!".

Note that an exported function (one that's public to the rest of the application, indicated by an initial uppercase character in the name) should have a comment preceding it. The comment format starts with //, followed by the name of the function, followed by the word "is," and everything after that is up to you. This comment helps with code documentation and readability.

You can also change the exported or public fields of an object and then call the method again. In the example below, we modify the sound field of poodle and call the Speak method again.

 ```go
poodle.sound = "Arf!"
poodle.Speak()
```

This time, the output will be "Arf!".

In Go, methods do not support method overrides, meaning each method must have its own unique name. However, like all functions, methods can return values. To do this, you simply declare the return type assigned to the method.

 ```go
func (d *Dog) SpeakThreeTimes() string {
    d.sound = "Bark!"
    return fmt
    .Sprintf("%v %v %v", d.sound, d.sound, d.sound)
}
```

In this example, we reassign the sound field of d and use the fmt.Sprintf function to return a formatted string with three placeholders (%v) matching the sound field. We then call the method SpeakThreeTimes in our main function.

 ```go
fmt.Println(poodle.SpeakThreeTimes())
```

The output will be "Bark! Bark! Bark!".

It's important to note that when passing the Dog object as the receiver, a copy of the object is made, not a reference. Therefore, if you modify the object within the method, it won't affect the version created in the main function. If you want to modify the original object, you can use pointers.

The ability to create custom methods for your own types makes Go behave more like a fully object-oriented language,