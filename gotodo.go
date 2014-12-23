package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/jack-zh/gotodo/task"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var noAct = errors.New("error")

var (
	file = flag.String("file", defaultFile(".zgotodo", "TODO"), "file in which to store tasks")
	now  = flag.Bool("now", false, "when adding, insert at head")
	done = flag.Bool("done", false, "don't actually add; just append to log file")
)

func defaultFile(name, env string) string {
	if f := os.Getenv(env); f != "" {
		return f
	}
	return filepath.Join(os.Getenv("HOME"), name)
}

const usage = `Usage:
	gotodo list
		Show all tasks
	gotodo list N
		Show task N
	gotodo rm N
		Remove task N
	gotodo done N
		Done task N
	gotodo undo N
		Undo task N
	gotodo doing N
		Doing task N
	gotodo add ...
		Add task to list
`

func printTask(task string, i string) {
	if strings.HasPrefix(task, "0") {
		task = strings.Replace(task, "0", "[Future]", 1)
	}
	if strings.HasPrefix(task, "1") {
		task = strings.Replace(task, "1", "[Doing ]", 1)
	}
	if strings.HasPrefix(task, "2") {
		task = strings.Replace(task, "2", "[Done  ]", 1)
	}
	fmt.Printf("%2s: %s\n", i, task)
}

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, usage)
		flag.PrintDefaults()
	}
	flag.Parse()

	list := task.NewList(*file)
	a, n := flag.Arg(0), len(flag.Args())

	err := noAct
	switch {
	case a == "list" && n == 1:
		var tasks []string
		tasks, err = list.Get()
		for i := 0; i < len(tasks); i++ {
			printTask(tasks[i], strconv.Itoa(i+1))
		}

	case a == "list" && n == 2:
		i, err2 := strconv.Atoi(flag.Arg(1))
		if err2 != nil {
			fmt.Fprint(os.Stdout, usage)
			break
		}
		var task string
		task, err = list.GetTask(i - 1)
		if err == nil {
			printTask(task, strconv.Itoa(i))
		}
	case a == "rm" && n == 2:
		i, err2 := strconv.Atoi(flag.Arg(1))
		if err2 != nil {
			fmt.Fprint(os.Stdout, usage)
			break
		}
		err = list.RemoveTask(i - 1)
		if err != nil {
			break
		}
	case a == "add" && n > 1:
		t := strings.Join(flag.Args()[1:], " ")
		err = list.AddTask(t, *now)

	case a == "doing" && n == 2:
		i, err3 := strconv.Atoi(flag.Args()[1])
		if err3 != nil {
			fmt.Fprint(os.Stdout, usage)
			break
		}
		err = list.DoingTask(i - 1)

	case a == "done" && n == 2:
		i, err4 := strconv.Atoi(flag.Args()[1])
		if err4 != nil {
			fmt.Fprint(os.Stdout, usage)
			break
		}
		err = list.DoneTask(i - 1)
	case a == "undo" && n == 2:
		i, err5 := strconv.Atoi(flag.Args()[1])
		if err5 != nil {
			fmt.Fprint(os.Stdout, usage)
			break
		}
		err = list.UndoTask(i - 1)
	default:
		fmt.Fprint(os.Stdout, usage)
	}
	if err != nil {
		fmt.Println(err)
	} else {
		if a != "list" {
			fmt.Println("\nSuccess!")
		}
	}
}
