package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter int64
	wg      sync.WaitGroup
	mutex   sync.Mutex
)

func main() {
	wg.Add(3)
	runtime.GOMAXPROCS(1)
	go incCounter(1)
	go incCounter(2)
	go incCounter(3)

	wg.Wait()
	fmt.Printf("finish:%d\n", counter)
}

func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		//第一个参数值必须是一个指针类型的值,因为该函数需要获得被操作值在内存中的存放位置,以便施加特殊的CPU指令
		//结束时会返回原子操作后的新值
		counter = atomic.AddInt64(&counter, 1)
	}
}
