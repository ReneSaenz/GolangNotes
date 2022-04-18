package main

/*
In Go, data from the structs and methods are sepatated
New methods can be added to the struct from anywhere within the package
*/

// Struct
type MessagePrinter struct {
	message string
}

// Struct methods
func (mp *MessagePrinter) PrintMessage() {
	println(mp.message)
}

// ===========================
func newMessagePrinter() *MessagePrinter {
	//create the struct
	mp := MessagePrinter{}

	mp.message = ""
	// return the address of the new struct, as indicated by the function signature
	return &mp
}

func main() {
	mp := newMessagePrinter()
	mp.message = "This is an emergency broadcast system"
	mp.PrintMessage()
}
