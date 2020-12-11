package main

import (
	"github.com/gfbett/gtd/persistence"
	"github.com/gfbett/gtd/tasklist"
	"log"
	"os"
)

const storageSubfolder = "/.gtdgo"

type GTD struct {
	storageFolder string
	inbox         *tasklist.TaskList
}

func LoadGTD() *GTD {
	gtd := new(GTD)
	home, _ := os.UserHomeDir()
	gtd.storageFolder = home + storageSubfolder
	gtd.inbox = tasklist.InitTaskList()
	res := persistence.Load(gtd.inbox, gtd.storageFolder)
	if res != true {
		log.Fatal("Unable to load inbox")
		return nil
	}
	return gtd
}

func (gtd *GTD) Store() bool {
	res := persistence.CreateFolderIfNotExists(gtd.storageFolder)
	if res == false {
		log.Fatal("error creating storage folder")
		return false
	}
	res = persistence.Store(gtd.inbox, gtd.storageFolder)
	if res == false {
		log.Fatal("error storing inbox")
		return false
	}
	return true
}
