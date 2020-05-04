package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// 使用一個邏輯處理器給調度器用
	runtime.GOMAXPROCS(2)
	t := time.Now()
	// wg + 2 表示要等待2個goroutine完成

	wg.Add(2)

	fmt.Println("start goroutines")

	// 創建goroutine
	go printPrime("A")
	go printPrime("B")

	fmt.Println("waiting to finish")
	wg.Wait()

	fmt.Println("\nfinish Program")
	fmt.Println("耗費：" + time.Now().Sub(t).String())
}

func printPrime(prefix string) {
	defer wg.Done()

next:
	for outer := 2000; outer < 50000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next //14版以後可以不用GoTo會自斷切換goroutine
			}
		}
		fmt.Printf("%s : %d\n", prefix, outer)
	}
	fmt.Println("completed", prefix)
}
