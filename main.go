package main

import (
	"fmt"
	"os"
)
import "flag"

var inbox = flag.String("i", "", "Adds a task in the inbox")

func main() {

	// inputs parameters declaration
	flag.Parse()

	if *inbox != "" {
		fmt.Println("TaskList task", *inbox)
	}

	gtd := InitGTD()


	home, _ := os.UserHomeDir()
	gtd.Store(home)
}
