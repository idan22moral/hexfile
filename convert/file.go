package convert

import (
	"io/fs"
	"os"
)

func FileToHex(filename string) (string, error) {
	content, err := os.ReadFile(filename)

	if err != nil {
		return "", err
	}

	var output []byte
	for _, b := range content {
		h, err := byteToHex(b)

		if err != nil {
			return "", err
		}

		output = append(output, h[0], h[1])
	}

	return string(output), nil
}

func HexToFile(hex string, filename string, perm fs.FileMode) error {
	var data []byte

	for i := 0; i < len(hex)-1; i += 2 {
		slice := []byte(hex[i : i+2])
		b, err := hexToByte(slice)

		if err != nil {
			return err
		}

		data = append(data, b)
	}

	err := os.WriteFile(filename, data, perm)

	if err != nil {
		return err
	}

	return nil
}
