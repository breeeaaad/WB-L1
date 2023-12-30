package main

import "time"

func sleep(d time.Duration) {
	<-time.After(d)
}
