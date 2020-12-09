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

	gtd := LoadGTD()

	if *inbox != "" {
		fmt.Println("Adding task", *inbox)
		gtd.inbox.AddTask(tasklist.NewTask(*inbox))
	} else {
		fmt.Println("Current tasks:")
		for i := 0; i < gtd.inbox.Size(); i++ {
			fmt.Println(fmt.Sprint(i) + " - " + gtd.inbox.GetTask(i).Name)

		}

	}

	gtd.Store()
}
