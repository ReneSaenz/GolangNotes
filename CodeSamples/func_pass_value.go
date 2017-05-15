package main


import(
"fmt"
)


/*

- Go by default, passes parameters by value

*/

func main() {
	
	name := "rene";
    course := "Docker Deep Dive";

    fmt.Println(name, "you are watching:", course);

    changeCourseVal(course);

    fmt.Println("you are now watching:", course);

    changeCourseRef(&course);

    fmt.Println("you are now watching:", course);

}

func changeCourseVal(course string) string {

    // Notice we do not use the colon-equals operator
    // We are not declaring a new variable.
    // We are just assigning a value to a variable
	course = "First look: Native docker clustering";

	fmt.Println("Trying to change 'course' to:", course);

	return course;

}

func changeCourseRef(course *string) string{
	*course = "First look: Native docker clustering";

	fmt.Println("Trying to change 'course' to:", *course);

	return *course;

}