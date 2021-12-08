/*
 * Date: 2021/12/08
 * File: pool_test.go
 */

// Package gopool TODO package function desc
package gopool

import (
	"fmt"
	"testing"
	"time"
)

func TestNewGoroutinePool(t *testing.T) {
	gpool := NewGoroutinePool(0, 0)
	work := &Worker{
		Action: func() {
			for i := 0; i < 10; i++ {
				fmt.Println("hello,world")
				time.Sleep(500 * time.Millisecond)
			}
		},
	}
	for i := 0; i < 10; i++ {
		gpool.Submit(work)
	}
	gpool.AwaitTermination()
}
