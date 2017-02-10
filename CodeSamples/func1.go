package main

import(
"fmt"
"strings"
)

func main() {
	
	module := "docker course";
	author := "rene";

	fmt.Println(converter(author, module));

}


func converter(author string, module string ) (ret1 string, ret2 string) {

	author = strings.ToUpper(author);
	module = strings.Title(module);

	return author, module;
	
}