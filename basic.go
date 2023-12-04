package main //package name
import (
	"bufio" //import bufio package
	"fmt"   //import fmt package
	"io/ioutil" //import io/ioutil package
	"os"
	"strings"
	"net/http" //import http package, for fetch( 
)

//import os package
func main() {
	loop_n_times(10);
	//print command line arguments
	command_line();
	command_line_second();
	command_line_third();
	//print duplicate lines
	//dup1();
	//fetch url
	fetch();
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
//dup1 prints the text of each line that appears more than
//once in the standard input, preceded by its count.
func dup1(){
	counts := make(map[string]int) //declare map of string and int
	input := bufio.NewScanner(os.Stdin) //declare input scanner
	for input.Scan(){ //loop through input
		counts[input.Text()]++ //increment count
	}
	//NOTE: ignoring potential errors from input.Err()
	for line, n:=range counts{ //loop through map counts and get line and count
		if n>1{ //if count is greater than 1
			fmt.Printf("%d\t%s\n",n,line) //print count and line
		}
	}
}
//dup2 prints the count and text of lines that appear more than once
//in the input. It reads from stdin or from a list of named files.
func dup2(){
	counts := make(map[string]int) //declare map of string and int
	files  := os.Args[1:] //declare files
	if(len(files)==0){ //if no files
		countLines(os.Stdin,counts) //read from stdin
	}else{
		for _, arg:=range files{
			f,err := os.Open(arg) //open file, return file and error
			if(err != nil){ //if error
				fmt.Fprintf(os.Stderr,"dup2: %v\n",err) //print error
				continue
			}
			countLines(f,counts) //read from file
			f.Close() //close file
		}
		for line, n:=range counts{ //loop through map counts and get line and count
			if n>1{ //if count is greater than 1
				fmt.Printf("%d\t%s\n",n,line) //print count and line
			}
		}
	}
}
//countLines counts the number of lines in the input
//it reads from stdin or from a list of named files
func countLines(f *os.File, counts map[string]int){
	input := bufio.NewScanner(f) //declare input scanner
	for input.Scan(){ //loop through input
		counts[input.Text()]++ //increment count
	}
}
//fetch prints the content found at a URL 
func fetch(){
	for _, url := range os.Args[1:]{ //loop through arguments
		resp, err  := http.Get(url) //get url
		if err != nil{ //if error
			fmt.Fprintf(os.Stderr,"fetch: %v\n",err) //print error
			os.Exit(1) //exit
		}
		b, err := ioutil.ReadAll(resp.Body) //read response body
		resp.Body.Close() //close response body
		if err != nil{ //if error
			fmt.Fprintf(os.Stderr,"fetch: reading %s: %v\n",url,err) //print error
			os.Exit(1) //exit
		}
		fmt.Printf("%s",b) //print response body
	}
}