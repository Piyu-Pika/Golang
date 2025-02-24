
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
 9. **If/Else Statements** (`main.go` lines 0-15)
    - Conditional execution based on boolean expressions
    - `if`, `else if`, and `else` blocks
    - Example of decision making in Go
    - Run with: `go run main.go`
 10. **For Loops** (`main.go` lines 0-20)
     - Iteration using `for` loops (Go's only loop construct)
     - Basic `for` loop structure
     - Simulating `while` loops
     - Iterating over arrays and slices with `range` (index and value)
     - Creating infinite loops and using `break`
     - Run with: `go run main.go`
 11. **Switch Statements** (`main.go` lines 0-15)
     - Multi-way branching based on value matching
     - `switch` statement with `case` clauses
      - Optional `default` case
     - Implicit `break` after each case
     - Run with: `go run main.go`
 12. **Maps** (`main.go` lines 0-15)
     - Declaration and initialization of maps (key-value pairs)
     - Accessing and modifying map elements
     - Checking for key existence
     - Using `len()` to get the number of key-value pairs
     - Run with: `go run main.go`
 13. **Structs** (`main.go` lines 0-15)
     - Declaration and initialization of structs (key-value pairs)
     - Accessing and modifying struct elements
     - Run with: `go run main.go`
 14. **Pointers** (`main.go` lines 0-15)
     - Declaring and initializing pointers
     - Dereferencing pointers
     - Pointer arithmetic
     - Run with: `go run main.go`
 15. **Data Conversion** (`main.go` lines 0-15)
     - Converting between different data types
     - Implicit type conversion
     - Explicit type conversion
     - Run with: `go run main.go`
 16. **Strings** (`main.go` lines 0-15)
     - Splitting strings
     - Counting words in a string
     - Trimming whitespace
     - Joining strings
     - Run with: `go run main.go`
17. **file-manage** (`main.go` lines 0-15)
    - Reading and writing files
    - Creating and deleting directories
    - Renaming files
    - Run with: `go run main.go`
18. **http** (`main.go` lines 0-15)
    - Making HTTP requests
    - Parsing JSON responses
    - Handling errors
    - Run with: `go run main.go`
19. **url-parse** (`main.go` lines 0-15)
    - Parsing URLs
    - Extracting query parameters
    - Run with: `go run main.go`
20. **apiCall** (`main.go` lines 0-15)
    - Making API calls
    - Parsing JSON responses
    - Handling errors
    - Run with: `go run main.go`
21. **json-parse** (`main.go` lines 0-15)
    - Parsing JSON
    - Accessing JSON elements
    - Run with: `go run main.go`    
22. **goroutines** (`main.go` lines 0-15)
    - Creating and managing goroutines
    - Passing data between goroutines
    - Synchronizing goroutines
    - Run with: `go run main.go`
23. **syncWait** (`main.go` lines 0-15)
    - Using `sync.WaitGroup` to synchronize goroutines
    - Using `sync.Mutex` to synchronize access to shared data
    - Run with: `go run main.go`


 

## Getting Started

1. Install Go from [go.dev/dl](https://go.dev/dl/)
2. For examples named `main.go`, create separate directories for each to avoid conflicts
3. Run individual examples using `go run filename.go`

Note: The examples use numbered line references from the original question - in practical use, each should be in its own directory with proper file organization.