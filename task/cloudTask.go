package task

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jack-zh/ztodo/utils"
	"io"
	"os"
	// "strings"
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

func (l *CloudTasks) CloudUpdateTask(n int, upstr string) error {
	// tasks, err := l.CloudGet()
	// if err != nil {
	// 	return err
	// }
	// if n >= len(tasks) || n < 0 {
	// 	return errors.New("index out of range")
	// }
	// f, err := os.Create(l.Filename)
	// if err != nil {
	// 	return err
	// }
	// defer f.Close()
	// for i, t := range tasks {
	// 	if i == n {
	// 		if strings.HasPrefix(t, "1") {
	// 			t = strings.Replace(t, "1", upstr, 1)
	// 		}
	// 		if strings.HasPrefix(t, "2") {
	// 			t = strings.Replace(t, "2", upstr, 1)
	// 		}
	// 		if strings.HasPrefix(t, "0") {
	// 			t = strings.Replace(t, "0", upstr, 1)
	// 		}
	// 		_, err = fmt.Fprintln(f, t)
	// 	} else {
	// 		_, err = fmt.Fprintln(f, t)
	// 	}
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	return nil
}

func (l *CloudTasks) CloudGetTask(n int) (string, error) {
	// tasks, err := l.CloudGet()
	// if err != nil {
	// 	return "", err
	// }
	// if n >= len(tasks) || n < 0 {
	// 	return "", errors.New("index out of range")
	// }
	// return tasks[n], nil
	return "nil", nil
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

func (l *CloudTasks) CloudRemoveTask(n int) error {
	// tasks, err := l.CloudGet()
	// if n >= len(tasks) || n < 0 {
	// 	return errors.New("index out of range")
	// }
	// if err != nil {
	// 	return err
	// }
	// f, err := os.Create(l.Filename)
	// if err != nil {
	// 	return err
	// }
	// defer f.Close()
	// for i, t := range tasks {
	// 	if i == n {
	// 		continue
	// 	}
	// 	_, err = fmt.Fprintln(f, t)
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	return nil
}

func (l *CloudTasks) CloudDoneTask(n int) error {
	return l.CloudUpdateTask(n, "2")
}

func (l *CloudTasks) CloudDoingTask(n int) error {
	return l.CloudUpdateTask(n, "1")
}

func (l *CloudTasks) CloudUndoTask(n int) error {
	return l.CloudUpdateTask(n, "0")
}

func (l *CloudTasks) CloudCleanTask() error {
	// tasks, err := l.CloudGet()
	// if err != nil {
	// 	return err
	// }
	// for i, t := range tasks {
	// 	if strings.HasPrefix(t, "2") {
	// 		err = l.CloudRemoveTask(i)
	// 		if err != nil {
	// 			return err
	// 		}
	// 	}
	// }
	return nil
}

func (l *CloudTasks) CloudClearTask() error {
	// tasks, err := l.CloudGet()
	// if err != nil {
	// 	return err
	// }
	// for i, _ := range tasks {
	// 	err = l.CloudRemoveTask(i)
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	return nil
}
