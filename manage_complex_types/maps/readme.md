# Store Unordered Values in Maps
In the Go programming language, a map is an unordered collection of key-value pairs. It acts as a hash table, allowing you to store data and retrieve items based on their keys. The keys in a map can be of any type that is comparable, although it's common to use strings for keys and any other type for associated values.

Let's explore maps using an example. Assume we have a new variable called states initialized using the make function. We specify the key type as string and the associated value type also as string. Initially, the map is empty. Here's an example:

 ```go
states := make(map[string]string) // Initializing an empty map
fmt.Println(states) // Output: map[]
```

To add items to the map, we use the key-value syntax. Here's an example that adds three states with their abbreviations:

 ```go
states := make(map[string]string)
states["WA"] = "Washington"
states["OR"] = "Oregon"
states["CA"] = "California"
fmt.Println(states) // Output: map[WA:Washington OR:Oregon CA:California]
```

We can retrieve an item from the map using its key. Here's an example that retrieves the value for the key "CA":

 ```go
california := states["CA"]
fmt.Println(california) // Output: California
```

To remove an item from the map, we can use the delete function:

 ```go
delete(states, "OR") // Removing the item with key "OR"
fmt.Println(states) // Output: map[WA:Washington CA:California]
```

Adding new items is similar to adding initial items:

 ```go
states["NY"] = "New York" // Adding a new item with key "NY"
fmt.Println(states) // Output: map[WA:Washington CA:California NY:New York]
```

Please note that the order of the map's display is not guaranteed, as it is an unordered collection. In the above example, the order appears to be alphabetical, but it may not always be the case.

To iterate over a map, we can use a for loop with the range keyword. Here's an example that loops through all key-value pairs in the states map:

 ```go
for k, v := range states {
    fmt.Printf("Key: %s, Value: %s\n", k, v)
}
```

However, the iteration order is not guaranteed either. If you want to guarantee a specific order, you need to manage it yourself. One way to achieve a sorted order is to extract the keys from the map as a slice of strings and then sort that slice. Here's an example:

 ```go
import (
    "fmt"
    "sort"
)

keys := make([]string, 0, len(states))
for k := range states {
    keys = append(keys, k)
}

sort.Strings(keys)

for _, k := range keys {
    fmt.Printf("Key: %s, Value: %s\n", k, states[k])
}
```

In this example, we extract the keys from the states map, sort them using the sort.Strings function, and then iterate over the sorted keys to output the corresponding values in a guaranteed order.

That covers the basics of using maps in Go to store unordered data collections and access items using their keys. By combining maps with slices, you can process the data in the desired order.