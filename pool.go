/*
 * Date: 2021/12/08
 * File: pool.go
 */

package gopool

import (
	"sync"
)

// 默认最大任务数
var defaultMaxWorkerCount int32 = 100

// 默认最大协程数
var defaultMaxGoroutineCount int32 = 10

// 协程池结构定义
type GoroutinePool struct {
	maxGoroutineCount int32         // 最大协程数
	maxWorkerCount    int32         // 最大任务数
	workerQueue       chan Workable // 任务队列
	waitGroup         sync.WaitGroup
}

/**
 * NewGoroutinePool
 * 实例化一个协程池
 */
func NewGoroutinePool(maxGoroutineCount int32, maxWorkerCount int32) *GoroutinePool {
	if maxGoroutineCount <= 0 {
		maxGoroutineCount = defaultMaxGoroutineCount
	}
	if maxWorkerCount <= 0 {
		maxWorkerCount = defaultMaxWorkerCount
	}
	pool := &GoroutinePool{
		maxGoroutineCount: maxGoroutineCount,
		workerQueue:       make(chan Workable, maxWorkerCount),
	}
	pool.start()
	return pool
}

/**
 * 提交并执行一个任务
 */
func (pool *GoroutinePool) Submit(worker Workable) {
	pool.waitGroup.Add(1)
	pool.workerQueue <- worker
}

/**
 * 等待所有的任务执行完成
 */
func (pool *GoroutinePool) AwaitTermination() {
	pool.waitGroup.Wait()
}

/**
 * 协程池开始执行
 */
func (pool *GoroutinePool) start() {
	for i := 0; int32(i) < pool.maxGoroutineCount; i++ {
		go pool.doWork()
	}
}

/**
 * 开始处理所有的任务
 */
func (pool *GoroutinePool) doWork() {
	for worker := range pool.workerQueue {
		worker.Work()
		pool.waitGroup.Done()
	}
}

/**
 * 关系线程池
 */
func (pool *GoroutinePool) Close() {
	close(pool.workerQueue)
}
