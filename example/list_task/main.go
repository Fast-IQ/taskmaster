package main

import (
	"fmt"
	"github.com/Fast-IQ/taskmaster"
	"os"
	"strings"
)

func main() {
	taskSrv, err := taskmaster.Connect()
	if err != nil {
		panic(err)
	}
	defer taskSrv.Disconnect()
	taskFolders, err := taskSrv.GetTaskFolders()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer taskFolders.Release()

	_ = ListFolder(&taskFolders, 1)
}

func ListFolder(taskFolders *taskmaster.TaskFolder, level int) error {
	for _, taskFolder := range taskFolders.SubFolders {
		fmt.Println(strings.Repeat("  ", level) + taskFolder.Name)
		if len(taskFolder.SubFolders) > 0 {
			_ = ListFolder(taskFolder, level+1)
		}
		for _, regTask := range taskFolder.RegisteredTasks {
			fmt.Println(strings.Repeat("  ", level+1) + "Task: " + regTask.Name)
		}
	}
	return nil
}
