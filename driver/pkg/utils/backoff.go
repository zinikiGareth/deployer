package utils

import (
	"log"
	"time"
)

func ExponentialBackoff(check func() bool) {
	var waitFor time.Duration = 1
	for {
		time.Sleep(waitFor * time.Second)
		if check() {
			return
		}
		waitFor = min(2*waitFor, 60)
		log.Printf("waiting another %ds\n", waitFor)
	}
}
