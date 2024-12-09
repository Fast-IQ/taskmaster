package main

import (
	"fmt"
	"github.com/Fast-IQ/taskmaster"
	"time"
)

func main() {
	//createT(nil)
	GetTask()
}

func createT(trig taskmaster.Trigger) {
	taskSrv, err := taskmaster.Connect()
	if err != nil {
		panic(err)
	}
	defer taskSrv.Disconnect()
	def := taskSrv.NewTaskDefinition()

	def.AddTrigger(taskmaster.DailyTrigger{
		DayInterval: taskmaster.EveryDay,
		TaskTrigger: taskmaster.TaskTrigger{
			StartBoundary: time.Now(),
			Enabled:       true,
		},
	})

	def.AddTrigger(trig)

	act := taskmaster.ExecAction{
		Path: "notepad.exe",
	}
	def.AddAction(act)

	_, ok, err := taskSrv.CreateTask("\\Sima-Land\\NewTask2", def, true)
	if ok {
		fmt.Println("task created")
	} else if err != nil {
		fmt.Println("task not created, error " + err.Error())
	}
}

func GetTask() {
	taskSrv, err := taskmaster.Connect()
	if err != nil {
		panic(err)
	}
	defer taskSrv.Disconnect()

	tf, err := taskSrv.GetTaskFolder("\\Sima-Land")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	rtc := tf.RegisteredTasks
	for _, t := range rtc {
		if t.Name == "NewTask" {
			createT(t.Definition.Triggers[0])
			fmt.Println(t.Definition.Triggers[0])
		}
	}

}
