package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell"
)

func InitScreen() (s tcell.Screen) {
	s, e := tcell.NewScreen()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	// ! the ; in the if statement allows for a value definition and then a check
	if e = s.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	return
}

func EventLoop(s tcell.Screen) {
	// ! channels can be any type if youre just closing them to set off a case
	quit := make(chan int)
	go func() {
		for {
			ev := s.PollEvent()
			// ? Not really sure how this works
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape, tcell.KeyEnter:
					close(quit)
					return
				}
			case *tcell.EventResize:
				s.Sync()
			}
		}
	}()

	// ! select statement waits for one of the cases to be called. Here we have 2 channels,
loop:
	for {
		select {
		case <-quit:
			break loop
		case <-time.After(time.Millisecond * 50):
		}
		Makebox(s)
	}
}
