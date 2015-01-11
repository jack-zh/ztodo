package task

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jack-zh/ztodo/utils"
	"io"
	"os"
	"strconv"
	"time"
)

type CloudTask struct {
	Task       string `json:"task"`
	Token      string `json:"token"`
	Createtime string `json:"createtime"`
	Doingtime  string `json:"doingtime"`
	Updatetime string `json:"updatetime"`
	Donetime   string `json:"donetime"`
	Status     string `json:"status"`
}

type CloudTasks struct {
	Filename string
	Tasks    []CloudTask
}

func CloudNewList(filename string) *CloudTasks {
	return &CloudTasks{filename, nil}
}

func (l *CloudTasks) CloudUpdateTaskStatus(n int, upstr string) error {
	l.CloudGetAllTaskByFile()
	if n > 0 && n <= len(l.Tasks) {
		l.Tasks[n-1].Status = upstr
	} else {
		return errors.New("index out of range")
	}
	return l.CloudTaskToFile()
}

func (l *CloudTasks) CloudGetAllTaskByFile() error {
	f, err := os.Open(l.Filename)
	if err != nil {
		if os.IsNotExist(err) {
			return errors.New("no file")
		}
		return errors.New("open file error")
	}

	var tasks_string string
	br := bufio.NewReader(f)
	for {
		t, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		tasks_string += string(t)
	}

	tasks := &l.Tasks
	return json.Unmarshal([]byte(tasks_string), tasks)
}

func (l *CloudTasks) CloudTaskToFile() error {
	f, err := os.Create(l.Filename)
	if err != nil {
		return err
	}
	defer f.Close()

	b, err := json.Marshal(l.Tasks)
	if err != nil {
		fmt.Println("error:", err)
	}
	f.WriteString(string(b))
	return nil
}

func (l *CloudTasks) CloudAddTask(s string) error {
	create_time_str := time.Now().Format("2006-01-02 15:04:05")
	doing_time_str := "2006-01-02 15:04:05"
	done_time_str := "2006-01-02 15:04:05"
	status := "Future"
	task_str := s
	token, _ := utils.GenUUID()
	task := CloudTask{
		Task:       task_str,
		Token:      token,
		Createtime: create_time_str,
		Doingtime:  doing_time_str,
		Donetime:   done_time_str,
		Status:     status,
		Updatetime: create_time_str,
	}
	l.CloudGetAllTaskByFile()
	l.Tasks = append(l.Tasks, task)
	return l.CloudTaskToFile()
}

func (l *CloudTasks) CloudRmOneTask(n int) {
	l.Tasks = append(l.Tasks[:n], l.Tasks[n+1:]...)
}

func (l *CloudTasks) CloudRemoveTask(n int) error {
	l.CloudGetAllTaskByFile()
	if n <= len(l.Tasks) && n > 0 {
		l.CloudRmOneTask(n - 1)
		return l.CloudTaskToFile()
	} else {
		return errors.New("index out of range")
	}
}

func (l *CloudTasks) CloudDoneTask(n int) error {
	return l.CloudUpdateTaskStatus(n, "Done")
}

func (l *CloudTasks) CloudDoingTask(n int) error {
	return l.CloudUpdateTaskStatus(n, "Doing")
}

func (l *CloudTasks) CloudUndoTask(n int) error {
	return l.CloudUpdateTaskStatus(n, "Future")
}

func (l *CloudTasks) CloudCleanTask() error {
	for n := 0; n < len(l.Tasks); n++ {
		if l.Tasks[n].Status == "Done" {
			l.CloudRmOneTask(n)
		}
	}
	return l.CloudTaskToFile()
}

func (l *CloudTasks) CloudClearTask() error {
	l.Tasks = nil
	return l.CloudTaskToFile()
}

func (l *CloudTasks) CloudTasksPrint(i int) {
	if i == -1 {
		for ti := 0; ti < len(l.Tasks); ti++ {
			task := l.Tasks[ti]
			fmt.Printf("%-3s: [%-6s] [%s] %s\n", strconv.Itoa(ti+1), task.Status, task.Updatetime, task.Task)
		}
	} else {
		if i <= len(l.Tasks) && i > 0 {
			task := l.Tasks[i-1]
			fmt.Printf("%-3s: [%-6s] [%s] %s\n", strconv.Itoa(i), task.Status, task.Updatetime, task.Task)
		}
	}
}

func (l *CloudTasks) CloudTasksPrintVerbose(i int) {
	if i == -1 {
		for ti := 0; ti < len(l.Tasks); ti++ {
			task := l.Tasks[ti]
			fmt.Printf("%13s:  %s\n%13s:  %s\n%13s:  %s\n%13s:  %s\n%13s:  %s\n%13s:  %s\n%13s:  %s\n%13s:  %s\n----------------------------------------\n",
				"token", task.Token,
				"num", strconv.Itoa(ti+1),
				"task", task.Task,
				"status", task.Status,
				"create time", task.Createtime,
				"doing time", task.Doingtime,
				"done time", task.Donetime,
				"update time", task.Updatetime)
		}
	} else {
		if i <= len(l.Tasks) && i > 0 {
			task := l.Tasks[i-1]
			fmt.Printf("%13s:  %s\n%13s:  %s\n%13s:  %s\n%13s:  %s\n%13s:  %s\n%13s:  %s\n%13s:  %s\n%13s:  %s\n----------------------------------------\n",
				"token", task.Token,
				"num", strconv.Itoa(i),
				"task", task.Task,
				"status", task.Status,
				"create time", task.Createtime,
				"doing time", task.Doingtime,
				"done time", task.Donetime,
				"update time", task.Updatetime)
		}
	}
}
