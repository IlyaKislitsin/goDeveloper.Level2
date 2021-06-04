package main

import (
	"context"
	"fmt"
	"os"
	"syscall"
	"time"
)

// Это промежуточная версия, пока не разобрался как задейсвовать контекст с таймаутом
// при поступлении сигнала...
func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())

	signals := make(chan os.Signal)

	go func() {
		ticker := 0
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			select {
			case <-ctx.Done():
			default:
				ticker++
				fmt.Println(ticker)
			}
		}
	}()

	go func() {
		defer close(signals)

		time.Sleep(5 * time.Second)
		signals <- syscall.SIGTERM
	}()

	select {
	case sygVal := <-signals:
		if sygVal == syscall.SIGTERM {
			fmt.Println("SYGTERM sygnal")
			cancelFunc()
		}
	case <-time.After(11 * time.Second):
		fmt.Println("Прорамма отработала без ошибок!")
	}

	time.Sleep(time.Second)
	fmt.Println("Конец")
}
