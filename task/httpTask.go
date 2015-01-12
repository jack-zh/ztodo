package task

import (
	"errors"
	"fmt"
	"github.com/jack-zh/ztodo/utils"
)

func (l *CloudTasks) CloudPullAll() (tasks CloudTasks, err error) {
	l.CloudGetUserConfigByFile()
	fmt.Println(l)
	return tasks, err
}

func (l *CloudTasks) CloudPullOne(num int) (task CloudTask, err error) {
	l.CloudGetUserConfigByFile()
	fmt.Println(l)
	return task, nil
}

func (l *CloudTasks) CloudPushAll() error {

	return nil
}

func (l *CloudTasks) CloudPushOne(num int) error {

	return nil
}

func (l *CloudTasks) Signup(username string, password string) error {
	l.CloudGetUserConfigByFile()
	fmt.Println(l)
	if l.UserConfig.Usertoken != "" {
		return errors.New("user has login. please logout first")
	} else {
		l.UserConfig.Usertoken, _ = utils.GenUUID()
		l.UserConfig.Pushtime = "2006-01-02 15:04:05"
		l.UserConfig.Pushtoken, _ = utils.GenUUID()
		l.UserConfig.Username = username
		l.UserConfig.Password = password
		return l.CloudSaveUserConfigToFile()
	}
	return nil
}

func (l *CloudTasks) ShowUserConfig() error {
	l.CloudGetUserConfigByFile()
	userconfig := l.UserConfig
	if userconfig.Usertoken != "" {
		fmt.Printf("%13s:  %s\n%13s:  %s\n%13s:  %s\n%13s:  %s\n%13s:  %s\n",
			"Username", userconfig.Username,
			"Pushtime", userconfig.Pushtime,
			"Password", userconfig.Password,
			"Usertoken", userconfig.Usertoken,
			"Pushtoken", userconfig.Pushtoken)
	} else {
		fmt.Println("Please login or signup first")
	}
	return nil
}

func (l *CloudTasks) Login(username string, password string) error {
	l.CloudGetUserConfigByFile()
	fmt.Println(l)
	if l.UserConfig.Usertoken != "" {
		return errors.New("user has login. please logout first or showlogin")
	} else {
		fmt.Println(username)
		fmt.Println(password)
	}
	return nil
}

func (l *CloudTasks) Logout() error {
	return l.CloudSaveUserConfigToFile()
}
