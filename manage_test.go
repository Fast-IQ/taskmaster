//go:build windows

package taskmaster

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func TestLocalConnect(t *testing.T) {
	taskService, err := Connect()
	if err != nil {
		t.Fatal(err)
	}
	taskService.Disconnect()
}

func TestCreateTask(t *testing.T) {
	var err error
	taskService, err := Connect()
	assert.NoError(t, err)

	defer taskService.Disconnect()

	// test ExecAction
	execTaskDef := taskService.NewTaskDefinition()
	popCalc := ExecAction{
		Command: "calc.exe",
	}
	execTaskDef.AddAction(popCalc)

	_, _, err = taskService.CreateTask("\\Sima-Land\\ExecAction", execTaskDef, true)
	assert.NoError(t, err)

	// test ComHandlerAction
	comHandlerDef := taskService.NewTaskDefinition()
	comHandlerDef.AddAction(ComHandlerAction{
		ClassID: "{F0001111-0000-0000-0000-0000FEEDACDC}",
	})

	_, _, err = taskService.CreateTask("\\Taskmaster\\ComHandlerAction", comHandlerDef, true)
	assert.NoError(t, err)

	// test BootTrigger
	bootTriggerDef := taskService.NewTaskDefinition()
	bootTriggerDef.AddAction(popCalc)
	bootTriggerDef.AddTrigger(BootTrigger{
		TaskTrigger: TaskTrigger{
			StartBoundary: time.Now(),
			Enabled:       false,
		},
	})
	_, _, err = taskService.CreateTask("\\Taskmaster\\BootTrigger", bootTriggerDef, true)
	assert.NoError(t, err)

	// test DailyTrigger
	dailyTriggerDef := taskService.NewTaskDefinition()
	dailyTriggerDef.AddAction(popCalc)
	dailyTriggerDef.AddTrigger(DailyTrigger{
		DayInterval: EveryDay,
		TaskTrigger: TaskTrigger{
			StartBoundary: time.Now(),
		},
	})
	_, _, err = taskService.CreateTask("\\Taskmaster\\DailyTrigger", dailyTriggerDef, true)
	assert.NoError(t, err)

	// test EventTrigger
	eventTriggerDef := taskService.NewTaskDefinition()
	eventTriggerDef.AddAction(popCalc)
	subscription := "<QueryList> <Query Id='1'> <Select Command='System'>*[System/Level=2]</Select></Query></QueryList>"
	eventTriggerDef.AddTrigger(EventTrigger{
		Subscription: subscription,
	})
	_, _, err = taskService.CreateTask("\\Taskmaster\\EventTrigger", eventTriggerDef, true)
	assert.NoError(t, err)

	// test IdleTrigger
	idleTriggerDef := taskService.NewTaskDefinition()
	idleTriggerDef.AddAction(popCalc)
	idleTriggerDef.AddTrigger(IdleTrigger{})
	_, _, err = taskService.CreateTask("\\Taskmaster\\IdleTrigger", idleTriggerDef, true)
	assert.NoError(t, err)

	// test LogonTrigger
	logonTriggerDef := taskService.NewTaskDefinition()
	logonTriggerDef.AddAction(popCalc)
	logonTriggerDef.AddTrigger(LogonTrigger{})
	_, _, err = taskService.CreateTask("\\Taskmaster\\LogonTrigger", logonTriggerDef, true)
	assert.NoError(t, err)

	// test MonthlyDOWTrigger
	monthlyDOWTriggerDef := taskService.NewTaskDefinition()
	monthlyDOWTriggerDef.AddAction(popCalc)
	monthlyDOWTriggerDef.AddTrigger(MonthlyDOWTrigger{
		DaysOfWeek:   Monday | Friday,
		WeeksOfMonth: First,
		MonthsOfYear: January | February,
		TaskTrigger: TaskTrigger{
			StartBoundary: time.Now(),
		},
	})
	_, _, err = taskService.CreateTask("\\Taskmaster\\MonthlyDOWTrigger", monthlyDOWTriggerDef, true)
	assert.NoError(t, err)

	// test MonthlyTrigger
	monthlyTriggerDef := taskService.NewTaskDefinition()
	monthlyTriggerDef.AddAction(popCalc)
	monthlyTriggerDef.AddTrigger(MonthlyTrigger{
		DaysOfMonth:  3,
		MonthsOfYear: February | March,
		TaskTrigger: TaskTrigger{
			StartBoundary: time.Now(),
		},
	})
	_, _, err = taskService.CreateTask("\\Taskmaster\\MonthlyTrigger", monthlyTriggerDef, true)
	assert.NoError(t, err)

	// test RegistrationTrigger
	registrationTriggerDef := taskService.NewTaskDefinition()
	registrationTriggerDef.AddAction(popCalc)
	registrationTriggerDef.AddTrigger(RegistrationTrigger{})
	_, _, err = taskService.CreateTask("\\Taskmaster\\RegistrationTrigger", registrationTriggerDef, true)
	assert.NoError(t, err)

	// test SessionStateChangeTrigger
	sessionStateChangeTriggerDef := taskService.NewTaskDefinition()
	sessionStateChangeTriggerDef.AddAction(popCalc)
	sessionStateChangeTriggerDef.AddTrigger(SessionStateChangeTrigger{
		StateChange: TASK_SESSION_LOCK,
	})
	_, _, err = taskService.CreateTask("\\Taskmaster\\SessionStateChangeTrigger", sessionStateChangeTriggerDef, true)
	assert.NoError(t, err)

	// test TimeTrigger
	timeTriggerDef := taskService.NewTaskDefinition()
	timeTriggerDef.AddAction(popCalc)
	timeTriggerDef.AddTrigger(TimeTrigger{
		TaskTrigger: TaskTrigger{
			StartBoundary: time.Now(),
		},
	})
	_, _, err = taskService.CreateTask("\\Taskmaster\\TimeTrigger", timeTriggerDef, true)
	assert.NoError(t, err)

	// test WeeklyTrigger
	weeklyTriggerDef := taskService.NewTaskDefinition()
	weeklyTriggerDef.AddAction(popCalc)
	weeklyTriggerDef.AddTrigger(WeeklyTrigger{
		DaysOfWeek:   Tuesday | Thursday,
		WeekInterval: EveryOtherWeek,
		TaskTrigger: TaskTrigger{
			StartBoundary: time.Now(),
		},
	})
	_, _, err = taskService.CreateTask("\\Taskmaster\\WeeklyTrigger", weeklyTriggerDef, true)
	assert.NoError(t, err)

	// test trying to create task where a task at the same path already exists and the 'overwrite' is set to false
	_, taskCreated, err := taskService.CreateTask("\\Taskmaster\\TimeTrigger", timeTriggerDef, false)
	assert.NoError(t, err)
	assert.EqualValues(t, taskCreated, false, "task shouldn't have been created")
	/*if taskCreated {
		t.Fatal("task shouldn't have been created")
	}*/
}

func TestUpdateTask(t *testing.T) {
	taskService, err := Connect()
	if err != nil {
		t.Fatal(err)
	}
	testTask := createTestTask(taskService)
	defer taskService.Disconnect()

	testTask.Definition.RegistrationInfo.Author = "Big Chungus"
	_, err = taskService.UpdateTask("\\Taskmaster\\TestTask", testTask.Definition)
	if err != nil {
		t.Fatal(err)
	}

	testTask, err = taskService.GetRegisteredTask("\\Taskmaster\\TestTask")
	if err != nil {
		t.Fatal(err)
	}
	if testTask.Definition.RegistrationInfo.Author != "Big Chungus" {
		t.Fatal("task was not updated")
	}
}

func TestGetRegisteredTasks(t *testing.T) {
	taskService, err := Connect()
	if err != nil {
		t.Fatal(err)
	}
	defer taskService.Disconnect()

	rtc, err := taskService.GetRegisteredTasks()
	if err != nil {
		t.Fatal(err)
	}
	rtc.Release()
}

func TestGetTaskFolders(t *testing.T) {
	taskService, err := Connect()
	if err != nil {
		t.Fatal(err)
	}
	defer taskService.Disconnect()

	tf, err := taskService.GetTaskFolders()
	if err != nil {
		t.Fatal(err)
	}
	tf.Release()
}

func TestDeleteTask(t *testing.T) {
	taskService, err := Connect()
	if err != nil {
		t.Fatal(err)
	}
	createTestTask(taskService)
	defer taskService.Disconnect()

	err = taskService.DeleteTask("\\Taskmaster\\TestTask")
	if err != nil {
		t.Fatal(err)
	}

	deletedTask, err := taskService.GetRegisteredTask("\\Taskmaster\\TestTask")
	if err == nil {
		t.Fatal("task shouldn't still exist")
		return
	}
	deletedTask.Release()
}

func TestDeleteFolder(t *testing.T) {
	taskService, err := Connect()
	if err != nil {
		t.Fatal(err)
	}
	createTestTask(taskService)
	defer taskService.Disconnect()

	var folderDeleted bool
	folderDeleted, err = taskService.DeleteFolder("\\Taskmaster", false)
	if err != nil {
		t.Fatal(err)
	}
	if folderDeleted == true {
		t.Error("folder shouldn't have been deleted")
	}

	folderDeleted, err = taskService.DeleteFolder("\\Taskmaster", true)
	if err != nil {
		t.Fatal(err)
	}
	if folderDeleted == false {
		t.Error("folder should have been deleted")
	}

	tasks, err := taskService.GetRegisteredTasks()
	if err != nil {
		t.Fatal(err)
	}
	taskmasterFolder, err := taskService.GetTaskFolder("\\Taskmaster")
	if err == nil {
		t.Fatal("folder shouldn't exist")
	}
	if taskmasterFolder.Name != "" {
		t.Error("folder struct should be defaultly constructed")
	}
	for _, task := range tasks {
		if strings.Split(task.Path, "\\")[1] == "Taskmaster" {
			t.Error("task should've been deleted")
		}
	}
}
