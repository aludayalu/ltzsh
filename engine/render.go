package engine

import (
	"fmt"
	"ltz/elements"
	"ltz/shared"
)

func Render(events <-chan shared.Event) {
	fmt.Println(Test().Render(shared.Render_Info{
		SuggestedDimensions: shared.SuggestedDimensions{
			Height: shared.CurrentTerminalDimensions.Height,
			Width: shared.CurrentTerminalDimensions.Width,
		},
	}))

	ProcessEvents(events)
}

func Test() elements.Element {
	return elements.Div{Text: "Hello!"}
}