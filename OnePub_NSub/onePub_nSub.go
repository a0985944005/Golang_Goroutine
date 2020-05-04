package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

// 一個發布者, 多個訂閱者
// 因為只有一個發布者對上channel, 所以由發布者自己決定什麼時候關閉通道
func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	// 隨機數字的最大值
	const Max = 100000
	// 訂閱者數量
	const NumSubscribers = 100

	wgSubscribers := sync.WaitGroup{}
	wgSubscribers.Add(NumSubscribers)

	// 資料通道
	dataCh := make(chan int)

	// 發布者
	go func() {
		for {
			// 當剛好出現0時
			if value := rand.Intn(Max); value == 0 {
				// 唯一的發布者可自己關閉通道
				close(dataCh)
				return
			} else {
				dataCh <- value
			}
		}
	}()

	//  訂閱者
	for i := 0; i < NumSubscribers; i++ {
		go func(i int) {
			defer wgSubscribers.Done()

			//一直從channel接收資料直到通道關閉, 且都沒資料為止
			for value := range dataCh {
				log.Printf("Subscriber%d : %d", i, value)
			}
		}(i)
	}

	wgSubscribers.Wait()
}
