package data

import (
	"bytes"
	"encoding/binary"
	"io"
	"math"
)

const (
	// TypeIdNil represents "nil" (empty) type
	TypeIdNil uint64 = math.MaxUint64 - iota
	// TypeIdNumber represents Number type
	TypeIdNumber
	// TypeIdBoolean represents Boolean type
	TypeIdBoolean
	// TypeIdString represents String type
	TypeIdString
	// TypeIdList represents List type
	TypeIdList
	// TypeIdDictionary represents Dictionary type
	TypeIdDictionary
)

const (
	// Int64Size is size of "int64" type value in bytes
	Int64Size = 8
	// Uint64Size is size of "uint64" type value in bytes
	Uint64Size = 8
	// NilSize is size of "nil" (empty) value in bytes
	NilSize = 0
	// Float64Size is size of "float64" type value in bytes
	Float64Size = 8
	// BoolSize is size of "boolean" value in bytes
	BoolSize = 1
)

// Serializable describes an entity that can be converted to a byte slice and restored from a byte slice
type Serializable interface {
	// BytesSize returns entity size in bytes
	BytesSize() int
	// Serialize transforms entity to bytes slice. Method can return error as second value on transformation error.
	Serialize() ([]byte, error)
	// Deserialize restores entity data from received bytes slice
	Deserialize(raw []byte) error
}

// Value describes Row value which
type Value interface {
	Serializable
	// TypeId returns type unique identifier. This identifier need to
	// serialize and deserialize mixed Value implementations collections.
	TypeId() uint64
}

type Serializer[T any] interface {
	// Serialize transforms entity to bytes slice. Method can return error as second value on transformation error.
	Serialize(wr io.Writer, value T) error
	// Deserialize restores entity data from received bytes slice
	Deserialize(wr io.Reader) (T, error)
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// Serialization/deserialization helpers
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

func serializeString(buf *bytes.Buffer, value string) error {
	length := int64(len(value))

	// Write value length as 64-bit integer
	err := binary.Write(buf, binary.BigEndian, length)
	if err != nil {
		return err
	}

	// If value is not empty then write entity name value as bytes slice
	if length > 0 {
		err = binary.Write(buf, binary.BigEndian, []byte(value))
		if err != nil {
			return err
		}
	}

	return nil
}

func deserializeString(buf *bytes.Buffer) (string, error) {
	var length int64
	err := binary.Read(buf, binary.BigEndian, &length)
	if err != nil {
		return "", err
	}

	if length > 0 {
		raw := make([]byte, length)
		err = binary.Read(buf, binary.BigEndian, raw)
		if err != nil {
			return "", err
		}

		return string(raw), nil
	} else {
		return "", nil
	}
}

func serializeValue(buf *bytes.Buffer, value Value) error {
	var err error

	// Serialize row value
	if value != nil {
		// Write value type unique identifier as 64-bit unsigned integer
		err = binary.Write(buf, binary.BigEndian, value.TypeId())
		if err != nil {
			return err
		}

		// Write value size as 64-bit integer
		bytesSize := int64(value.BytesSize())
		err = binary.Write(buf, binary.BigEndian, bytesSize)
		if err != nil {
			return err
		}

		// If value size is not zero then serialize and write it
		if bytesSize > 0 {
			var rawValue []byte
			rawValue, err = value.Serialize()
			if err != nil {
				return err
			}

			err = binary.Write(buf, binary.BigEndian, rawValue)
			if err != nil {
				return err
			}
		}
	} else {
		// If value is "nil" then just write marker for "nil" value
		err = binary.Write(buf, binary.BigEndian, TypeIdNil)
		if err != nil {
			return err
		}
	}

	return nil
}

func deserializeValue(buf *bytes.Buffer) (Value, error) {
	var err error

	var typeId uint64
	err = binary.Read(buf, binary.BigEndian, &typeId)
	if err != nil {
		return nil, err
	}

	var value Value = nil
	if typeId != TypeIdNil {
		var bytesSize int64

		err = binary.Read(buf, binary.BigEndian, &bytesSize)
		if err != nil {
			return nil, err
		}

		if bytesSize > 0 {
			valueRaw := make([]byte, bytesSize)
			err = binary.Read(buf, binary.BigEndian, valueRaw)
			if err != nil {
				return nil, err
			}

			value, err = ValueFactory.createValueByTypeId(typeId)
			if err != nil {
				return nil, err
			}

			err = value.Deserialize(valueRaw)
			if err != nil {
				return nil, err
			}
		}
	}

	return value, nil
}
