package engine

import (
	"fmt"
	"ltz/arena"
	"ltz/elements"
	"ltz/shared"
	"sync"
)

var buffer []byte = nil
var bufferMutex sync.Mutex

var dom elements.Element = nil
var domMutex sync.Mutex

var allocationBufferInitialSize uint64 = 100 * 1024 * 1024 // 100 MB

func SetCursor(row int, col int) {
    fmt.Printf("\033[%d;%dH", row, col)
}

func ClearScreen() {
	fmt.Print("\033[2J")
}

func Run(events chan shared.Event) {
	arena_group := arena.NewArenaGroup(allocationBufferInitialSize)
	SetCursor(1, 1)

	domMutex.Lock()
	dom = Test()
	domMutex.Unlock()

	FirstPaint := dom.Render(shared.Render_Info{
		Dimensions: shared.RenderingDimensions{
			SuggestedHeight: shared.CurrentTerminalDimensions.Height,
			SuggestedWidth: shared.CurrentTerminalDimensions.Width,
		},
		Arena_Group: arena_group,
		// Buffer: &buffer,
	})

	IncrementalPrint(FirstPaint)

	ProcessEvents(events)
}

func Test() elements.Element {
	return elements.Text{Text: "Hello!"}
}

func IncrementalPrint(result shared.RenderResult) {
	bufferMutex.Lock()
	defer bufferMutex.Unlock()
}