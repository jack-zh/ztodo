package task

import ()

func (l *CloudTasks) CloudPullAll() (tasks CloudTasks, err error) {

	return tasks, nil
}

func (l *CloudTasks) CloudPullOne(num int) (task CloudTask, err error) {

	return task, nil
}

func (l *CloudTasks) CloudPushAll() error {

	return nil
}

func (l *CloudTasks) CloudPushOne(num int) error {

	return nil
}
