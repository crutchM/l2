package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
	time "time"
)

func main() {
	r, err := ntp.Query("ntp1.stratum2.ru")
	if err != nil {
		log.Fatal(err)
	}
	t := time.Now().Add(r.ClockOffset)

	fmt.Printf("current time: %d:%d:%d", t.Hour(), t.Minute(), t.Second())
}
