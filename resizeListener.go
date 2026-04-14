package main

import (
	"ltz/shared"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/term"
)

// written originally by ChatGPT
func resizeListener(events chan<- shared.Event) {
	w, h, _ := term.GetSize(int(os.Stdout.Fd()))
	shared.CurrentTerminalDimensions = shared.TermDimensions{
		Height: h,
		Width: w,
	}
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGWINCH)

	for range sig {
		w, h, err := term.GetSize(int(os.Stdout.Fd()))
		if err != nil {
			continue
		}

		events <- shared.Event{
			Type: shared.ENUM_EVENT_RESIZE,
			ResideData: &shared.ResizeEventData{
				Height: h,
				Width:  w,
			},
		}
	}
}