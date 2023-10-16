# Go Facts

- Go Basics

## Beginners Basic

1.  Go structure consist of componenet and each component have a meaning
2.  Go program is always declared in package and is case sensitive

         ```go

             package <package-name>

             import (
                 "fmt"
             )

             func main(){
               fmt.prinln("yo");
             }

         ```

3.  package name should not start with capital letters
4.  Statically typed, not allowed to change type at runtime
5.  exported function start with capital letter
6.

### Rough notes

- dynamically-typed languages perform type checking at runtime, while statically typed languages perform type checking at compile time
-

// Lexer checks the grammer of the given code, reports if it is or inaccordance to the language grammer
// or checks whether you have return synticaly well formed code or not

// always put a semuilcon at the end of the line
// staticaly typped langauge
// everything has a type in go
// types: String, Boolean, uint,float, complex
// more types
// Array, slices , map , structs, pointers, functions, interfaces, channels

### Documentation Comments in Go

- Go uses documentation comments for documenting code and generating package documentation.
- Package comments should be written before the package clause, starting with "Package x..."
- Function comments should be written before the function declaration, starting with the - function name, and explain arguments, purpose, and return values.

## MS certification

1.  Go differs from other programming languages in how it organizes project files. First, Go works under the concept of workspaces. A workspace is simply a location where your application source code lives. All Go projects share the same workspace. But Go has started to change this approach beginning with version 1.11. You don't have to worry about that yet because we'll cover workspaces in the next module. For now, your Go workspace is located at $HOME/go, but you can set up a different location for all your projects, if needed.

To set your workspace to a different location, you can use the $GOPATH environment variable. This environment variable helps to avoid future problems when working with more complex projects.

Configure your workspace by adding the following command to your ~/.profile:

Each Go workspace has three basic folders:

bin: Contains executables from applications.
src: Includes all application source code that lives in your workstation.
pkg: Contains compiled versions of the available libraries. The compiler can link against these libraries without recompiling them.

2. The package main statement is how we tell Go that the app we're creating is an executable program (a file you can run).
   A package is a set of common source code files. Every executable app has this first line, even if the project or file has a different name.

   - every executable program should be part of the main package.
     Challenge: Change the package name

Can you confirm the importance of the main package?

In your Go file, change the package name in the first line.
Save the file, and run your program again in the Terminal.
What happens? Do you see "Hello World!"? Is a binary executable produced?

Show answer

"Hello World!" isn't displayed. A binary executable isn't produced. You should see an error: `go run: cannot run non-main package`. Why? Every executable program should be part of the `main` package.

## Declare and use variables

- var firstName string The previous statement declares a variable named firstName of type string.
- var firstName, lastName string
- var firstName, lastName string
  var age int
  -Another way to write the previous statement is by using parenthesis after the var keyword, as if you had a dedicated block to declare variables

      ```golang

            var (
              firstName, lastName string
              age int
          )

- Initialize variables

  - you don't need to specify its type because Go will infer the type when you initialize the variable with a value

        ```golang

              var (
              firstName string = "John"
              lastName  string = "Doe"
              age       int    = 32
          )


              var (
          firstName string = "John"
          lastName  string = "Doe"
          age       int    = 32
          )

  - declare in single line

    var (
    firstName, lastName, age = "John", "Doe", 32
    )

    - when you use the colon equal sign, the variable you're declaring has to be a new one.

    firstName, lastName := "John", "Doe"
    age := 32

- Constants Go has an interesting concept related to constants that's called iota.Go has an interesting concept related to constants that's called iota. declare constants without using them and you won't get an error message. You can't use the colon equal sign to declare constants

  const (
  StatusOK = 0
  StatusConnectionReset = 1
  StatusOtherError = 2
  )

## Learn about basic data types

- Go is a strongly typed language.bound to a specific data type and will only accept values that match that type.
- In Go, we have four categories of data types:

  - Basic types: numbers, strings, and booleans
  - Aggregate types: arrays and structs
  - Reference types: pointers, slices, maps, functions, and channels
  - Interface types: interface

### Integer type

- int8, int16, int32, and int64 ,which are ints with a size of 8, 16, 32, or 64 bits,uint8, uint16, uint32, and uint64.for a 32 bit computer by default int is a 32bit integer ,and for 64bit int is 64bit.constant 9223372036854775807 overflows int

      var integer8 int8 = 127
      var integer16 int16 = 32767
      var integer32 int32 = 2147483647
      var integer64 int64 = 9223372036854775807
      fmt.Println(integer8, integer16, integer32, integer64)

- Go int of d/f sizes are not same and will require type casting if to be operated together else

       invalid operation: integer16 + integer32 (mismatched types int16 and int32)

- Unicode is a superset of ASCII, and assigns each character a standard number called “Unicode Code Point”, or in Go language, a “Rune”.A rune is simply an alias for int32 data type. It's used to represent a Unicode character (or a Unicode code point).

  - String is a sequence of bytes and not of a Rune. A string may contain Unicode text encoded in UTF-8. But, the Go source code encodes as UTF-8, therefore, no need to encode the string in UTF-8.
  - UTF-8 encodes all the Unicode in the range of 1 to 4 bytes, where 1 byte is used for ASCII and the rest for the Rune.
  - ASCII contains a total of 256 elements and out of which, 128 are characters and 0-127 are identified as code points.
  - number 71 is represented, which represents the Unicode character for G

          rune := 'G'
          fmt.Println(rune)//71

### Floating-point numbers

- GO has float 32 and float 64, can store much larger number that int64 : https://stackoverflow.com/questions/20633163/how-can-a-4-byte-store-such-large-value-of-3-4e38/20633632#20633632

            fmt.Println(math.MaxFloat32, math.MaxFloat64)
            3.4028234663852886e+38 1.7976931348623157e+308
            const e = 2.71828
            const Avogadro = 6.02214129e23
            const Planck = 6.62606957e-34
            Note  Go infers the data types from the values used.

### Bool type

- In go we can convert true/false to 0/1 implicitly we have to do typecasting.
-

### Strings

-the keyword string is used to represent a string data type. "" a string , '' characheter.

      fullName := "John Doe \t(alias \"Foo\")\n"
      fmt.Println(fullName)

### default values

- 0 for int types (and all of its subtypes, like int64)
- +0.000000e+000 for float32 and float64 types
- false for bool types
- An empty value for string types

      var defaultInt int
      var defaultFloat32 float32
      var defaultFloat64 float64
      var defaultBool bool
      var defaultString string
      fmt.Println(defaultInt, defaultFloat32, defaultFloat64, defaultBool, defaultString)

### Type Conversions

- no impilcit casting, go provides native function. or we can use https://pkg.go.dev/strconv
- \_ means we're not going to use that variable's value, and ignore it

## funtions

### main function

- All executable program need to have main funciton.Go package dont require main function since they are meant to be used by another packages.
- we can access command line arguments in Go using https://pkg.go.dev/os os package.
- os.Args variable, which holds all the arguments passed to the program.

      package main

      import (
          "fmt"
          "os"
          "strconv"
      )

      func main() {
          number1, _ := strconv.Atoi(os.Args[1])
          number2, _ := strconv.Atoi(os.Args[2])
          fmt.Println("Sum:", number1+number2)
      }

      go run main.go 3 5
      Sum: 8

- syntax for creacting a pakage,func keyword to define a function and then assign a name to it

        func name(parameters) (results) {
          body-content
       }

       func sum(number1 string, number2 string) (result int) {

- we can return more than one value

  func calc(number1 string, number2 string) (sum int, mul int, sub int) {

  int1, _ := strconv.Atoi(number1)
  int2, _ := strconv.Atoi(number2)
  sum = int1 + int2
  mul = int1 \* int2
  sub = int1 - int2
  return
  }

- The \_ variable is the idiomatic way for Go to ignore return values

#### Function Parameters

- Pass by value

       package main

       import "fmt"

        func main() {
          firstName := "John"
          updateName(firstName)
          fmt.Println(firstName)
        }

        func updateName(name string) {
         name = "David"
        }

- pass by address. A pointer is a variable that contains the memory address of another variable.
  When you send a pointer to a function, you're not passing a value, you're passing a memory address. So every change you make to that variable affects the caller.

      The & operator takes the address of the object that follows it.
      The * operator dereferences a pointer. It gives you access to the object at the address contained in the pointer.

- example

      package main

      import "fmt"

      func main() {
        firstName := "John"
        updateName(&firstName)
        fmt.Println(firstName)
      }

      func updateName(name *string) {// it's a pointer to a string.)
      *name = "David"
      }

- for storing \* on right , for using star on left

### packages

- Packages in Go are like libraries or modules in other programming languages. You can package your code and reuse it somewhere else. The source code of a package can be distributed in more than one .go file.

- If a program is part of the main package, Go generates a binary file. When that file runs, it calls the main() function. when you use the main package, your program will produce a standalone executable.

- when a program is part of a package other than main, Go doesn't generate a binary file. It generates a package archive file (a file that has the .a extension).

- A package uses the last part of its import path as its name

      import "math/cmplx"
      cmplx.Inf()

- If you want something to be private, start its name with a lowercase letter.
- If you want something to be public, start its name with an uppercase letter.

### module

- Go modules typically contain packages that offer related functionality.
- A package's module also specifies the context that Go needs to run the code you've grouped together. This contextual information includes the Go version your code is written for.

### facts

- What happens when you declare a variable but don't use it?
  You get an error. The program won't compile.
  Correct. Go requires that variables are used in programs. It will throw an error stating that the variable was declared but not used.
  A standalone application uses the package main. A package (library) uses a different name.

### Use control flows in Go

### if

- we don't need parentheses in conditions. else clause is optional. But braces are still required.
- Go doesn't offer support for ternary if statements to reduce line
- go supports compound if statments , means we can declare variables in condition block of if and can only be used with in if else and else if block , its idiomatic way.

### switch

- to avoid chaining of if statments.also does not need paranthesis in switch conditon
- if none matches in switch cases, the block simply ends, we can also have default case .
- we can also match multiple expression in te case statment by using comma, against the switch variable
  example:

          ```golang
            func location(city string) (string, string) {
              var region string
              var continent string
              switch city {
              case "Delhi", "Hyderabad", "Mumbai", "Chennai", "Kochi":
                  region, continent = "India", "Asia"
              case "Lafayette", "Louisville", "Boulder":
                  region, continent = "Colorado", "USA"
              case "Irvine", "Los Angeles", "San Diego":
                  region, continent = "California", "USA"
              default:
                  region, continent = "Unknown", "Unknown"
              }
              return region, continent
          }
        ```

- A switch condtion section can also invoke a function.we can write case statements for possible return values.
- we can also omit the condition block, and simply rely on case statments.we can execute funtion within case statment conditon, executed if true

      switch {
      case email.MatchString(contact):
          fmt.Println(contact, "is an email")
      case phone.MatchString(contact):
          fmt.Println(contact, "is a phone number")
      default:
          fmt.Println(contact, "is not recognized")
      }

- we can omit a condition in a switch statement like you do in an if statement. This pattern is like comparing a true value as if you were forcing the switch statement to run all the time.
- Generally in switch only one case stament is executed via the match or the true , when function in case, we generally beak out of it once one of the case is executed , if we want thatTo make the logic fall through to the next immediate case, use the fallthrough keyword.
  he logic goes to the next case statement immediately without validating the case condtion.
- variable assignment can also happen in switch condition

### Looop

-for loop expressions don't require parentheses. But braces are required.Go has no while keyword.

- prestatements and poststatements optional.To write a proper infinite loop, follow the for keyword with braces,
- use the continue keyword to skip the current iteration of a loo

### defer, panic, and recover

asd

## Use data types and structs, arrays, slices, and maps in Go

### Use arrays

- Arrays in Go are a fixed-length data structure of a particular type. zero or more elements. define the size and type when you declare or initialize them.
- can't resize once created, lays foundation for map and slices
- . By default, Go initializes each element with the default data type.
- length(<array>)
- examples

         ```golang
             var a [3]int
             a[1] = 10
             b:=a[len(a)-1]
             cities := [5]string{"New York", "Paris", "Berlin", "Madrid"}
             q := [...]int{1, 2, 3}
             numbers := [...]int{99: -1}

         ```

  we don't need to assign a value to all elements. As we've seen before, the last position has an empty string because it's the default value for a string data type

- Another way of declaring and initializing an array when you don't know how many positions you'll needs
  - Another interesting way of initializing an array is to use an ellipsis and only specify a value for the last position.numbers[99]==-1 true, length is 100

### Multidimensional arrays

    ```golang

            var twoD [3][5]int
          for i := 0; i < 3; i++ {
              for j := 0; j < 5; j++ {
                  twoD[i][j] = (i + 1) * (j + 1)
              }
              fmt.Println("Row", i, twoD[i])
          }

          var threeD [3][5][2]int
          for i := 0; i < 3; i++ {
              for j := 0; j < 5; j++ {
                  for k := 0; k < 2; k++ {
                      threeD[i][j][k] = (i + 1) * (j + 1) * (k + 1)
                  }
              }
          }

    ```

### slices

- slice is a data type in Go that represents a sequence of elements of the same type.And size of slice is dynamic.not fixed.
- A slice is a data structure on top of an array or another slice
- With a slice, you can access the whole underlying array or only a subsequence of elements.
- A slice has only three components:

Pointer to the first reachable element of the underlying array. This element isn't necessarily the array's first element, array[0].
Length of the slice. The number of elements in the slice.
Capacity of the slice. The number of elements between the start of the slice and the end of the underlying array.
