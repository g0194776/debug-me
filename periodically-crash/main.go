package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	var exitCode int
	var endTime string
	flag.StringVar(&endTime, "end", "", "when to crash, absolute time. format: 2006-01-02T15:04:05")
	flag.IntVar(&exitCode, "code", 1, "which state code you want to return after process exit. default: 1")
	flag.Parse()
	if len(endTime) == 0 {
		endTime = time.Now().Add(time.Second * 30).Format(time.RFC3339)
		fmt.Println("Argument \"end\" are missed, We've set it to 30s later.")
	}
	t, err := time.Parse(time.RFC3339, endTime)
	if err != nil {
		panic(fmt.Sprintf("The end time you typed were illegal, please try it again! error: %s", err.Error()))
	}
	wt := t.Sub(time.Now())
	if wt.Seconds() <= 0 {
		fmt.Printf("The end time you typed %s had expired, keep running!\n", endTime)
		select {} // block forever
	}

	fmt.Printf("Waiting %f seconds...\n", wt.Seconds())
	time.Sleep(wt)
	os.Exit(exitCode)
}
