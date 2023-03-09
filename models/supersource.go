package models

type SuperSourceBoxState struct {
	Flag       byte
	BoxId      uint8
	Enabled    bool
	Source     uint16
	X          int16
	Y          int16
	Size       uint16
	Cropped    bool
	CropTop    uint16
	CropBottom uint16
	CropLeft   uint16
	CropRight  uint16
}

type SuperSourceBoxes = map[int]SuperSourceBoxState
