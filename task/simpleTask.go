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

type SimpleList struct {
	filename string
}

func SimpleNewList(filename string) *SimpleList {
	return &SimpleList{filename}
}

func (l *SimpleList) SimpleUpdateTask(n int, upstr string) error {
	tasks, err := l.SimpleGet()
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

func (l *SimpleList) SimpleAddTask(s string) error {
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

func (l *SimpleList) SimpleRemoveTask(n int) error {
	tasks, err := l.SimpleGet()
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

func (l *SimpleList) SimpleGetTask(n int) (string, error) {
	tasks, err := l.SimpleGet()
	if err != nil {
		return "", err
	}
	if n >= len(tasks) || n < 0 {
		return "", errors.New("index out of range")
	}
	return tasks[n], nil
}

func (l *SimpleList) SimpleGet() ([]string, error) {
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

func (l *SimpleList) SimpleDoneTask(n int) error {
	return l.SimpleUpdateTask(n, "2")
}

func (l *SimpleList) SimpleDoingTask(n int) error {
	return l.SimpleUpdateTask(n, "1")
}

func (l *SimpleList) SimpleUndoTask(n int) error {
	return l.SimpleUpdateTask(n, "0")
}

func (l *SimpleList) SimpleCleanTask() error {
	tasks, err := l.SimpleGet()
	if err != nil {
		return err
	}
	for i, t := range tasks {
		if strings.HasPrefix(t, "2") {
			err = l.SimpleRemoveTask(i)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (l *SimpleList) SimpleClearTask() error {
	tasks, err := l.SimpleGet()
	if err != nil {
		return err
	}
	for i, _ := range tasks {
		err = l.SimpleRemoveTask(i)
		if err != nil {
			return err
		}
	}
	return nil
}
