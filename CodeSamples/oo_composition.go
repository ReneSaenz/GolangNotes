package main

/*
Go does not support inheritance.
Composition is used instead
*/

type MessagePrinter struct {
	message string
}

func (mp *MessagePrinter) PrintMessage() {
	println(mp.message)
}

/*
 Composition
 Since there are no fields in this struct, Go is going to treat all the
 fields and methods of the struct MessagePrinter struct, as part of this
 new struct
*/
type EnhancedMessagePrinter struct {
	MessagePrinter
}

func main() {
	mp := MessagePrinter{"foo"}
	mp.PrintMessage()

	emp := EnhancedMessagePrinter{}
	emp.message = "bar"
	emp.PrintMessage()
}
