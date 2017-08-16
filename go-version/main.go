package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"time"
)

// Suffix means it is the text file
const Suffix = ".txt"

// SavePath means the draft save path
var SavePath = ""

func main() {
	path := flag.String("path", "./draft", "draft path default current")
	updateTime := flag.String("time", "00:00:00", "update draft time")
	flag.Parse()
	_, err := ioutil.ReadDir(*path)
	if err != nil {
		log.Println("create the dir: " + *path)
		err = os.Mkdir(*path, 0777)
	}
	SavePath = *path
	scheduleTask(*updateTime)
}

func parseUpdateTime(updateTime string) (time.Time, error) {
	now := time.Now()
	t, err := time.ParseInLocation("2006-01-02 15:04:05", now.Format("2006-01-02")+" "+updateTime, time.Local)
	// t = t.Local()
	if err != nil {
		return time.Time{}, err
	}
	if t.Before(now) {
		t = t.Add(time.Hour * 24)
	}
	return t, nil
}

func scheduleTask(updateTime string) error {
	for {
		executeTime, err := parseUpdateTime(updateTime)
		if err != nil {
			log.Printf("parse update time error: %v\n", err)
			continue
		}
		log.Printf("%v will create the draft\n", executeTime.Format("2006-01-02 15:04:05"))
		log.Printf("sleep %v\n", time.Duration(executeTime.Unix()-time.Now().Unix())*time.Second)
		time.Sleep(time.Duration(executeTime.Unix()-time.Now().Unix()) * time.Second)
		err = createCurrentDraft(SavePath)
		if err != nil {
			log.Printf("create current draft error: %v\n", err)
			continue
		}
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
	log.Println("create current.txt")
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
