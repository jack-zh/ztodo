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

var noAct = errors.New("didn't act")

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
	gotodo rm N
		Remove task N
	gotodo add ...
		Add task to list
`

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
			fmt.Printf("%2d: %s\n", i+1, tasks[i])
		}
	case a == "rm" && n == 2:
		i, err2 := strconv.Atoi(flag.Arg(1))
		if err2 != nil && n == 2 {
			break
		}
		err = list.RemoveTask(i - 1)
		if err != nil || n == 2 {
			break
		}
	case a == "add" && n > 1:
		t := strings.Join(flag.Args()[1:], " ")
		err = list.AddTask(t, *now)
	default:
		fmt.Fprint(os.Stdout, usage)
	}
}
