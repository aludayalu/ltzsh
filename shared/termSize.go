package shared

type TermDimensions struct {
	Height int
	Width int
	TotalPixelHeight int
	TotalPixelWidth int
	CellHeight int
	CellWidth int
}

var CurrentTerminalDimensions = TermDimensions{}