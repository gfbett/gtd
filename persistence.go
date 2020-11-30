package main

import (
	"io/ioutil"
	"log"
	"os"
)

type Writable interface {
	toJson() string
	loadFromJson(data []byte)
	fileName() string
}

func CreateFolderIfNotExists(dir string) bool {
	_, err := os.Stat(dir)
	if err != nil && os.IsNotExist(err) {
		err = os.Mkdir(dir, os.ModeDir|os.ModePerm)
		if err == nil {
			return true
		} else {
			return false
		}
	}
	return true
}

func Store(writable Writable, path string) bool {

	filename := writable.fileName()
	data := writable.toJson()
	file, err := os.Create(path + "/" + filename)
	if err != nil {
		return false
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	_, err = file.WriteString(data)
	if err != nil {
		return false
	}
	return true
}

func Load(writable Writable, path string) bool {
	filename := writable.fileName()
	file, err := os.Open(path + "/" + filename)
	if err != nil {
		return false
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	all, err := ioutil.ReadAll(file)
	if err != nil {
		return false
	}
	writable.loadFromJson(all)
	return true
}
