package main

import (
	"fmt"
	"time"
)

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	start := time.Now()

	result := merge(
		sig(10*time.Second),
		sig(20*time.Second),
		sig(11*time.Second),
		sig(9*time.Second),
		sig(6*time.Second),
		sig(33*time.Second),
		sig(5*time.Second),
		sig(50*time.Second),
	)
	<-result
	fmt.Printf("elapsed time: %v", time.Since(start))
}

func merge(channels ...<-chan interface{}) <-chan interface{} {
	result := make(chan interface{})
	go func() {
		for {
			for i := range channels {
				select {
				case <-channels[i]:
					close(result)
					return
				default:
					continue
				}
			}
		}
	}()

	return result
}
