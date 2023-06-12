Group Related Values in Structs
In Go, a struct is a data structure that can encapsulate both data and methods. It is similar to the struct concept in computer science and classes in Java. However, Go does not have an inheritance model like Java, so each struct is independent with its own fields for data management and optional methods.

To create a custom struct type in Go, use the type keyword followed by the desired type name. In this example, let's create a struct called Dog:

 ```go
type Dog struct {
    Breed  string
    Weight int
}
 ```
Note that if the type name starts with an uppercase letter, it is exported and can be used by other parts of the application. If it starts with a lowercase letter, it is non-exported or private.

In the main function, we can create an instance of the Dog struct:

 ```go
poodle := Dog{
    Breed:  "Poodle",
    Weight: 10,
}
fmt.Printf("%+v\n", poodle)
 ```

The %+v format specifier in the Printf function is used to display the struct's field names and values. This is useful for debugging purposes.

To access the individual fields of a struct, use the dot operator. For example, to access the breed and weight of the poodle object, you can use:

 ```go
fmt.Printf("The dog's breed is %s\n", poodle.Breed)
fmt.Printf("The dog's weight is %d\n", poodle.Weight)
 ```

You can modify the values of struct fields directly. For example, to change the weight of the poodle object, you can do:

 ```go
poodle.Weight = 9
fmt.Printf("The dog's new weight is %d\n", poodle.Weight)
 ```
 
That covers the basics of using structs in Go to group and store related values. Structs can also encapsulate functionality in the form of methods, which will be covered in the next chapter.

Please note that when copying and pasting the Markdown text, you may need to adjust the formatting or indentation as needed.