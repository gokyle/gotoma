package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var (
	pomodoro    time.Duration
	breakTime   time.Duration
	nPomodoros  int
	task        *string
	shouldSpeak bool
	logfile     string
)

func main() {
	config()
	for i := 0; i < nPomodoros; i++ {
		doAPomodoro()
		if i < (nPomodoros - 1) {
			takeABreak()
		}
	}
}

func duration(n int) time.Duration {
	var d time.Duration
	d, err := time.ParseDuration(fmt.Sprintf("%dm", n))
	if err != nil {
		fmt.Println("[!] couldn't parse duration: ", err)
		os.Exit(1)
	}
	return d
}

func config() {
	pTime := flag.Int("t", 25, "Length of a pomodoro.")
	bTime := flag.Int("b", 5, "Length of a break.")
	nPoms := flag.Int("n", 1, "How many pomodoros to do.")
	task = flag.String("m", "new pomodoro",
		"Name of the task being done")
	fSpeak := flag.Bool("q", false, "Don't speak on new events.")
	fLogf := flag.String("l", "", "Write to the given logfile.")
	flag.Parse()

	pomodoro = duration(*pTime)
	breakTime = duration(*bTime)
	nPomodoros = *nPoms
	shouldSpeak = !(*fSpeak)
	logfile = *fLogf
}

func doAPomodoro() {
        event := fmt.Sprintf("starting pomodoro: %s", *task)
        eventOut(event)
	<-time.After(pomodoro)
	event = fmt.Sprintf("finished pomodoro: %s", *task)
        eventOut(event)
}

func takeABreak() {
        eventOut("starting break")
	<-time.After(breakTime)
        eventOut("break's over")
}

func eventOut(event string) {
	go fmt.Println("[+] ", event)
        go logEvent(event)
	speaker(event)
}
