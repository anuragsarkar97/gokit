package concurrent

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestNewFutureTask(t *testing.T) {
	w := NewWorkerPool(15)
	for i := 0; i < 10; i++ {
		go func(i2 int) {
			w.AddTask(NewFutureTask(func() {
				fmt.Println("start " + strconv.Itoa(i2))
				fmt.Println("      done " + strconv.Itoa(i2))
			}))
		}(i)
	}
	time.Sleep(1 * time.Second)
	go w.Run()
	w.Wait()

}
