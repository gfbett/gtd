package main

import "fmt"
import "flag"

var inbox = flag.String("i", "", "Adds a task in the inbox")

func main() {

	// inputs parameters declaration
	flag.Parse()

	if *inbox != "" {
		fmt.Println("TaskList task", *inbox)
	}
	fmt.Println("GTD GO")

}
