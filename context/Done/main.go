package main

import (
	"context"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// キャンセルされるまでnumをひたすら送信し続けるチャネルを生成
func generator(ctx context.Context, num int) <-chan int {
	out := make(chan int)
	go func() {
		defer wg.Done()

	LOOP:
		for {
			select {
			case <-ctx.Done():
				break LOOP
			case out <- num: // キャンセルされてなければnumを送信
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
