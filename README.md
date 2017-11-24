# Go Language Notes

The **go** programming language was created by Google in 2007.
<<<<<<< HEAD
It is a [compiled](https://en.wikipedia.org/wiki/Compiler), [statically typed language](https://en.wikipedia.org/wiki/Type_system#STATIC).
=======

It is a [compiled](https://en.wikipedia.org/wiki/Compiler), 
[statically typed language](https://en.wikipedia.org/wiki/Type_system#STATIC).

>>>>>>> 797de085abacb78d8f472561916eee2510cca51f
It has pointers and
[garbage collection](https://en.wikipedia.org/wiki/Garbage_collection_(computer_science)).

Limited [structural typing](https://en.wikipedia.org/wiki/Structural_type_system), [memory safety](https://en.wikipedia.org/wiki/Memory_safety) features and concurrent programming.

**go** is compiled, concurrent, garbage-collected, statically typed language.

**go** is efficient, scalable and productive.

Why **Go**?
- [Part 1](http://golang-basic.blogspot.com/2014/05/basic-golang-why-and-what-part-1.html)
- [Part 2](http://golang-basic.blogspot.com/2014/05/basic-golang-why-and-what-part-2.html)


### Getting workspace ready

All source code for go is organized in a single directory. This directory is known as the workspace.

The workspace is a directory hierarchy with three directories at its root:
- **src** contains Go source files
- **pkg** contains package objects
- **bin** contains executable commands

A workspace contains many version control repositories. Each one contains one or more packages.
The path to a package's directory determines its **import path**

The **go** tool builds packages from code and installs the resulting binaries to the **pkg** and **bin** directories.

The **src** subdirectory typically contains multiple version control repositories that track the development of one or more source packages.

### The **go** environment variables

<<<<<<< HEAD
When you installed **go** it creates several environment variables for you. As a user, you are responsible for definiting two environment variables.


1. The **GOPATH** environment variable specifies the location of your workspace.
To get started, create a workspace directory and set **GOPATH** to point to that directory. 
The workspace directory can be located wherever you like.
=======
The **GOPATH** environment variable specifies the location of your workspace.
To get started, create a workspace directory and set **GOPATH** to point to that directory. The workspace directory can be located wherever you like.
>>>>>>> 797de085abacb78d8f472561916eee2510cca51f
For example:
```
$ mkdir $HOME/go_workspace

$ export GOPATH=$HOME/go_workspace
```

2. The **GOROOT** environment variable specifies the location where **go** was downloaded and installed.
For example
```
$ export GOROOT=/usr/local/go1.8.3
```

### Semantics in GO
The semantics of GO statements is generally C-like.
It is a compiled, statically typed procedural language with pointers.<br>
Go makes many small changes to `C` semantics, mostly in the service of robustness.
- there is **no pointer arithmetic**
- there are **no implicit numeric conversations**
- array bounds are always checked
- there are **no type aliases (after type x int, X and int are distinct types not aliases)**
- assignment is not an expression

There are some other changes from traditional `C`
- concurrency
- garbage collection
- interface types
- reflection
- type switches


### Syntax in GO
- Syntax is critical to tooling.
Go is designed with clarity and tooling in mind and has a clear syntax.
- Difference in `C` and `Go` syntax: The declared name appears before the type and there are mode keywords.

**Go** syntax
```
var fn func([]int) int
type T struct {a, b int }
```
**C** syntax
```
int (fn)(int[]);
struct T { int a, b; }
```

- A method is just a function with a special parameter, its **receiver**, which can be passed to the function using the standard "dot" notation.
Method declaration syntax places the receiver in parenthesis before the function name.
Here is a method of type **T**
```
func (x T) Abs() float64
```

- Here is a variable (closure) with a type **T** argument. **Go has first-class functions and closures**
```
negAbs := func(x T) float64 { return -Abs(x) }
```
- Go functions can **return multiple values**.
A common case is to return the function result and and **error** value as a pair.
```
func ReadByte() (c byte, err error)
c, err := ReadByte()
if err != nil { ... }
```
- Go does not support default function arguments
  - The lack of default arguments requires **more** functions or methods to be defined, as on function cannot hold the entire interface.
  - This leads to a **clearer API** that is easier to understand.
  - Those functions all need separate names, too, which makes it clear which combinations exists.


### Packages in go

Each **go** source file starts with a **package** name.

- GO style suggests keeping package name **concise**, **clear** and **obvious**
- Every GO program is made up of packages
- Programs start running in package **main**
- The package name is the same as the last element of the import path. For example `import "math/rand"` translates to package rand
- Package names may not be unique
- Package names can be overridden in each importing source file by providing a local identifier in the import clause
- Package path is the slash-separated directory path (path/to/package) of the source package in the repository. (package paths are unique)
- GO allows importing source file using the package name to reference to items in the package. For example
```
var dec = json.NewDecoder(reader)
```

### Remote packages in GO

The allocation of the space of import paths is delegated to URLs, which makes the naming of packages decentralized and therefore scalable, in contrast to centralized registries used by other language.

The **go get tool** is used to get remote repositories.
- It recursively downloads all dependencies
- It uses go build tool to fetch from remote URL and install it locally.
```
$ go get github.com/4ad/doozer  //shell command to fetch packages
import "github.com/4ad/doozer"  //Doozer client import statement
var client doozer.Conn  //client use of packages
```



  ### Resources

  For more information, visit [**GOPATH**](https://golang.org/cmd/go/#hdr-GOPATH_environment_variable)

  For more information, visit
  [The **Go** programming language](https://golang.org/doc/effective_go.html)
