package cmds

import (
	"bytes"
	"encoding/binary"
)

type SSBPCmd struct { // SuperSourceBoxProperties
	Flag       byte
	BoxId      uint8
	Enabled    bool
	Ignore1    byte // TODO: Find out what this signifies
	Source     uint16
	X          int16
	Y          int16
	Size       uint16
	Cropped    bool // for some reason this is encoded as two bytes?
	Ignore2    byte
	CropTop    uint16
	CropBottom uint16
	CropLeft   uint16
	CropRight  uint16
	Ignore3    uint16 // last two bytes are empty?
}

func (SSBPCmd) Slug() string {
	return "SSBP"
}

func (c *SSBPCmd) MarshalBinary() ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	if err := binary.Write(buf, binary.BigEndian, *c); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c *SSBPCmd) UnmarshalBinary(data []byte) error {
	reader := bytes.NewReader(data)
	err := binary.Read(reader, binary.BigEndian, c)
	return err
}
