package data

import (
	"bytes"
)

type Row struct {
	entity    string
	component string
	field     string
	value     Value
}

// BytesSize returns full size of Row instance in bytes
func (r *Row) BytesSize() int {
	// The entity size is equal to the sum of the sizes of the following values
	//  - String length value (64-bit integer);
	//  - String itself (slice of bytes of length equal to the above value);
	entitySize := Int64Size + len(r.entity)

	// The component size is equal to the sum of the sizes of the following values
	//  - String length value (64-bit integer);
	//  - String itself (slice of bytes of length equal to the above value);
	componentSize := Int64Size + len(r.component)

	// The field size is equal to the sum of the sizes of the following values
	//  - String length value (64-bit integer);
	//  - String itself (slice of bytes of length equal to the above value);
	fieldSize := Int64Size + len(r.field)

	// The value size is equal to the sum of the sizes of the following values:
	//  - Value type identifier (64-bit unsigned integer);
	//  - Value size value (64-bit integer);
	//  - Value itself (slice of bytes of length equal to the above value);
	valueSize := Uint64Size + Int64Size + r.value.BytesSize()

	// The total row size is the sum of all values above
	return entitySize + componentSize + fieldSize + valueSize
}

func (r *Row) Serialize() ([]byte, error) {
	var err error
	buf := bytes.NewBuffer(make([]byte, 0))

	err = serializeString(buf, r.entity)
	if err != nil {
		return nil, err
	}

	err = serializeString(buf, r.component)
	if err != nil {
		return nil, err
	}

	err = serializeString(buf, r.component)
	if err != nil {
		return nil, err
	}

	err = serializeValue(buf, r.value)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (r *Row) Deserialize(raw []byte) error {
	var err error
	buf := bytes.NewBuffer(raw)

	entity, err := deserializeString(buf)
	if err != nil {
		return err
	}

	component, err := deserializeString(buf)
	if err != nil {
		return err
	}

	field, err := deserializeString(buf)
	if err != nil {
		return err
	}

	value, err := deserializeValue(buf)
	if err != nil {
		return err
	}

	r.entity = entity
	r.component = component
	r.field = field
	r.value = value

	return nil
}
