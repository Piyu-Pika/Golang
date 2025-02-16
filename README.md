
# Go Learning Examples

This repository contains sample Go code demonstrating various fundamental concepts. Each example focuses on different aspects of Go programming.

## Examples

1. **Hello World & Variables** (`main.go` lines 0-35)
   - Basic program structure
   - Variable declaration and initialization
   - Multiple data types (int32, float64, rune, bool, slice)
   - Run with: `go run main.go`

2. **Formatted Printing** (`main.go` lines 0-15)
   - Short variable declarations
   - Print formatting using `fmt.Printf`
   - Different format specifiers (%d, %s, %f)
   - Run with: `go run main.go`

3. **Array Operations** (`main.go` lines 0-19)
   - Array declaration and initialization
   - Length calculation with `len()`
   - Array iteration
   - Run with: `go run main.go`

4. **Function Basics** (`main.go` lines 0-18)
   - Function declaration and invocation
   - Parameter passing
   - Return values
   - Run with: `go run main.go`

5. **Input Handling** (`main.go` lines 0-17)
   - Reading user input with `bufio`
   - Basic input validation
   - Run with: `go run main.go`

6. **Error Handling** (`main.go` lines 0-21)
   - Custom error creation
   - Multiple return values
   - Error checking pattern
   - Run with: `go run main.go`

7. **Custom Package** (`myutil.go`)
   - Contains a utility function `PrintMessage`
   - Usage example (create separate main file):
     ```go
     package main
     import "./myutil"
     func main() {
         myutil.PrintMessage("Hello from package!")
     }
     ```
 8. **Slices** (`main.go` lines 0-10)
    - Basic slice declaration and initialization
    - Length calculation with `len()`
    - make function in slices
    - capacity of slices
    - run with: `go run main.go`
 9. **if else** (`main.go` lines 0-10)
    - if else statement
    - else if statement
    - else statement
    - use to check the condition
    - run with: `go run main.go`
 10. **for loop** (`main.go` lines 0-10)
    - for loop
    - only for loop available in go
    - create an infinite loop using break
    - get index,value from the array and slice using range
    - similarly we can use range to iterate over array and slice
    - run with: `go run main.go`
 11. **switch case** (`main.go` lines 0-10)
    - switch case
    - use to check the condition
    - can use multiple case
    - run with: `go run main.go`
 12. **map** (`main.go` lines 0-10)
    - map declaration and initialization
    - length calculation with `len()`
    - run with: `go run main.go`


## Getting Started

1. Install Go from [go.dev/dl](https://go.dev/dl/)
2. For examples named `main.go`, create separate directories for each to avoid conflicts
3. Run individual examples using `go run filename.go`

Note: The examples use numbered line references from the original question - in practical use, each should be in its own directory with proper file organization.