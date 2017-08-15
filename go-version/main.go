package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// Suffix means it is the text file
const Suffix = ".txt"

func main() {
	path := flag.String("path", "./draft", "draft path default current")
	_, err := ioutil.ReadDir(*path)
	if err != nil {
		fmt.Println("create the dir: " + *path)
		err = os.Mkdir(*path, 0777)
	}
	err = createCurrentDraft(*path)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func createCurrentDraft(path string) error {
	currentPath := path + "/current"
	_, err := os.Stat(currentPath + Suffix)
	if err == nil {
		now := time.Now().Format("2006-01-02")
		err = renameFile(currentPath, path+"/"+now)
		if err != nil {
			return err
		}
	}
	fmt.Println("create current.txt")
	os.Create(currentPath + Suffix)
	return nil
}

func renameFile(oldPath, newPath string) error {
	_, err := os.Stat(newPath + Suffix)
	if err == nil {
		err = renameFile(newPath, newPath+"_1")
		if err != nil {
			return err
		}
	}
	err = os.Rename(oldPath+Suffix, newPath+Suffix)
	if err != nil {
		return err
	}
	return nil
}
