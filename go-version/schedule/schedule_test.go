package schedule

import (
	"log"
	"testing"
	"time"
)

type MockTask struct {
}

func (t *MockTask) Start() {
	log.Println("start mock task")
}

func Test_schedule(t *testing.T) {
	now := time.Now()
	time := now.Add(2 * time.Second).Format("15:04:05")
	t.Logf("time is %v", time)
	s, err := NewSchedule(time, &MockTask{})
	if err != nil {
		t.Fatalf("%v", err)
	}
	s.Run()
}
