package components

import (
	"ltz/shared"
)

type Element interface {
	shared.Renderable
}

type Props = map[string]any