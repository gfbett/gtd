package main

import (
	"fmt"
	"github.com/gfbett/gtd/tasklist"
)
import "flag"

var inbox = flag.String("i", "", "Adds a task in the inbox")

func main() {

	// inputs parameters declaration
	flag.Parse()

	gtd := InitGTD()

	if *inbox != "" {
		fmt.Println("Adding task", *inbox)
		gtd.inbox.AddTask(tasklist.NewTask(*inbox))
	}

	gtd.Store()
}
