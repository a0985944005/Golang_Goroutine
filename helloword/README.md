# Golang_Goroutine
wg.add(int)
int為有幾個goroutine需要等待
wg.Done()
每個須等待的goroutine執行完都要執行done()表示結束
wg.Wait()
會等待所有須等待的goroutine都執行結束
