package main

import (
	"fmt"
	"sync"
	"time"
)

func TaskA(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	fmt.Println("task a done...")
}

func TaskB(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(3 * time.Second)
	fmt.Println("task b done...")
}

func TaskC(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(5 * time.Second)
	fmt.Println("task c done...")
}

func main() {
	var wg sync.WaitGroup

	//開啟一個 goroutine 前，在 main goroutine 中呼叫 wg.Add
	//為什麼強調一定要在 main goroutine 中呼叫 wg.Add，而不能在
	//新建的 goroutine 中呼叫呢？我的理解是如果延遲到新建的
	//goroutine 中呼叫 wg.Add 就有可能造成 wg.Wait 先執行。若是
	//如此，必定不能達到如期的效果。
	//傳遞給 goroutine 的 WaitGroup 變數必須為指標型別，因為在
	//Golang 中所有的函式引數都是值傳遞，也就是在函式體內會複製一
	//份引數的副本。如果不使用指標型別就無法引用到同一個 WaitGroup
	//變數，便也不能依賴 WaitGroup 來實現同步了。
	wg.Add(1)
	go TaskA(&wg)

	wg.Add(1)
	go TaskC(&wg)

	wg.Add(1)
	go TaskB(&wg)



	//阻塞等待直至所有其它的 goroutine 都已執行完畢
	wg.Wait()
	fmt.Println("all the tasks done...")


	// 也可以一次設定幾個goroutine去執行
	wg.Add(3)
	go TaskA(&wg)
	go TaskB(&wg)
	go TaskC(&wg)

	//阻塞等待直至所有其它的 goroutine 都已執行完畢
	wg.Wait()
	fmt.Println("all the tasks done...")
}geniusboss
