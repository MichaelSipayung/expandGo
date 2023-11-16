package main //package name
import (
	"fmt" //import fmt package
	"os"
	"strings"
)

//import os package
func main() {
	loop_n_times(10);
	//print command line arguments
	command_line();
	command_line_second();
	command_line_third();
}
//echo version 1. print it's command line arguments
func command_line(){
	println("command line arguments v:1.0")
	var s, sep string //declare variables
	for i:=1; i<len(os.Args); i++{ //loop through arguments
		s+=sep +os.Args[i] //concatenate string
		sep=" " //add space between each argument
	}
	fmt.Println(s) //print string
}
//pass n as argument to this function 
func loop_n_times(n int){
	for i := 0; i < n; i++ {
		fmt.Printf("hello, world :%d \n", i)
	}
}
//echo version 2. print it's command line arguments
func command_line_second(){
	println("command line arguments v:2.0")
	var s,sep string;
	//using blank identifier
	for _, arg:=range os.Args[1:]{ //loop through arguments, ignore first argument
		s+=sep+arg;
		sep = " "
	}
	fmt.Println(s)
}
//echo version 3. print it's command line arguments
func command_line_third(){
	println("command line arguments v:3.0")
	fmt.Println(strings.Join(os.Args[1:]," "))
}