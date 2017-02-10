package main


import(
"fmt"
)


func main(){

	mySlice := make([]int, 1, 4)

	fmt.Printf("Length is %d Capacity %d",
		len(mySlice), cap(mySlice))


	for i := 1; i < 17; i++ {
		mySlice = append(mySlice, i)
		fmt.Printf("\nCapacity is: %d", 
			cap(mySlice))
	}


    fmt.Println("\n\n")

    // Append a slice to a slice
    slice1 := []int{1, 2, 3, 4, 5}
    fmt.Println(slice1)

	for _, d := range slice1 {
		fmt.Println("for range loop:", d)
	}


	newSlice := []int{10, 20, 30}
	slice1 = append(slice1, newSlice...)

	fmt.Println(slice1)

}