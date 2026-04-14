package main

import (
	"flag"
	"fmt"
	"ltz/engine"
	"ltz/shared"
	"time"
)

func main() {
	var events chan shared.Event = make(chan shared.Event)

	debugMode := flag.Bool("debug", false, "to enable keyboard debugging information")
	flag.Parse()

	listener_cleanup := func(){}

	go terminalListener(events, &listener_cleanup)

	go resizeListener(events)

	if *debugMode {
		KeyboardDebugging(events)
	} else {
		time.Sleep(time.Millisecond * 100)
		engine.Render(events)
	}

	listener_cleanup()
}

func KeyboardDebugging(events <-chan shared.Event) {
	for event := range events {
		if event.Type == shared.ENUM_EVENT_MOUSE {
			fmt.Println(event.MouseData)
		}
		if event.Type == shared.ENUM_EVENT_KEY {
			fmt.Println(event.KeyData)
		}
		if event.Type == shared.ENUM_EVENT_RESIZE {
			fmt.Println(event.ResideData)
		}
	}
}