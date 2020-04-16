package task

import "github.com/robfig/cron"

// Start 启动任务
func Start()  {
	task := cron.New()

	// TODO 添加定时任务

	task.Run()
}