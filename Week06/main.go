package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var RejectRequestErr = errors.New("触发限流")
type DataItem struct {
	success int
	fail int
	timeout int
	reject int
}

type Window struct {
	tail int
	head int
	size int
	max int
	preTimestamp int64
	interval int
	data []*DataItem
	m sync.Mutex
	sum DataItem
}

func GetNowMs() int64 {
	return time.Now().UnixNano()/ 1e6
}

func (w *Window) GetCount() DataItem {
	return w.sum
}

func (w *Window) getCap() int {
	return 1 + (w.tail + w.size - w.head) % w.size
}

func (w *Window) IncreSuccess() error {
	// 检查是否超过阈值
	if w.sum.success + 1 > w.max {
		return RejectRequestErr
	}
	d := w.getCurrent()
	d.success += 1
	w.sum.success += 1
	return nil
}


func (w *Window) Flush() error {
	diff := int(GetNowMs() - w.preTimestamp)
	// 应该是第几个bucket
	bucketNum := diff / w.interval + 1
	// 如果还没满,加diff个
	if w.size >= bucketNum {
		for i := 0; i < (bucketNum - w.getCap()); i++ {
			w.tail = (w.tail + 1) % w.size
			w.data[w.tail] = &DataItem{}
		}
		return w.IncreSuccess()
	}

	// 如果满了，就开始踢
	// 干掉total - w.size
	newHead := (w.head + bucketNum - w.size) % w.size
	for {
		if w.head == newHead {
			break
		}
		w.sum.success -= w.data[w.head].success
		w.data[w.head] = nil
		// ...fail、timeout、reject
		w.head = (w.head + 1) % w.size
	}
	w.preTimestamp = w.preTimestamp + int64(bucketNum - w.size) * int64(w.interval)

	// 把后面都补充上
	for i := 0; i < w.size - w.getCap(); i++ {
		w.tail = (w.tail + 1) % w.size
		w.data[w.tail] = &DataItem{}
	}
	return w.IncreSuccess()
}

func (w *Window) getCurrent() *DataItem {
	if w.data[w.tail] == nil {
		w.data[w.tail] = &DataItem{}
	}
	return w.data[w.tail]
}


func (w *Window) Add() error {
	// 判断是否有时钟回调

	w.m.Lock()
	defer w.m.Unlock()
	return w.Flush()
}

func NewWindow(size int, max int) *Window{
	return &Window{
		tail: 0,
		head: 0,
		size: size,
		max: max,
		interval: 1 * 1000,
		preTimestamp: GetNowMs(),
		data: make([]*DataItem, size),
		sum: DataItem{},
	}

}


func main() {
	w := NewWindow(5, 200)
	for j :=0; j <= 20; j++ {
		j := j
		time.Sleep(time.Duration(rand.Intn(1000) * int(time.Millisecond)))
		go func() {
			for i := 0; i <= 10; i++ {
				time.Sleep(time.Duration(rand.Intn(100) * int(time.Millisecond)))
				err := w.Add()
				if err != nil {
					fmt.Println(err, j, i)
				}
			}
		}()
	}

	time.Sleep(12 * time.Second)
	fmt.Println(w.GetCount().success, w.data)
	for _, d := range w.data {
		fmt.Println(d)
	}
}

