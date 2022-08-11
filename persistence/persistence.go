package persistence

import (
	"io/ioutil"
	"log"
	"os"
)

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

	filename := writable.FileName()
	data := writable.ToStorableString()
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

func FileExists(writable Writable, path string) (bool, error) {
	filename := writable.FileName()
	_, err := os.Stat(path + "/" + filename)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func Load(writable Writable, path string) bool {
	filename := writable.FileName()
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
		log.Fatal("Unable to load file")
		return false
	}
	return writable.LoadFromStorableString(string(all))
}
