package shared

import (
	"ltz/arena"
)

type Coordinate struct {
	Column int
	Row int
}

type Renderable interface {
	Render(render_info Render_Info)RenderResult
}

type Render_Info struct {
	Arena_Group *arena.ArenaGroup
	Dimensions RenderingDimensions
}

type RenderResult struct {
	Buffer *[]Cell // 2D in shape
	Rows int
	Columns int
}

type RenderingDimensions struct {
	SuggestedHeight int
	SuggestedWidth int
}

type Cell struct {
	Data []byte
	DataVisualWidth uint8 // 0, 1, 2
}