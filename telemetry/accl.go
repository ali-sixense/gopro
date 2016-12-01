package telemetry

import (
	"encoding/binary"
	"errors"
)

type ACCL struct {
	X float64
	Y float64
	Z float64
}

func (accl *ACCL) Parse(bytes []byte, scale *SCAL) error {
	if 6 != len(bytes) {
		return errors.New("Invalid length ACCL packet")
	}

	// ¯\_(ツ)_/¯
	accl.X = float64(int16(binary.BigEndian.Uint16(bytes[0:2])) / int16(scale.Values[0]))
	accl.Y = float64(int16(binary.BigEndian.Uint16(bytes[2:4])) / int16(scale.Values[0]))
	accl.Z = float64(int16(binary.BigEndian.Uint16(bytes[4:6])) / int16(scale.Values[0]))

	return nil
}