package elements

import (
	"ltz/arena"
	"ltz/shared"
)

// Text handles wrapping wrt parent element although might overflow in height by default.
type Text struct {
	Styles map[string]string
	Listeners map[int]func()
	Text string // Use SanatizedText for dangerous text values
}

func (element Text)Render(render_info shared.Render_Info)shared.RenderResult {
	graphemes := shared.Graphemes(element.Text)
	output_buffer_len := uint64(0)

	for _, item := range graphemes {
		output_buffer_len += item.Width
	}

	output_buffer := arena.AllocSlice[shared.Cell](render_info.Arena_Group, output_buffer_len)

	j := 0

	for i := 0; i < len(graphemes); i++ {
		// append
		output_buffer[j].Data = arena.AllocSlice[byte](render_info.Arena_Group, uint64(len(graphemes[i].Data)))
	}

	return shared.RenderResult{
		Buffer: &output_buffer,
	}
}