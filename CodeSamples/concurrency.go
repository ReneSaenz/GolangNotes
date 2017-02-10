package main


import(
"fmt"
"time"
"sync"
)


/*

Concurency is the COMPOSITION of independently executing processes
Parallelism is the SIMULTANEOUS EXECUTION of (possibly related)  computations

Concurrency is about dealing  with lots of things at once
Parallelism is about doing lots of things at once


+ goroutines vs OS Threads

- Lighter weight
- Go manages goroutines (simple for programmers)
- Less switching
- Faster startup times
- Safer communication

Go concurrency model: Actor model
CSP (Communicating Sequential Processes)

*/


func main() {

   var waitGrp sync.WaitGroup
   waitGrp.Add(2)

   // Self-executing, annonymous functions

   go func(){

   	time.Sleep(3 * time.Second)
   	fmt.Println("Hello")

   	defer waitGrp.Done();

   }()


   go func(){

   	fmt.Println("Pluralsight")

   	defer waitGrp.Done();

   }()


   waitGrp.Wait();
}