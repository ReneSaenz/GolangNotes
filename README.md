# Go Language Notes

The ___go___ programming language was created by Google in 2007.
It is a [compiled](https://en.wikipedia.org/wiki/Compiler), [statically typed language](https://en.wikipedia.org/wiki/Type_system#STATIC).
It has pointers and
[garbage collection](https://en.wikipedia.org/wiki/Garbage_collection_(computer_science).
Limited [structural typing](https://en.wikipedia.org/wiki/Structural_type_system), [memory safety](https://en.wikipedia.org/wiki/Memory_safety) features and concurrent programming.

## Getting workspace ready

All source code for go is organized in a single directory. This directory is known as the workspace.

The workspace is a directory hierarchy with three directories at its root:
- ___src___ contains Go source files
- ___pkg___ contains package objects
- ___bin___ contains executable commands

A workspace contains many version control repositories. Each one contains one or more packages.
The path to a package's directory determines its ___import path___

The ___go___ tool builds packages from code and installs the resulting binaries to the ___pkg___ and ___bin___ directories.

The ___src___ subdirectory typically contains multiple version control repositories that track the development of one or more source packages.

## The ___GOPATH___ environment variable

The ___GOPATH___ environment variable specifies the location of your workspace.
To get started, create a workspace directory and set ___GOPATH___ to point to that directory. The workspace directory can be located wherever you like.
For example:

```$ mkdir $HOME/go_workspace```

```$ export GOPATH=$HOME/go_workspace```

## Resources

For more information, visit [___GOPATH___](https://golang.org/cmd/go/#hdr-GOPATH_environment_variable)

For more information, visit
[Go programming language](https://golang.org/doc/effective_go.html)
