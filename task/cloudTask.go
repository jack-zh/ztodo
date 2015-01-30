package task

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/jack-zh/ztodo/utils"
	"github.com/jack-zh/ztodo/zterminal/color"
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

type UserConfig struct {
	Usertoken string `json:"usertoken"`
	Pushtime  string `json:"pushtime"`
	Pushtoken string `json:"pushtoken"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type CloudTasks struct {
	WorkFilename       string
	WorkTasks          []CloudTask
	BackupFilename     string
	BackupTasks        []CloudTask
	UserConfigFilename string
	UserConfig         UserConfig
}

func CloudNewList(workfilename string, backfilename string, userconfig_filename string) *CloudTasks {
	return &CloudTasks{workfilename, nil, backfilename, nil, userconfig_filename, UserConfig{}}
}

func (l *CloudTasks) CloudGetUserConfigByFile() error {
	fd, err := os.Open(l.UserConfigFilename)
	if err != nil {
		return nil
	}
	str := ""
	br := bufio.NewReader(fd)
	for {
		linestr, _, readerr := br.ReadLine()
		if readerr == io.EOF {
			break
		}
		if readerr != nil {
			return readerr
		}
		str += string(linestr)
	}

	userconfig := &l.UserConfig
	jsonerr := json.Unmarshal([]byte(str), userconfig)
	if jsonerr != nil {
		return jsonerr
	}
	return nil
}

func (l *CloudTasks) CloudSaveUserConfigToFile() error {
	fd, err := os.Create(l.UserConfigFilename)
	if err != nil {
		return err
	}
	defer fd.Close()

	jsonstr, err := json.Marshal(l.UserConfig)
	if err != nil {
		fmt.Println("error:", err)
	}
	fd.WriteString(string(jsonstr))
	return nil
}

func (l *CloudTasks) CloudUpdateTaskStatus(n int, upstr string) error {
	l.CloudGetAllWorkTaskByFile()
	if n > 0 && n <= len(l.WorkTasks) {
		l.WorkTasks[n-1].Status = upstr
		time_str := time.Now().Format("2006-01-02 15:04:05")
		l.WorkTasks[n-1].Updatetime = time_str
		if upstr == "Future" {
			l.WorkTasks[n-1].Doingtime = "2006-01-02 15:04:05"
			l.WorkTasks[n-1].Donetime = "2006-01-02 15:04:05"
		} else if upstr == "Done" {
			l.WorkTasks[n-1].Donetime = time_str
		} else {
			l.WorkTasks[n-1].Doingtime = time_str
		}
	} else {
		return errors.New("index out of range")
	}
	return l.CloudTaskToFile()
}

func (l *CloudTasks) CloudGetAllBackupTaskByFile() error {
	// backup task get by json str
	backupfd, backuperropenfile := os.Open(l.BackupFilename)
	backup_tasks_string := ""
	if backuperropenfile == nil {
		backupbr := bufio.NewReader(backupfd)
		for {
			backupstr, _, backuperr := backupbr.ReadLine()
			if backuperr == io.EOF {
				break
			}
			if backuperr != nil {
				return backuperr
			}
			backup_tasks_string += string(backupstr)
		}

		backuptasks := &l.BackupTasks
		backupjsonunmarshalerr := json.Unmarshal([]byte(backup_tasks_string), backuptasks)
		if backupjsonunmarshalerr != nil {
			return backupjsonunmarshalerr
		}
	}
	return nil
}

func (l *CloudTasks) CloudGetAllWorkTaskByFile() error {

	// work task get by json str
	workfd, workerropenfile := os.Open(l.WorkFilename)
	work_tasks_string := ""
	if workerropenfile == nil {

		work_br := bufio.NewReader(workfd)
		for {
			workstr, _, workerr := work_br.ReadLine()
			if workerr == io.EOF {
				break
			}
			if workerr != nil {
				return workerr
			}
			work_tasks_string += string(workstr)
		}
		worktasks := &l.WorkTasks
		workjsonunmarshalerr := json.Unmarshal([]byte(work_tasks_string), worktasks)
		if workjsonunmarshalerr != nil {

			return workjsonunmarshalerr
		}
	}
	return nil
}

func (l *CloudTasks) CloudTaskToFile() error {
	// work tasks json to file
	workfd, err := os.Create(l.WorkFilename)
	if err != nil {
		return err
	}
	defer workfd.Close()

	workjsonstr, err := json.Marshal(l.WorkTasks)
	if err != nil {
		fmt.Println("error:", err)
	}
	workfd.WriteString(string(workjsonstr))

	// backup task json to file
	backfd, err := os.Create(l.BackupFilename)
	if err != nil {
		return err
	}
	defer backfd.Close()

	backupjsonstr, err := json.Marshal(l.BackupTasks)
	if err != nil {
		fmt.Println("error:", err)
	}
	backfd.WriteString(string(backupjsonstr))

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
	l.CloudGetAllWorkTaskByFile()
	l.WorkTasks = append(l.WorkTasks, task)
	return l.CloudTaskToFile()
}

func (l *CloudTasks) CloudRmOneTask(n int) {
	task := l.WorkTasks[n]
	task.Updatetime = time.Now().Format("2006-01-02 15:04:05")
	l.BackupTasks = append(l.BackupTasks, task)
	l.WorkTasks = append(l.WorkTasks[:n], l.WorkTasks[n+1:]...)
}

func (l *CloudTasks) CloudRemoveTask(n int) error {
	l.CloudGetAllWorkTaskByFile()
	l.CloudGetAllBackupTaskByFile()
	if n <= len(l.WorkTasks) && n > 0 {
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
	l.CloudGetAllWorkTaskByFile()
	l.CloudGetAllBackupTaskByFile()
	for n := len(l.WorkTasks) - 1; n >= 0; n-- {
		if l.WorkTasks[n].Status == "Done" {
			l.CloudRmOneTask(n)
		}
	}
	return l.CloudTaskToFile()
}

func (l *CloudTasks) CloudClearTask() error {
	l.CloudGetAllWorkTaskByFile()
	l.CloudGetAllBackupTaskByFile()
	for n := len(l.WorkTasks) - 1; n >= 0; n-- {
		l.CloudRmOneTask(n)
	}
	return l.CloudTaskToFile()
}

func (l *CloudTasks) CloudTasksPrint(i int) {
	if i == -1 {
		for ti := 0; ti < len(l.WorkTasks); ti++ {
			task := l.WorkTasks[ti]
			printTask(task, ti+1)
		}
	} else {
		if i <= len(l.WorkTasks) && i > 0 {
			task := l.WorkTasks[i-1]
			printTask(task, i)
		}
	}
}

func (l *CloudTasks) CloudTasksPrintVerbose(i int) {
	if i == -1 {
		for ti := 0; ti < len(l.WorkTasks); ti++ {
			task := l.WorkTasks[ti]
			printTaskVerbose(task, ti+1)
		}

	} else {
		if i <= len(l.WorkTasks) && i > 0 {
			task := l.WorkTasks[i-1]
			printTaskVerbose(task, i)
		}
	}
}

func printTask(task CloudTask, i int) {
	fmt.Printf("%-3s: [%-6s] [%s] %s\n", strconv.Itoa(i), task.Status, task.Updatetime, task.Task)
	c := color.New(color.FgCyan).Add(color.Underline)
	c.Printf("Prints cyan text with an underline.")
}

func printTaskVerbose(task CloudTask, i int) {
	if task.Doingtime != "2006-01-02 15:04:05" && task.Donetime != "2006-01-02 15:04:05" {
		fmt.Printf("%13s:  %s\n%13s:  %s\n%13s:  %s\n%13s:  %s\n%13s:  %s\n%13s:  %s\n%13s:  %s\n%13s:  %s\n----------------------------------------\n",
			"token", task.Token,
			"num", strconv.Itoa(i),
			"task", task.Task,
			"status", task.Status,
			"create time", task.Createtime,
			"doing time", task.Doingtime,
			"done time", task.Donetime,
			"update time", task.Updatetime)
	} else if task.Doingtime == "2006-01-02 15:04:05" && task.Donetime != "2006-01-02 15:04:05" {
		fmt.Printf("%13s:  %s\n%13s:  %s\n%13s:  %s\n%13s:  %s\n%13s:  %s\n%13s:  %s\n%13s:  %s\n----------------------------------------\n",
			"token", task.Token,
			"num", strconv.Itoa(i),
			"task", task.Task,
			"status", task.Status,
			"create time", task.Createtime,
			"done time", task.Donetime,
			"update time", task.Updatetime)
	} else if task.Doingtime != "2006-01-02 15:04:05" && task.Donetime == "2006-01-02 15:04:05" {
		fmt.Printf("%13s:  %s\n%13s:  %s\n%13s:  %s\n%13s:  %s\n%13s:  %s\n%13s:  %s\n%13s:  %s\n----------------------------------------\n",
			"token", task.Token,
			"num", strconv.Itoa(i),
			"task", task.Task,
			"status", task.Status,
			"create time", task.Createtime,
			"doing time", task.Doingtime,
			"update time", task.Updatetime)
	} else {
		fmt.Printf("%13s:  %s\n%13s:  %s\n%13s:  %s\n%13s:  %s\n%13s:  %s\n%13s:  %s\n----------------------------------------\n",
			"token", task.Token,
			"num", strconv.Itoa(i),
			"task", task.Task,
			"status", task.Status,
			"create time", task.Createtime,
			"update time", task.Updatetime)
	}
}
