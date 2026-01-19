package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/heathcliff26/cultures-trainer/pkg/trainer"
)

var globalResources []int32 = make([]int32, len(trainer.StorageLocations))

func main() {
	globalResources[0] = 400
	globalResources[7] = 500

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			m := make(map[string]int32, len(globalResources))
			for i, item := range trainer.StorageLocations {
				m[item] = globalResources[i]
			}
			fmt.Println(m)
		case <-quit:
			fmt.Println("Exiting")
			return
		}
	}
}
