package concurrenttask

import (
	"time"
	uflog "uframework/log"
)

func NewParallelTaskManager(taskIntervalNum int) *ParallelTaskManager {
	return &ParallelTaskManager{
		taskIntervalNum: taskIntervalNum,
		hasRun:          false,
	}
}

func NewParallelTaskManagerV2(gorountinNum int) *ParallelTaskManager {
	return &ParallelTaskManager{
		goroutineSum: gorountinNum,
		hasRun:       false,
	}
}

func (this *ParallelTaskManager) PushTask(task ParallelTask) {
	if this.TaskList == nil {
		this.TaskList = make([]ParallelTask, 0, this.taskIntervalNum)
	}

	this.TaskList = append(this.TaskList, task)
}

func (this *ParallelTaskManager) Process(mode RUNNING_MODE) {
	if this.hasRun {
		return
	}

	begin := time.Now()
	if mode == RUNNING_WITH_PARALLEL {
		this.concurrentProcess()
	} else {
		this.lineProcess()
	}
	end := time.Now()
	this.RunTime = end.Sub(begin)
	if this.CallBack != nil {
		this.CallBack(this)
	}

	this.hasRun = true
}

func (this *ParallelTaskManager) lineProcess() {
	for _, task := range this.TaskList {
		task.Run()
	}
}

func (this *ParallelTaskManager) concurrentProcess() {
	if len(this.TaskList) == 0 {
		return
	}

	// caculate go rountine number
	goroutinNum := 0
	taskIntervalNum := 0

	if this.goroutineSum == 0 && this.taskIntervalNum != 0 {
		taskIntervalNum = this.taskIntervalNum
		switch {
		case len(this.TaskList)%this.taskIntervalNum == 0:
			goroutinNum = len(this.TaskList) / this.taskIntervalNum
		case len(this.TaskList) < this.taskIntervalNum:
			goroutinNum = 1
		default:
			goroutinNum = len(this.TaskList)/this.taskIntervalNum + 1
		}
	} else if this.goroutineSum != 0 && this.taskIntervalNum == 0 {
		goroutinNum = this.goroutineSum
		switch {
		case len(this.TaskList)%this.goroutineSum == 0:
			taskIntervalNum = len(this.TaskList) / this.goroutineSum
		case len(this.TaskList) < this.goroutineSum:
			taskIntervalNum = 1
			goroutinNum = len(this.TaskList)
		default:
			taskIntervalNum = len(this.TaskList) / this.goroutineSum
			goroutinNum = this.goroutineSum + 1
		}
	} else {
		return
	}

	signalContainer := goroutinNum/signal_container_proportion_on_go_routine_num + 1
	signal := make(chan int, signalContainer)

	// concurrent deal
	uflog.INFOF("signalContainer:%d goroutinNum:%d taskIntervalNum:%d this.TaskList len:%d", signalContainer, goroutinNum, taskIntervalNum, len(this.TaskList))
	for i := 0; i < goroutinNum; i++ {
		go func(index int) {
			begin := index * taskIntervalNum
			end := begin + taskIntervalNum - 1
			if end >= len(this.TaskList) {
				end = len(this.TaskList) - 1
			}

			uflog.INFOF("index:%d begin:%d end:%d TaskList:%#v", index, begin, end, this.TaskList)
			for _, task := range this.TaskList[begin : end+1] {
				task.Run()
			}
			signal <- 1
		}(i)
	}

	index := 0
	for range signal {
		index++
		if index == goroutinNum {
			break
		}
	}
}

func (this *ParallelTaskManager) RegisterEndCallBack(callback func(*ParallelTaskManager)) {
	this.CallBack = callback
}
