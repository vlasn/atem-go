package cmds_test

import (
	"testing"

	"github.com/mraerino/atem-go/packet/cmds"
	"github.com/stretchr/testify/assert"
)

func TestSuperSourceUnmarshalBigPicture(t *testing.T) {
	cmdBytes := []byte{
		0, 0, 1, 5, 0, 1, 0, 0, 0, 0, 3, 232, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 102, 84}
	ssbpCmd := new(cmds.SSBPCmd)
	err := ssbpCmd.UnmarshalBinary(cmdBytes)
	assert.Nil(t, err)
	assert.True(t, ssbpCmd.Enabled)
	assert.EqualValues(t, 1, ssbpCmd.Source, "source")
	assert.EqualValues(t, 0, ssbpCmd.X)
	assert.EqualValues(t, 0, ssbpCmd.Y)
	assert.EqualValues(t, 1000, ssbpCmd.Size)
	assert.True(t, true, ssbpCmd.Cropped)
	assert.EqualValues(t, 0, ssbpCmd.CropTop)
	assert.EqualValues(t, 0, ssbpCmd.CropBottom)
	assert.EqualValues(t, 0, ssbpCmd.CropLeft)
	assert.EqualValues(t, 0, ssbpCmd.CropRight)
}
func TestSuperSourceUnmarshalFourthTopLeft(t *testing.T) {
	cmdBytes := []byte{0, 0, 1, 0, 0, 1, 253, 10, 1, 169, 1, 161, 0, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	ssbpCmd := new(cmds.SSBPCmd)
	err := ssbpCmd.UnmarshalBinary(cmdBytes)
	assert.Nil(t, err)
	assert.True(t, ssbpCmd.Enabled)
	assert.EqualValues(t, 1, ssbpCmd.Source, "source")
	assert.EqualValues(t, -758, ssbpCmd.X)
	assert.EqualValues(t, 425, ssbpCmd.Y)
	assert.False(t, ssbpCmd.Cropped)
	assert.EqualValues(t, 0, ssbpCmd.CropTop)
	assert.EqualValues(t, 0, ssbpCmd.CropBottom)
	assert.EqualValues(t, 0, ssbpCmd.CropLeft)
	assert.EqualValues(t, 0, ssbpCmd.CropRight)

	marshalled, err := ssbpCmd.MarshalBinary()
	assert.Nil(t, err)
	assert.NotNil(t, marshalled)
	assert.Equal(t, cmdBytes, marshalled)
}

func TestCroppedSuperSource(t *testing.T) {
	cmdBytes := []byte{0, 0, 1, 0, 0, 1, 253, 10, 1, 169, 1, 161, 1, 8, 3, 232, 7, 208, 11, 184, 15, 160, 0, 0}
	ssbpCmd := new(cmds.SSBPCmd)
	err := ssbpCmd.UnmarshalBinary(cmdBytes)
	assert.Nil(t, err)
	assert.True(t, ssbpCmd.Enabled)
	assert.EqualValues(t, 1, ssbpCmd.Source, "source")
	assert.EqualValues(t, -758, ssbpCmd.X)
	assert.EqualValues(t, 425, ssbpCmd.Y)
	assert.True(t, ssbpCmd.Cropped)
	assert.EqualValues(t, 1000, ssbpCmd.CropTop)
	assert.EqualValues(t, 2000, ssbpCmd.CropBottom)
	assert.EqualValues(t, 3000, ssbpCmd.CropLeft)
	assert.EqualValues(t, 4000, ssbpCmd.CropRight)

	marshalled, err := ssbpCmd.MarshalBinary()
	assert.Nil(t, err)
	assert.NotNil(t, marshalled)
	assert.Equal(t, cmdBytes, marshalled)
}
