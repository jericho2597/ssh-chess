package app

import (
	"container/heap"
	"sync"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

var SchedulerService *Scheduler

type Task struct {
	SessionID string
	Msg       tea.Msg
	ExecuteAt time.Time
}

// taskHeap is a min-heap of Tasks
type taskHeap []*Task

func (h taskHeap) Len() int           { return len(h) }
func (h taskHeap) Less(i, j int) bool { return h[i].ExecuteAt.Before(h[j].ExecuteAt) }
func (h taskHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *taskHeap) Push(x interface{}) {
	*h = append(*h, x.(*Task))
}

func (h *taskHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Scheduler struct {
	tasks          taskHeap
	mutex          sync.Mutex
	cond           *sync.Cond
	quit           chan struct{}
	sessionManager *SessionManager
}

func NewScheduler(sm *SessionManager) {
	SchedulerService = &Scheduler{
		tasks:          make(taskHeap, 0),
		quit:           make(chan struct{}),
		sessionManager: sm,
	}
	SchedulerService.cond = sync.NewCond(&SchedulerService.mutex)
	heap.Init(&SchedulerService.tasks)
	go SchedulerService.run()
}

func (s *Scheduler) Schedule(sessionID string, msg tea.Msg, delay time.Duration) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	task := &Task{
		SessionID: sessionID,
		Msg:       msg,
		ExecuteAt: time.Now().Add(delay),
	}
	heap.Push(&s.tasks, task)
	s.cond.Signal()
	return nil
}

func (s *Scheduler) run() {
	for {
		s.mutex.Lock()
		for s.tasks.Len() == 0 {
			select {
			case <-s.quit:
				s.mutex.Unlock()
				return
			default:
				s.cond.Wait()
			}
		}

		now := time.Now()
		task := s.tasks[0]

		if now.Before(task.ExecuteAt) {
			timer := time.NewTimer(task.ExecuteAt.Sub(now))
			s.mutex.Unlock()
			select {
			case <-timer.C:
				// Time to execute
			case <-s.quit:
				timer.Stop()
				return
			}
			s.mutex.Lock()
		}

		heap.Pop(&s.tasks)
		s.mutex.Unlock()

		go s.executeTask(task)
	}
}

func (s *Scheduler) executeTask(task *Task) {
	s.sessionManager.Send(task.SessionID, task.Msg)
}

// Shutdown stops the scheduler
func (s *Scheduler) Shutdown() {
	close(s.quit)
	s.cond.Broadcast()
}
