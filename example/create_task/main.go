package main

import (
	"fmt"
	"github.com/Fast-IQ/taskmaster"
	"github.com/rickb777/date/period"
	"time"
)

func main() {
	_ = CreateTaskRun2Action()
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
		Command: "notepad.exe",
	}
	def.AddAction(act)

	_, ok, err := taskSrv.CreateTask("\\Taskmaster\\NewTask2", def, true)
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

	tf, err := taskSrv.GetTaskFolder("\\Taskmaster")
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

func CreateTaskRun2Action() error {
	taskSrv, err := taskmaster.Connect()
	if err != nil {
		panic(err)
	}
	defer taskSrv.Disconnect()

	def := taskSrv.NewTaskDefinition()
	//def.Principal.UserID = "Administrator"
	//def.Principal.GroupID = "Administrators"
	def.Principal.LogonType = taskmaster.TASK_LOGON_INTERACTIVE_TOKEN_OR_PASSWORD
	def.AddAction(taskmaster.ExecAction{
		Command: "taskkill.exe",
		Args:    "/f /im:run.exe",
	})
	def.AddAction(taskmaster.ExecAction{
		Command: "run.exe",
		Args:    "",
	})
	def.AddTrigger(taskmaster.DailyTrigger{
		TaskTrigger: taskmaster.TaskTrigger{
			StartBoundary: time.Date(2024, 01, 01, 12, 00, 05, 0, time.Local),
			Enabled:       true,
		},
		DayInterval: 1,
		RandomDelay: period.NewHMS(1, 0, 0),
	})
	_, _, err = taskSrv.CreateTask("\\Taskmaster\\Run_Start", def, true)
	if err != nil {
		return err
	}
	return nil
}
