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
		c := time.Tick(1 * time.Second)
		now := time.Now()
		for {
			fmt.Printf("\r%v", time.Since(now).Round(time.Second))
			<-c
		}
	}()
	<-block
}
