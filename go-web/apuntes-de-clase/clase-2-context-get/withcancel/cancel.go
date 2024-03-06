package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"
)

func SimulateOperation(ctx context.Context) {
	fmt.Println("SimulateOperation started")
	for i := 0; i < 5; i++ {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("Doing something...")
		case <-ctx.Done():
			fmt.Println("Operation canceled")
			return
		}
	}
	fmt.Println("SimulateOperation finished")
}

func main() {
	ctx := context.Background()
	ctx, cancelFunc := context.WithCancel(ctx)
	go SimulateOperation(ctx)
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	cancelFunc()
	fmt.Println("End of the program")
}
