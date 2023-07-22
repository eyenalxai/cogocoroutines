package scheduler

type Task struct {
	run      func() bool
	callback func()
}

type Scheduler struct {
	tasks []Task
}

func (s *Scheduler) AddTask(t Task) {
	s.tasks = append(s.tasks, t)
}

func (s *Scheduler) Run() {
	for len(s.tasks) > 0 {
		task := s.tasks[0]
		s.tasks = s.tasks[1:]

		if !task.run() {
			s.tasks = append(s.tasks, task)
		} else {
			if task.callback != nil {
				task.callback()
			}
		}
	}
}

func NewTask(run func() bool, optionalCallback ...func()) Task {
	var callback func()
	if len(optionalCallback) > 0 {
		callback = optionalCallback[0]
	}

	return Task{
		run:      run,
		callback: callback,
	}
}
