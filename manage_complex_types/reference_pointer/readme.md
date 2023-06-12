# Reference Values with Pointers in Go
In Go, you can use pointers, which are variables that store the memory address of another variable. Pointers in Go are similar to pointers in C and other similar languages. Let's explore how pointers work in Go.

To declare a pointer, you use the VAR keyword followed by the variable name and the type, indicated by an asterisk *. If you don't assign anything to the pointer variable, it will be nil.

```go
 
var P *int // Declaring a pointer variable
 ``` 

When outputting the value of a pointer, if you just use the pointer variable name, it will show as nil. However, if you use the asterisk before the variable name (*P), indicating that it is a pointer and pointing to another variable, it will crash the application if the pointer is not pointing to a valid memory address.

To point a pointer variable at a valid variable, you use the ampersand & operator followed by the variable name.

 ```go
 
var anInt = 42
var P = &anInt // Pointing the pointer at the memory address of anInt
 ``` 

Now, when you output the value of P, it will display the value of the variable it is pointing to.

You can also use pointers with other types, such as floating-point values. Here's an example:

 ```go
 
var value1 = 42.13
var pointer1 = &value1 // Pointing the pointer at the memory address of value1
 ``` 

To output the value using the pointer, you use the asterisk before the pointer variable name (*pointer1).

When changing a value through a pointer, it also affects the original variable. This behavior is similar to reference variables in languages like Java and C#. Here's an example:

 ```go
 
*pointer1 = *pointer1 / 31 // Changing the value through the pointer
 ``` 
 
If you output the value using both the pointer and the original variable, you'll see that the value has been changed.

In Go, pointers are flexible and can be changed at runtime to point to different variables. Unlike in Java, pointers in Go don't have to initially point to any particular value.

If you're familiar with pointers in languages like C or C#, you'll find pointers in Go to be very similar and equally valuable.

That covers the basics of working with pointers in Go. Pointers allow you to reference and modify variables indirectly by manipulating their memory addresses.