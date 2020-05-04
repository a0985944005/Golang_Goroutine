package main

import (
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

// 不能讓發布者或是訂閱者來關閉資料通道, 且不能讓發布者這邊來關閉額外的訊息通道來通知其他所有角色退出.
// 引入主持人這角色在這情境下, 來關閉訊息通道
// 主持人目的是為了可以知道是誰關閉了channel

func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)
	//先設定亂數的最大範圍
	const Max = 100000
	//訂閱者數量上限
	const NumSubscribers = 10
	//發布者數量上限
	const NumPublishers = 1000

	wgSubscribers := sync.WaitGroup{}
	wgSubscribers.Add(NumSubscribers)

	//資料通道
	dataCh := make(chan int)
	//停止通道,主持人出來後會關閉此通道
	stopCh := make(chan struct{})
	//一個長度為1的通道，住要視主持人關閉通道時要看主持人是誰
	toStop := make(chan string, 1)
	//主持人(為某個subscriber或publisher)
	var stoppedBy string

	//主持人，一開始因為頻道內沒有東西會先block自己
	//直到主持人取值成功，就會關閉通道
	go func() {
		stoppedBy = <-toStop
		close(stopCh)
	}()

	//創建Publishers
	for i := 0; i < NumPublishers; i++ {
		go func(id string) {
			for {
				value := rand.Intn(Max)
				if value == 0 {
					//達成條件的某個創建Publishers -> 主持人
					select {
					case toStop <- "publishers#" + id:
					default:
					}
					return
				}
				//嘗試從停止通道中取值，或者不阻塞往下繼續執行
				select {
				case <-stopCh:
					return
				default:
				}

				//嘗試停止通道中取值，或者發送資料到資料通道
				select {
				case <-stopCh:
					return
				case dataCh <- value:
				}

			}
		}(strconv.Itoa(i)) //轉成字串
	}

	//創建Subscribers
	for i := 0; i < NumSubscribers; i++ {
		go func(id string) {
			defer wgSubscribers.Done()
			for {
				//嘗試從停止通道中取值，或者不阻塞往下繼續執行
				select {
				case <-stopCh:
					return
				default:
				}
				//嘗試從停止通道中取值，或者資料通道取值
				select {
				case <-stopCh:
					return
				case value := <-dataCh:
					if value == Max-1 {
						select {
						case toStop <- "receiver#" + id:
						default:
						}
						return
					}
					log.Println(value)
				}
			}
		}(strconv.Itoa(i))
	}
	wgSubscribers.Wait()
	log.Println("stopped by", stoppedBy)
}
