package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var taskTraker = make(map[string]map[string]string)
var isWorking bool = false

//var actions = [7]

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	var action string
	var context []string
	//var isWorking = false

taskTrakerActionLoop:
	for {
		action, context = listenAction(scanner)
		//fmt.Println(action, context)

		switch action {
		case "StartApp":
			actionStartApp()
		case "Quit":
			break taskTrakerActionLoop
		case "Print":
			if !isWorking {
				continue
			}
			actionPrint()
		case "Add":
			if !isWorking {
				continue
			}
			actionAdd(context)
		case "Del":
			if !isWorking {
				continue
			}
		case "Find":
			if !isWorking {
				continue
			}
		}
	}

}

// get action name as string and additional data of action in slice
func listenAction(scanner *bufio.Scanner) (action string, context []string) {
	scanner.Scan()
	arr := strings.Split(scanner.Text(), " ")

	if len(arr) == 1 {
		return arr[0], nil
	}
	if len(arr) > 1 {
		return arr[0], arr[1:]
	}

	return "", nil
}

// transform "1-2-3" format to "0001-02-03"
func normalizeData(date string) string {
	arr := strings.Split(date, "-")
	if len(arr) != 3 {
		return date
	}
	var arrInt = make([]int, len(arr), len(arr))
	for i, v := range arr {
		arrInt[i], _ = strconv.Atoi(v)
	}
	return fmt.Sprintf("%.4d-%.2d-%.2d", arrInt[0], arrInt[1], arrInt[2])
}

// Action StartApp
func actionStartApp() {
	isWorking = true
}

func actionAdd(context []string) {
	var date, task string = context[0], context[1]
	date = normalizeData(date)

	if _, ok := taskTraker[date]; !ok {
		taskTraker[date] = make(map[string]string)
	}
	taskTraker[date][task] = task
}

func actionPrint() {
	// sort dates
	dates := make([]string, 0, len(taskTraker))
	for date := range taskTraker {
		dates = append(dates, date)
	}
	sort.Strings(dates)

	var tasks []string
	for _, date := range dates {

		// sort tasks of every date
		tasks = make([]string, 0, len(taskTraker[date]))
		for task := range taskTraker[date] {
			tasks = append(tasks, task)
		}
		sort.Strings(tasks)

		// final print
		for _, task := range tasks {
			fmt.Println(date, task)
		}

	}
}
