package main

import (
	"context"
	"log"
	"os"
	"time"
)

var logger *log.Logger

func do2(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			logger.Println(ctx.Err())
			return
		default:
			logger.Println("航海王")
			time.Sleep(1 * time.Second)
		}
	}
}

func do1(ctx context.Context) {
	select {
	case <-ctx.Done():
		logger.Println(ctx.Err())
		return
	default:
		logger.Println("哥吉拉")
		time.Sleep(1 * time.Second)

	}
}

func main() {
	logger = log.New(os.Stdout, "", log.Ltime)
	timeout := 3 * time.Second
	// 建立一個timeout context,  3秒後沒返回就發出超時
	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	defer cancel()
	logger.Println("開始選擇")
	go do1(ctx)
	go do2(ctx)

	time.Sleep(4 * time.Second) // 比timeout多一秒
}
