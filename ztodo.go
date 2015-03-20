package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/jack-zh/ztodo/task"
	"github.com/jack-zh/ztodo/utils"
)

var noAct = errors.New("")

var userconfig_filename = filepath.Join(os.Getenv("HOME"), ".ztodo", "userconfig.json")
var cloud_work_tasks_filename = filepath.Join(os.Getenv("HOME"), ".ztodo", "worktasks.json")
var cloud_backup_tasks_filename = filepath.Join(os.Getenv("HOME"), ".ztodo", "backuptasks.json")
var simple_tasks_filename = filepath.Join(os.Getenv("HOME"), ".ztodo", "simpletasks")

const version = "0.4.8"
const build_time = "2015-03-23"
const usage = `Incorrect Usage.

NAME:
   ztodo - a command line todo list!

USAGE:
   ztodo [global options] command [command options] [arguments...]

VERSION:
   version:` + version + " (" + build_time + ") build" + `

AUTHOR:
  Jack.z - <zzh.coder@qq.com>

COMMANDS:

	ztodo list|ls [verbose]    -- Show all tasks
	ztodo list|ls N [verbose]  -- Show task N
	ztodo rm|remove N          -- Remove task N
	ztodo done N               -- Done task N
	ztodo undo N               -- Undo task N
	ztodo doing N              -- Doing task N
	ztodo clean                -- Rm done task
	ztodo clear                -- Rm all task
	ztodo add ...              -- Add task to list

GLOBAL OPTIONS:
	ztodo version              -- Show ztodo version
	ztodo help                 -- Show usage
`

func printSimpleTask(t string, i string) {
	if strings.HasPrefix(t, "0") {
		t = strings.Replace(t, "0", "[New]", 1)
	}
	if strings.HasPrefix(t, "1") {
		t = strings.Replace(t, "1", "[Doing ]", 1)
	}
	if strings.HasPrefix(t, "2") {
		t = strings.Replace(t, "2", "[Done  ]", 1)
	}
	fmt.Printf("%2s: %s\n", i, t)
}

func dirCheck() error {
	var filename = filepath.Join(os.Getenv("HOME"), ".ztodo")
	finfo, err := os.Stat(filename)
	if err != nil {
		os.Mkdir(filename, os.ModePerm)
		return nil
	}
	if finfo.IsDir() {
		return nil
	} else {
		return errors.New("$HOME/.ztodo is a file not dir.")
	}
}

func printUsgaes() {
	fmt.Println("Have a nice day.\n")
	fmt.Fprint(os.Stdout, usage)
}

func main() {
	errdir := dirCheck()
	if errdir != nil {
		os.Exit(1)
	}
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, usage)
		flag.PrintDefaults()
	}
	flag.Parse()

	simplelist := task.SimpleNewList(simple_tasks_filename)
	cloudlist := task.CloudNewList(cloud_work_tasks_filename, cloud_backup_tasks_filename, userconfig_filename)
	a, n := flag.Arg(0), len(flag.Args())

	a = strings.ToLower(a)
	if a == "ls" {
		a = "list"
	} else if a == "remove" {
		a = "rm"
	} else if a == "simplels" {
		a = "simplelist"
	} else if a == "simpleremove" {
		a = "simplerm"
	}

	err := noAct
	switch {
	case a == "version" && n == 1:
		fmt.Println("ztodo version " + version + " (" + build_time + ") build")
		err = nil

	case a == "help" && n == 1:
		fmt.Println(usage)
		err = nil

	case a == "simplelist" && n == 1:
		var tasks []string
		tasks, err = simplelist.SimpleGet()
		for i := 0; i < len(tasks); i++ {
			printSimpleTask(tasks[i], strconv.Itoa(i+1))
		}
	case a == "simplelist" && n == 2:
		i, err2 := strconv.Atoi(flag.Arg(1))
		if err2 != nil {
			printUsgaes()
			break
		}
		var task string
		task, err = simplelist.SimpleGetTask(i - 1)
		if err == nil {
			printSimpleTask(task, strconv.Itoa(i))
		}
	case a == "simplerm" && n == 2:
		i, err2 := strconv.Atoi(flag.Arg(1))
		if err2 != nil {
			printUsgaes()
			break
		}
		err = simplelist.SimpleRemoveTask(i - 1)
		if err != nil {
			break
		}
	case a == "simpleadd" && n > 1:
		t := strings.Join(flag.Args()[1:], " ")
		err = simplelist.SimpleAddTask(t)
		err = cloudlist.CloudAddTask(t)

	case a == "simpledoing" && n == 2:
		i, err3 := strconv.Atoi(flag.Args()[1])
		if err3 != nil {
			printUsgaes()
			break
		}
		err = simplelist.SimpleDoingTask(i - 1)

	case a == "simpledone" && n == 2:
		i, err4 := strconv.Atoi(flag.Args()[1])
		if err4 != nil {
			printUsgaes()
			break
		}
		err = simplelist.SimpleDoneTask(i - 1)
	case a == "simpleundo" && n == 2:
		i, err5 := strconv.Atoi(flag.Args()[1])
		if err5 != nil {
			printUsgaes()
			break
		}
		err = simplelist.SimpleUndoTask(i - 1)
	case a == "simpleclean" && n == 1:
		err = simplelist.SimpleCleanTask()
	case a == "simpleclear" && n == 1:
		err = simplelist.SimpleClearTask()

	case a == "list" && n == 1:
		err = cloudlist.CloudGetAllWorkTaskByFile()
		if err == nil {
			cloudlist.CloudTasksPrint(-1)
		}

	case a == "list" && n == 2:
		if flag.Arg(1) == "verbose" {
			err = cloudlist.CloudGetAllWorkTaskByFile()
			if err == nil {
				cloudlist.CloudTasksPrintVerbose(-1)
			}
		} else {
			i, err2 := strconv.Atoi(flag.Arg(1))
			if err2 != nil {
				printUsgaes()
				break
			}
			err = cloudlist.CloudGetAllWorkTaskByFile()
			if err == nil {
				cloudlist.CloudTasksPrint(i)
			}
		}
	case a == "list" && n == 3:
		if flag.Arg(2) == "verbose" {
			i, err2 := strconv.Atoi(flag.Arg(1))
			if err2 != nil {
				printUsgaes()
				break
			}
			err = cloudlist.CloudGetAllWorkTaskByFile()
			if err == nil {
				cloudlist.CloudTasksPrintVerbose(i)
			}
		} else {
			printUsgaes()
		}

	case a == "rm" && n == 2:
		i, err2 := strconv.Atoi(flag.Arg(1))
		if err2 != nil {
			printUsgaes()
			break
		}
		err = cloudlist.CloudRemoveTask(i)
		if err != nil {
			break
		}
	case a == "add" && n > 1:
		t := strings.Join(flag.Args()[1:], " ")
		err = cloudlist.CloudAddTask(t)

	case a == "doing" && n == 2:
		i, err3 := strconv.Atoi(flag.Args()[1])
		if err3 != nil {
			printUsgaes()
			break
		}
		err = cloudlist.CloudDoingTask(i)

	case a == "done" && n == 2:
		i, err4 := strconv.Atoi(flag.Args()[1])
		if err4 != nil {
			printUsgaes()
			break
		}
		err = cloudlist.CloudDoneTask(i)
	case a == "undo" && n == 2:
		i, err5 := strconv.Atoi(flag.Args()[1])
		if err5 != nil {
			printUsgaes()
			break
		}
		err = cloudlist.CloudUndoTask(i)
	case a == "clean" && n == 1:
		err = cloudlist.CloudCleanTask()
	case a == "clear" && n == 1:
		err = cloudlist.CloudClearTask()

	case a == "pull" && n == 1:
		_, _ = cloudlist.CloudPullAll()
	case a == "pull" && n == 2:
		i, err6 := strconv.Atoi(flag.Args()[1])
		if err6 != nil {
			printUsgaes()
			break
		}
		_, _ = cloudlist.CloudPullOne(i)
	case a == "push" && n == 1:
		_ = cloudlist.CloudPushAll()
	case a == "push" && n == 2:
		i, err7 := strconv.Atoi(flag.Args()[1])
		if err7 != nil {
			printUsgaes()
			break
		}
		_ = cloudlist.CloudPushOne(i)
	case a == "signup" && n == 1:
		username, password, retypepassword := utils.CredentialsRetype()
		if password == retypepassword {
			err = cloudlist.Signup(username, password)
		} else {
			err = errors.New("Mismatch")
		}
	case a == "login" && n == 1:
		username, password := utils.Credentials()
		err = cloudlist.Login(username, password)

	case a == "logout" && n == 1:
		err = cloudlist.Logout()

	case a == "user" && n == 1:
		err = cloudlist.ShowUserConfig()

	case a == "staying-up" && n == 1:
		fmt.Print("Happy New Year.")

	default:
		printUsgaes()
		err = nil
		os.Exit(0)
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
