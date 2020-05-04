# Golang_Context
Golang_Context
Context可以讓各goroutine間通訊
## context.Background()會返回一個ontext

type Context interface {
	// 獲取設置好的截止時間 ; 第二個bool返回值表示有沒有設置截止時間
	Deadline() (deadline time.Time, ok bool)

	// 返回一個 readonly channel, 如果該channel可以被讀取,
    //表示parent context 發起了cancel請求, 就能透過Done方法收到訊號後, 作結束操作.
	Done() <-chan struct{}

	// 返回取消的錯誤原因, 為什麼context被取消
	Err() error

	// 讓goroutine共享資料, 透過獲得該Context上綁定的值, 是一組KV pair, 是thread safe的;
	// 不存在則返回nil
	Value(key interface{}) interface{}
}


## ctx, cancel := context.WithCancel(context.Background())
// func WithCancel(parent Context) (ctx Context, ## cancel CancelFunc)
cancel 會獲得一個可以停止context的func