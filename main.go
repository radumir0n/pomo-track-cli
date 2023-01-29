package main

import (
	"fmt"
	"time"
)

const taskNamePrompt = "Enter the name of your task: "
const workTimePrompt = "Enter work time in minutes: "
const breakTimePrompt = "Enter break time in minutes: "
const cycleCountPrompt = "Enter how many cycles to work: "

const defaultWorkTime = 25 * 60
const defaultBreakTime = 5 * 60
const defaultCycleCount = 4

func main() {
	task := promptUser()
	
	prevElapsedTime := 0
	for {
		elapsedTime := task.getElapsedTimeInSeconds()

		if elapsedTime != prevElapsedTime {
			time := getCountdownTimer(task)
			printRunningTask(task.taskName, time)

			prevElapsedTime = elapsedTime
		}
	}
}

func printRunningTask(task string, t Time) {
	// get countdown time in hours-mins-seconds (ex: 25:00)
	// run a loop that counts down the time until 00:00
	// when the timer is up switch to break mode and start new timer (ex: 05:00)
	// this marks the end of a cycle, check how many cycles there are, do the book keeping, redo the timer if there are more
	fmt.Printf("\r%s: %02d:%02d", task, t.minutes, t.seconds)
}

func getCountdownTimer(task Task) Time {
	elapsedSeconds := task.getElapsedTimeInSeconds()
	workTime := task.workTime * 60
	timeRemaining := workTime - elapsedSeconds
	
	var t Time
	t.minutes = timeRemaining / 60
	t.seconds = timeRemaining - t.minutes * 60

	return t
}

func checkPromptNumericInputs(message string, defaultValue int) int {
	var input int

	fmt.Print(message)
	fmt.Scanln(&input)

	if input == 0 {
		input = defaultValue 
	}

	return input
}

func promptUser() Task {
	// var taskName string
	
	// for {
	// 	fmt.Print(taskNamePrompt)
	// 	fmt.Scanln(&taskName)

	// 	if taskName != "" {
	// 		break
	// 	}
	// }

	// t := Task {
	// 	taskName: taskName,
	// 	startTime: time.Now(),
	// 	workTime: checkPromptNumericInputs(workTimePrompt, defaultWorkTime),
	// 	breakTime: checkPromptNumericInputs(breakTimePrompt, defaultBreakTime),
	// 	cycleCount: checkPromptNumericInputs(cycleCountPrompt, defaultCycleCount),
	// 	isWorkTime: true,
	// }

	return Task {
		taskName: "test",
		startTime: time.Now(),
		workTime: 25,
		breakTime: 5,
		cycleCount: 4,
		isWorkTime: true,
	}

	// return t
}