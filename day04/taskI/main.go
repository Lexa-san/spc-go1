package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	var taskNumber int
	scanner.Scan()
	taskNumber, _ = strconv.Atoi(scanner.Text())

	tasks := make([]string, taskNumber, taskNumber)

	for i := 0; i < taskNumber; i++ {
		scanner.Scan()
		tasks[i] = scanner.Text()
	}

	var reminderNumber int
	scanner.Scan()
	reminderNumber, _ = strconv.Atoi(scanner.Text())

	reminders := make([]string, reminderNumber, reminderNumber)

	var taskIndex int
	for i := 0; i < reminderNumber; i++ {
		scanner.Scan()
		taskIndex, _ = strconv.Atoi(scanner.Text())
		reminders[i] = tasks[taskIndex-1]
	}

	for _, val := range reminders {
		fmt.Println(val)
	}
}
