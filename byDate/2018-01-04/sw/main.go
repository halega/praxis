package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	block := make(chan os.Signal, 1)
	signal.Notify(block, os.Interrupt)
	go func() {
		now := time.Now()
		for {
			fmt.Printf("\r%v", time.Since(now).Round(time.Second))
			<-time.Tick(1 * time.Second)
		}
	}()
	<-block
}
