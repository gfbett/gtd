package main

import (
	"log"
)

type GTD struct {
	inbox *TaskList
}

func InitGTD() *GTD {
	gtd := new(GTD)
	gtd.inbox = InitTaskList()
	return gtd
}

func (gtd *GTD) Store(homeFolder string) bool {
	storageFolder := homeFolder + "/.gtdgo"
	res := CreateFolderIfNotExists(storageFolder)
	if res == false {
		log.Fatal("error creating storage folder")
		return false
	}
	res = Store(gtd.inbox, storageFolder)
	if res == false {
		log.Fatal("error storing inbox")
		return false
	}

	return true
}
