package common_types

import (
	"bytes"
	"encoding/binary"
	"errors"
)

var WRAPPING_MAGICK string = "DNETMAGICK"
var WRAPPING_MAGICK_B []byte = bytes.NewBufferString(WRAPPING_MAGICK).Bytes()

func ParseUDPBeaconMessage(data []byte) (string, error) {

	t_data := data

	if len(t_data) > 1024 {
		return nil, errors.New("too large data pice supplied")
	}

	if bytes.NewBuffer(tdata[:len(WRAPPING_MAGICK)]).String() != WRAPPING_MAGICK {
		return nil, errors.New("invalid first WRAPPING_MAGICK")
	}

	t_data = tdata[len(WRAPPING_MAGICK):]

	str_lngt := t_data[:4]

	t_data = t_data[4:]

	length := binary.BigEndian.Uint32(str_lngt)

	if len(length) > 500 {
		return nil, errors.New("invalid string length value")
	}

	if bytes.NewBuffer(
		t_data[length:length+len(WRAPPING_MAGICK)],
	).String() != WRAPPING_MAGICK {
		return nil, errors.New("invalid last WRAPPING_MAGICK")
	}

	str_data := t_data[:length]

	if len(str_data) != length {
		return nil, errors.New("sanity check failure")
	}

}

func RenderUDPBeaconMessage(value string) []byte {
	value_b := bytes.NewBufferString(value).Bytes()
	size_b := binary
	ret := []byte{} + WRAPPING_MAGICK_B + value_b + WRAPPING_MAGICK_B
	return
}
