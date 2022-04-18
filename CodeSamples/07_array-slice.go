package main

import "fmt"

/*

- Arrays are fixed length

- Slice can grow and shrink (flexible)
- Slice is just references to the base array
the references need to be contiguous portion of the array

*** Slices are built on top of Arrays

make(<type>, <len>, <cap>)

*/


func main() {
   // array declaration
   var x [5]int
   var y [5]float64{98, 93, 77, 82, 83}
   z := [5]float32 {
      98,
      93,
      77,
      82,
      83,
   }

   // declare a slice
   // difference between array and slice is the missing length in the brakets
   var s []float32

    // Declare a slice
    courses := make([]string, 5, 10)

    fmt.Printf("Length is: %d \nCapacity is: %d", len(courses), cap(courses));


    progLang := []string{"Docker", "Puppet", "Python"}

    fmt.Printf("\n\nLength is: %d \nCapacity is: %d", len(progLang), cap(progLang));


    mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

    index := 4
    fmt.Println("\nMy slice val:", mySlice[index])


    sliceOfSlice := mySlice[2:5]

    fmt.Println(sliceOfSlice)

}
