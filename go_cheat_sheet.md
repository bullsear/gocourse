# Go (Golang) Cheat Sheet

## 1. Basic Syntax
### Variables
```go
var x int = 10        // Explicit declaration
y := 20               // Short variable declaration
const PI = 3.14       // Constant
var s string = "Hello" // String
```

### Functions
```go
func add(a, b int) int { // Function with parameters and return type
    return a + b
}

func main() {           // Entry point
    fmt.Println(add(2, 3))
}
```

## 2. Data Types
```go
// Basic types
var i int = 42
var f float64 = 3.14
var b bool = true
var s string = "text"

// Complex types
var a [3]int = [3]int{1, 2, 3} // Array
var slice []int = []int{1, 2, 3} // Slice
var m map[string]int = map[string]int{"one": 1} // Map
```

## 3. Control Structures
### If-Else
```go
if x > 10 {
    fmt.Println("Large")
} else if x == 10 {
    fmt.Println("Equal")
} else {
    fmt.Println("Small")
}
```

### Loops
```go
// For loop (only loop type in Go)
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

// Range loop
for i, v := range slice {
    fmt.Println(i, v)
}
```

### Switch
```go
switch day := "Monday"; day {
case "Monday":
    fmt.Println("Start of week")
default:
    fmt.Println("Other day")
}
```

## 4. Slices and Arrays
```go
// Array (fixed length)
var arr [3]int = [3]int{1, 2, 3}

// Slice (dynamic length)
slice := []int{1, 2, 3}
slice = append(slice, 4) // Add element
slice[1] = 5             // Modify element
```

## 5. Maps
```go
m := make(map[string]int)
m["key"] = 42        // Add/Update
value := m["key"]    // Access
delete(m, "key")     // Delete
val, ok := m["key"]  // Check existence
```

## 6. Structs
```go
type Person struct {
    Name string
    Age  int
}

p := Person{Name: "Alice", Age: 30}
p.Age = 31           // Modify field
fmt.Println(p.Name)  // Access field
```

## 7. Methods
```go
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 { // Method with receiver
    return math.Pi * c.Radius * c.Radius
}

c := Circle{Radius: 5}
fmt.Println(c.Area())
```

## 8. Interfaces
```go
type Shape interface {
    Area() float64
}

func Describe(s Shape) {
    fmt.Println("Area:", s.Area())
}
```

## 9. Goroutines and Channels
```go
// Goroutine
go func() {
    fmt.Println("Running concurrently")
}()

// Channel
ch := make(chan int)
go func() {
    ch <- 42 // Send
}()
value := <-ch // Receive
```

## 10. Error Handling
```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

result, err := divide(10, 0)
if err != nil {
    fmt.Println("Error:", err)
}
```

## 11. Packages and Imports
```go
package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Println(math.Sqrt(16))
}
```

## 12. Defer
```go
func main() {
    defer fmt.Println("Done") // Executes at function exit
    fmt.Println("Working")
}
```

## 13. Pointers
```go
x := 10
var p *int = &x // Pointer to x
*p = 20         // Modify through pointer
fmt.Println(x)  // Outputs 20
```

## 14. Type Assertions
```go
var i interface{} = "hello"
s, ok := i.(string) // Type assertion
if ok {
    fmt.Println(s)
}
```

## 15. Concurrency Patterns
### WaitGroup
```go
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    // Work
}()
wg.Wait()
```

This cheat sheet covers the core features of Go with concise examples. For more details, refer to the official Go documentation at https://golang.org/doc/.

</xaiArtifact>