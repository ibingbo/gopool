/*
 * Date: 2021/12/08
 * File: worker.go
 */

// Package gopool TODO package function desc
package gopool

// 任务接口
type Workable interface {
	Work()
}

// 任务结构
type Worker struct {
	Action func()
}

/**
 * Work
 * 启动任务
 */
func (worker *Worker) Work() {
	worker.Action()
}
