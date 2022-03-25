package channel

import (
	"fmt"
	"time"
)

func RunUnbufferedChannelExample() {
	done := make(chan struct{})
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
			time.Sleep(1 * time.Second)
		}
		done <- struct{}{}
	}()
	<-done
	fmt.Println("finished!")
}
