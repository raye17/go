package common

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"time"
)

// 定义一个执行者，可以执行任何任务，任务是限制完成的
// 执行者可以通过发送信号终止

type Runner struct {
	tasks     []func(int)      //表示任务
	complete  chan error       //用于通知任务完成
	timeout   <-chan time.Time //限制完成时间
	interrupt chan os.Signal   //终止信号
}

func NewRunner(tm time.Duration) *Runner {
	return &Runner{
		complete:  make(chan error),
		timeout:   time.After(tm),
		interrupt: make(chan os.Signal, 1),
	}
}
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

var (
	ErrTimeout   = errors.New("执行超时")
	ErrInterrupt = errors.New("执行被中断")
)

func (r *Runner) isInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}
func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.isInterrupt() {
			return ErrInterrupt
		}
		task(id)
	}
	return nil
}
func (r *Runner) Start() error {
	signal.Notify(r.interrupt, os.Interrupt)
	go func() {
		r.complete <- r.run()
	}()
	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout
	}

}
func CreateTask() func(int) {
	return func(id int) {
		log.Println("正在执行任务：", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
