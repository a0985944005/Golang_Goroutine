package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	const Max = 100000
	// 發布者數量
	const NumPublishers = 1000

	wgSubscriber := sync.WaitGroup{}
	wgSubscriber.Add(1)
	// 資料通道
	dataCh := make(chan int)
	// 停止訊號通道, 發訊號給他的是訂閱者, 訂閱者因為自己不能關閉通道, 會違反原則
	// 發布者收到停止訊號後, 就會停止發布並且返回
	stopCh := make(chan struct{})

	// 創建多個發布者
	for i := 0; i < NumPublishers; i++ {
		go func() {
			for {
				// 如果只有一個select 內有從stopCh取值跟送值給dataCh這兩個case.
				// 當同時兩個條件都滿足下, 是會發生隨機挑一個case去執行的無法預估的情況.
				// 所以第一個select只會有從stopCh取值作提早返回和default case避免阻塞用.
				select {
				// 發布者對資料通道是發布者的角色
				// 但是對停止訊號通道則是訂閱者的角色
				case <-stopCh:
					return
				default:
				}

				select {
				case <-stopCh:
					return
				case dataCh <- rand.Intn(Max):
				}
			}
		}()
	}

	//  訂閱者
	go func() {
		defer wgSubscriber.Done()

		for value := range dataCh {
			if value == Max-1 {
				// 訂閱者對停止事件通道的角色則是發布的作用,
				// 所以由他負責關閉沒有違反原則, 且也只有他一位.
				close(stopCh) ///當CLOSE CHANNEL 時 訂閱的也會收到?
				return
			}
			//那另外一個datachannel裡面的值如何處理   GC會處理

			log.Println(value)
		}
	}()

	wgSubscriber.Wait()
}
