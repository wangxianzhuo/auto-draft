package schedule

import (
	"fmt"
	"log"
	"time"
)

type Schedule struct {
	Time     string
	ExecTime time.Time
	Task     Task
}

type Task interface {
	Start()
}

func (s *Schedule) Run() {
	for {
		log.Printf("%v will create the draft, sleep %v\n", s.ExecTime.Format("2006-01-02 15:04:05"),
			time.Duration(s.ExecTime.Unix()-time.Now().Unix())*time.Second)
		time.Sleep(time.Duration(s.ExecTime.Unix()-time.Now().Unix()) * time.Second)
		s.Task.Start()
		s.UpdateExecuteTime()
	}
}

func (s *Schedule) UpdateExecuteTime() error {
	execTime, err := parseExecTime(s.Time)
	if err != nil {
		return fmt.Errorf("parse execute time error: %v", err)
	}
	s.ExecTime = execTime
	return nil
}

func NewSchedule(time string, task Task) (*Schedule, error) {
	s := Schedule{Time: time, Task: task}
	execTime, err := parseExecTime(time)
	if err != nil {
		return nil, fmt.Errorf("parse execute time error: %v", err)
	}
	s.ExecTime = execTime
	log.Printf("execute time is: %v\n", execTime)
	return &s, nil
}

func parseExecTime(execTime string) (time.Time, error) {
	now := time.Now()
	t, err := time.ParseInLocation("2006-01-02 15:04:05", now.Format("2006-01-02")+" "+execTime, time.Local)
	if err != nil {
		return time.Time{}, err
	}
	if t.Before(now) {
		t = t.Add(time.Hour * 24)
	}
	return t, nil
}
