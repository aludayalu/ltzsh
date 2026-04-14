package elements

import (
	"ltz/shared"
)

type Div struct {
	Styles map[string]string
	Children []Element
	Listeners map[int]func()
	Text string
}

func (Div)Render(render_info shared.Render_Info)shared.RenderResult {
	return shared.RenderResult{}
}

