package golanglearning

import (
	"fmt"
	"sync"
	"time"
)

// Task represents a work item
type Task struct {
	ID        int
	Priority  int
	Duration  time.Duration
	DependsOn []int
}

// TaskManager handles task execution and dependencies
type TaskManager struct {
	tasks     []Task
	completed map[int]chan struct{}
	wg        sync.WaitGroup
}

func NewTaskManager(tasks []Task) *TaskManager {
	completed := make(map[int]chan struct{})
	for _, task := range tasks {
		completed[task.ID] = make(chan struct{})
	}
	return &TaskManager{
		tasks:     tasks,
		completed: completed,
	}
}

func (tm *TaskManager) waitForDependencies(task Task) {
	for _, depID := range task.DependsOn {
		<-tm.completed[depID]
	}
}

func (tm *TaskManager) processTask(task Task, results chan<- string) {
	defer tm.wg.Done()
	defer close(tm.completed[task.ID])

	// Wait for dependencies to complete
	tm.waitForDependencies(task)

	fmt.Printf("Starting task %d (Priority: %d)\n", task.ID, task.Priority)
	time.Sleep(task.Duration)
	results <- fmt.Sprintf("Task %d completed in %v", task.ID, task.Duration)
}
// all will have like a common tracker and like global results thing
//and to check if all  tasks of group are completed we need to have this wait group
// and we need to have a way to track if all tasks are completed

func SimulateGoRoutines() {
	// Define tasks with dependencies
	tasks := []Task{
		{ID: 1, Priority: 1, Duration: time.Second * 2, DependsOn: []int{}},
		{ID: 2, Priority: 2, Duration: time.Second * 3, DependsOn: []int{1}},
		{ID: 3, Priority: 1, Duration: time.Second * 1, DependsOn: []int{}},
		{ID: 4, Priority: 3, Duration: time.Second * 2, DependsOn: []int{2, 3}},
		{ID: 5, Priority: 1, Duration: time.Second * 1, DependsOn: []int{}},
	}

	taskManager := NewTaskManager(tasks)
	results := make(chan string, len(tasks))

	for _, task := range tasks {
		taskManager.wg.Add(1)
		go taskManager.processTask(task, results)
	}

	go func() {
		taskManager.wg.Wait()
		close(results)
	}()

	fmt.Println("----------------------")
	for result := range results {
		fmt.Println(result)
	}

	fmt.Println("\nAll tasks completed!")
}
