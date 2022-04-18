package convert

import (
	"errors"
)

func byteToHex(b byte) ([]byte, error) {
	toHexChar := func(n byte) (byte, error) {
		if n <= 9 {
			return n + '0', nil
		} else if n >= 10 && n <= 15 {
			return (n - 10) + 'A', nil
		}

		return 0, errors.New("hex character consists of [0-9A-F]")
	}

	xh, err := toHexChar(b >> 4)

	if err != nil {
		return nil, err
	}

	xl, err := toHexChar(b & 0b00001111)

	if err != nil {
		return nil, err
	}

	return []byte{xh, xl}, nil
}

func hexToByte(h []byte) (byte, error) {
	if len(h) != 2 {
		return 0, errors.New("hexadecimal representation of byte must contain 2 characters")
	}

	toByte := func(c byte) (byte, error) {
		if c >= '0' && c <= '9' {
			return c - '0', nil
		} else if c >= 'A' && c <= 'F' {
			return c - 'A' + 10, nil
		} else {
			return 0, errors.New("hex character consists of [0-9A-F]")
		}
	}

	xh, err := toByte(h[0])

	if err != nil {
		return 0, err
	}

	xl, err := toByte(h[1])

	if err != nil {
		return 0, err
	}

	return (xh << 4) + xl, err
}
