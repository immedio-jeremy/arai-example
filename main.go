package main

import (
	"example.com/arai-example/developer"
	"example.com/arai-example/task"
)

func main() {

	task.PerformBackendTask(developer.BackendDeveloper{})

	task.PerformBackendTask(developer.FullStackDeveloper{})

	task.PerformFrontendTask(developer.FullStackDeveloper{})

	task.PerformFullStackTask(developer.FullStackDeveloper{})

	task.PerformFrontendTask(developer.FrontendClone{})

	task.PerformFrontendTask(developer.FrontendRobot{})
}
