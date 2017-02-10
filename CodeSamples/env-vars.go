package main


import(
"fmt"
"os"
)

func main(){

	displayEnv();

}


func displayEnv() {

	for _, env := range os.Environ(){
		fmt.Println(env);
	}
	
}