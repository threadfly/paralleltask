package main

import (
	"fmt"

	tasks "gitlab.ucloudadmin.com/TreasureGo/paralleltask"
)

type myTask struct {
	handler func()
}

func NewMyTask(h func()) *myTask {
	return &myTask{
		handler: h,
	}
}

func (this *myTask) Run() {
	if this.handler != nil {
		this.handler()
	}
}

func (this *myTask) IsSucc() bool {
	return true
}

func main() {
	taskMgr := tasks.NewParallelTaskManager(1)

	f := func() {
		for i := 0; i < 1000000000; i++ {
		}
	}

	taskMgr.PushTask(NewMyTask(f))
	taskMgr.PushTask(NewMyTask(f))
	taskMgr.PushTask(NewMyTask(f))
	taskMgr.PushTask(NewMyTask(f))
	taskMgr.PushTask(NewMyTask(f))
	taskMgr.PushTask(NewMyTask(f))
	taskMgr.PushTask(NewMyTask(f))
	taskMgr.PushTask(NewMyTask(f))
	taskMgr.PushTask(NewMyTask(f))
	taskMgr.PushTask(NewMyTask(f))
	taskMgr.PushTask(NewMyTask(f))
	taskMgr.PushTask(NewMyTask(f))
	taskMgr.PushTask(NewMyTask(f))

	taskMgr.Process(tasks.RUNNING_WITH_LINE)
	fmt.Printf("line running mode take time:%s\n", taskMgr.RunTime)

	taskMgr.Process(tasks.RUNNING_WITH_CONCURRENT)
	fmt.Printf("concurrent running mode take time:%s\n", taskMgr.RunTime)
}
