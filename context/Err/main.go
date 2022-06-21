package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// キャンセルされるまでnumをひたすら送信し続けるチャネルを生成
func generator(ctx context.Context, num int) <-chan int {
	out := make(chan int)
	go func() {
		defer wg.Done()

		for {

			// Err()メソッドでcontextがどういう状態でDone()されたかをキャッチする
			select {
			case <-ctx.Done():
				if err := ctx.Err(); errors.Is(err, context.Canceled) {
					// キャンセルされていた場合
					fmt.Println("canceled")
				} else if errors.Is(err, context.DeadlineExceeded) {
					// タイムアウトだった場合
					fmt.Println("deadline")
				}
			}
		}

		close(out)
		fmt.Println("generator closed")
	}()
	return out
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	gen := generator(ctx, 1)

	wg.Add(1)

	for i := 0; i < 5; i++ {
		fmt.Println(<-gen)
	}
	cancel()

	wg.Wait()
}
