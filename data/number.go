package data

import (
	"bytes"
	"encoding/binary"
)

// Number is alias type for native "float64" type which implements Serializable interface
type Number float64

func (n *Number) TypeId() uint64 {
	return TypeIdNumber
}

func (n *Number) BytesSize() int {
	return Int64Size
}

func (n *Number) Serialize() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))

	err := binary.Write(buf, binary.BigEndian, float64(*n))
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (n *Number) Deserialize(raw []byte) error {
	buf := bytes.NewBuffer(raw)

	var value float64
	err := binary.Read(buf, binary.BigEndian, &value)
	if err != nil {
		return err
	}

	*n = Number(value)
	return nil
}
