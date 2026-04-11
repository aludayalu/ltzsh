package components

import (
	"ltz/shared"
)

type Div_Element struct {

}

func div(props ...any) Element {
	return Div_Element{}
}

func (b Div_Element) Render(render_info shared.Render_Info) shared.RenderResult {
	return shared.RenderResult{}
}