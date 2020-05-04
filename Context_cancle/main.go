package main

import (
	"context"
	"log"
	"os"
	"time"
)

var logger *log.Logger
var key string = "name"

func main() {
	logger = log.New(os.Stdout, "", log.Ltime)

	//建立一個cancel context
	// func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
	ctx, cancel := context.WithCancel(context.Background())

	//建立數個withValue context,繼承於ctx
	valueCtx1 := context.WithValue(ctx, key, "哥吉拉")
	valueCtx2 := context.WithValue(ctx, key, "航海王")
	go watch(valueCtx1)
	go watch(valueCtx2)

	//等待10秒後停止任務
	time.Sleep(10 * time.Second)
	//發布任務停止
	cancel()
	logger.Println("哥吉拉,航海王都停止了!")

	// 確保工作結束
	time.Sleep(1 * time.Second)
}

func watch(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			//接收到關閉訊號
			logger.Println("任務" + ctx.Value(key).(string) + ":停止~")
			return
		default:
			//工作中
			logger.Println("任務" + ctx.Value(key).(string) + ":運作~")
		}
	}
}
