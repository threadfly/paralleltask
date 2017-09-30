package concurrenttask

import (
	"time"
)

const (
	signal_container_proportion_on_go_routine_num = 4
)

type ParallelTask interface {
	Run()
	IsSucc() bool
}

type ParallelTaskManager struct {
	taskIntervalNum int
	TaskList        []ParallelTask
	RunTime         time.Duration
	CallBack        func(*ParallelTaskManager)
	goroutineSum    int
	hasRun          bool
}

type RUNNING_MODE int

const (
	RUNNING_WITH_PARALLEL RUNNING_MODE = 0
	RUNNING_WITH_LINE     RUNNING_MODE = 1
)
