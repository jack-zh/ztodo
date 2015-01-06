package task

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type CloudTask struct {
	task       string
	token      string
	createtime string
	doingtime  string
	updatetime string
	donetime   string
	status     string
}

type CloudTasks struct {
	filename string
	tasks    []CloudTask
}

func CloudNewList(filename string) *CloudTasks {
	return &CloudTasks{filename, nil}
}

func (l *CloudTasks) CloudUpdateTask(n int, upstr string) error {
	tasks, err := l.CloudGet()
	if err != nil {
		return err
	}
	if n >= len(tasks) || n < 0 {
		return errors.New("index out of range")
	}
	f, err := os.Create(l.filename)
	if err != nil {
		return err
	}
	defer f.Close()
	for i, t := range tasks {
		if i == n {
			if strings.HasPrefix(t, "1") {
				t = strings.Replace(t, "1", upstr, 1)
			}
			if strings.HasPrefix(t, "2") {
				t = strings.Replace(t, "2", upstr, 1)
			}
			if strings.HasPrefix(t, "0") {
				t = strings.Replace(t, "0", upstr, 1)
			}
			_, err = fmt.Fprintln(f, t)
		} else {
			_, err = fmt.Fprintln(f, t)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (l *CloudTasks) CloudAddTask(s string) error {
	s = strings.Join([]string{"0", time.Now().Format("[2006-01-02 15:04:05]"), s}, " ")
	var flags = os.O_WRONLY | os.O_CREATE | os.O_APPEND
	f, err := os.OpenFile(l.filename, flags, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = fmt.Fprintln(f, s)
	if err != nil {
		return err
	}
	return err
}

func (l *CloudTasks) CloudRemoveTask(n int) error {
	tasks, err := l.CloudGet()
	if n >= len(tasks) || n < 0 {
		return errors.New("index out of range")
	}
	if err != nil {
		return err
	}
	f, err := os.Create(l.filename)
	if err != nil {
		return err
	}
	defer f.Close()
	for i, t := range tasks {
		if i == n {
			continue
		}
		_, err = fmt.Fprintln(f, t)
		if err != nil {
			return err
		}
	}
	return nil
}

func (l *CloudTasks) CloudGetTask(n int) (string, error) {
	tasks, err := l.CloudGet()
	if err != nil {
		return "", err
	}
	if n >= len(tasks) || n < 0 {
		return "", errors.New("index out of range")
	}
	return tasks[n], nil
}

func (l *CloudTasks) CloudGet() ([]string, error) {
	f, err := os.Open(l.filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	var tasks []string
	br := bufio.NewReader(f)
	for {
		t, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, string(t))
	}
	return tasks, nil
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
	tasks, err := l.CloudGet()
	if err != nil {
		return err
	}
	for i, t := range tasks {
		if strings.HasPrefix(t, "2") {
			err = l.CloudRemoveTask(i)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (l *CloudTasks) CloudClearTask() error {
	tasks, err := l.CloudGet()
	if err != nil {
		return err
	}
	for i, _ := range tasks {
		err = l.CloudRemoveTask(i)
		if err != nil {
			return err
		}
	}
	return nil
}
