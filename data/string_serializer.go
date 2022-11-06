package data

import (
	"encoding/binary"
	"io"
)

type stringSerializer struct{}

func (ss *stringSerializer) Serialize(wr io.Writer, value string) error {
	length := int64(len(value))

	// Write value length as 64-bit integer
	err := binary.Write(wr, binary.BigEndian, length)
	if err != nil {
		return err
	}

	// If value is not empty then write entity name value as bytes slice
	if length > 0 {
		err = binary.Write(wr, binary.BigEndian, []byte(value))
		if err != nil {
			return err
		}
	}

	return nil
}

func (ss *stringSerializer) Deserialize(wr io.Reader) (string, error) {
	var length int64
	err := binary.Read(wr, binary.BigEndian, &length)
	if err != nil {
		return "", err
	}

	if length > 0 {
		raw := make([]byte, length)
		err = binary.Read(wr, binary.BigEndian, raw)
		if err != nil {
			return "", err
		}

		return string(raw), nil
	} else {
		return "", nil
	}
}
