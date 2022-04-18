package main

import "fmt"
import "errors"


type argError struct{
	arg int
	reason string
}

func f1(arg int)(int, error){
	if arg == 42{
		//errors.New() construct a basic error value with the given error message
		return -1, errors.New("Cannot work with 42")
	}

	// A nil value in the error position indicates that there was no error
	return arg + 3, nil
}


func f2(arg int)(int, error){
	if arg == 42{
		// in this case we use &argError syntax to build a
		// new struct, supplying values for the two fields arg and prob
		return -1, &argError{arg, "Cannot work with it"}
	}

	return arg + 3, nil
}



/*
its possible to use custom types as errors by implementing the Error() method on them.
Here is a variant on the example above that uses a custom type to explicitly represent an argument error
 */

func (e *argError) Error() string{
	return fmt.Sprintf("%d - %s", e.arg, e.reason)
}


/*
The two loops below test out each of our error-returning functions.
Note that the use of an inline error check on the if like is a
common idiom in the Go code
 */
func main(){

	for _, i := range []int{7, 42}{
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else{
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{7, 42}{
		if r, e := f2(i); e != nil{
			fmt.Println("f2 failed:", e)
		} else{
			fmt.Println("f2 worked:", r)
		}
	}

	/*
	if you want to programmatically use the data in a custom error
	you'll need to get the error as an instance of the custom
	error type via type assertion
	 */
	_, e := f2(42)
	if ae, ok := e.(*argError); ok{
		fmt.Println(ae.arg)
		fmt.Println(ae.reason)
	}
}

