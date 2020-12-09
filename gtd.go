package main

import (
	"github.com/gfbett/gtd/persistence"
	"github.com/gfbett/gtd/tasklist"
	"log"
	"os"
)

const storageSubfolder = "/.gtdgo"

type GTD struct {
	inbox *tasklist.TaskList
}

func InitGTD() *GTD {
	gtd := new(GTD)
	gtd.inbox = tasklist.InitTaskList()
	return gtd
}

func (gtd *GTD) Store() bool {
	home, _ := os.UserHomeDir()
	storageFolder := home + storageSubfolder
	res := persistence.CreateFolderIfNotExists(storageFolder)
	if res == false {
		log.Fatal("error creating storage folder")
		return false
	}
	res = persistence.Store(gtd.inbox, storageFolder)
	if res == false {
		log.Fatal("error storing inbox")
		return false
	}
	return true
}
