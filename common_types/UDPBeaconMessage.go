package common_types

import (
	"bytes"
	"encoding/binary"
	"errors"
)

// TODO: remove address from message. only peer address presented by net.Conn
// shoulde be used

var WRAPPING_MAGICK string = "DNETMAGICK"
var WRAPPING_MAGICK_B []byte = bytes.NewBufferString(WRAPPING_MAGICK).Bytes()

func ParseUDPBeaconMessage(data []byte) (string, error) {

	t_data := data

	if len(t_data) > 1024 {
		return "", errors.New("too large data pice supplied")
	}

	if bytes.NewBuffer(t_data[:len(WRAPPING_MAGICK)]).String() !=
		WRAPPING_MAGICK {
		return "", errors.New("invalid first WRAPPING_MAGICK")
	}

	t_data = t_data[len(WRAPPING_MAGICK):]

	str_lngt := t_data[:4]

	t_data = t_data[4:]

	length := binary.BigEndian.Uint32(str_lngt)

	if length > 500 {
		return "", errors.New("invalid string length value")
	}

	if bytes.NewBuffer(
		t_data[length:length+uint32(len(WRAPPING_MAGICK))],
	).String() != WRAPPING_MAGICK {
		return "", errors.New("invalid last WRAPPING_MAGICK")
	}

	str_data := t_data[:length]

	if uint32(len(str_data)) != length {
		return "", errors.New("sanity check failure")
	}

	return bytes.NewBuffer(str_data).String(), nil
}

func RenderUDPBeaconMessage(value string) []byte {

	var size_b []byte = []byte{0, 0, 0, 0}

	value_b := bytes.NewBufferString(value).Bytes()

	binary.BigEndian.PutUint32((size_b), uint32(len(value_b)))
	ret := make([]byte, 0)
	ret = append(
		ret,
		WRAPPING_MAGICK_B...,
	)
	ret = append(
		ret,
		size_b...,
	)
	ret = append(
		ret,
		value_b...,
	)
	ret = append(
		ret,
		WRAPPING_MAGICK_B...,
	)
	return ret
}
