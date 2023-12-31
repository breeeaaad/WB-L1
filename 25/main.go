package main

import "time"

func sleep(d time.Duration) {
	<-time.After(d)
}

func main() {
	sleep(5 * time.Second)
}
