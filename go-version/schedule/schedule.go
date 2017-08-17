package schedule

type Schedule struct {
	time string
	task Task
}

type Task interface {
	Start()
}

func (s *Schedule) Run() {

}
