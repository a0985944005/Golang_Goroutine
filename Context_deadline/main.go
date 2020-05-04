package main

import (
	"context"
	"log"
	"os"
	"time"
)

var logger *log.Logger

func do(ctx context.Context) {
	if deadline, ok := ctx.Deadline(); ok == true { //檢查是否存在方法
		logger.Println("deadline: ", deadline)
	}
	for {
		select {
		case <-ctx.Done():
			logger.Println("deadline時間到")
			logger.Println(ctx.Err()) //context deadline exceeded
			return
		default:
			logger.Println("繼續....")
		}
	}

}
func main() {
	logger = log.New(os.Stdout, "", log.Ltime)

	deadline := time.Now().Add(3 * time.Second)
	// 現在時間的2秒後的時間就是deadline
	ctx, cancel := context.WithDeadline(context.Background(), deadline)

	defer cancel()
	logger.Println("開始執行任務")
	go do(ctx)

	time.Sleep(4 * time.Second) //比deadline多一秒
	logger.Println("任務結束囉")
}
