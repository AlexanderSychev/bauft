package runtime

import "errors"

type Number float64

func (n *Number) Type() Type {
	return TypeNumber
}

func (n *Number) Unwrap() any {
	return float64(*n)
}

func (n *Number) Wrap(value any) error {
	switch value.(type) {
	case float64:
		*n = Number(value.(float64))
	case *float64:
		*n = Number(*value.(*float64))
	case float32:
		*n = Number(value.(float32))
	case *float32:
		*n = Number(*value.(*float32))
	case int:
		*n = Number(value.(int))
	case *int:
		*n = Number(*value.(*int))
	case int8:
		*n = Number(value.(int8))
	case *int8:
		*n = Number(*value.(*int8))
	case int16:
		*n = Number(value.(int16))
	case *int16:
		*n = Number(*value.(*int16))
	case int32:
		*n = Number(value.(int32))
	case *int32:
		*n = Number(*value.(*int32))
	case int64:
		*n = Number(value.(int64))
	case *int64:
		*n = Number(*value.(*int64))
	case uint:
		*n = Number(value.(uint))
	case *uint:
		*n = Number(*value.(*uint))
	case uint8:
		*n = Number(value.(uint8))
	case *uint8:
		*n = Number(*value.(*uint8))
	case uint16:
		*n = Number(value.(uint16))
	case *uint16:
		*n = Number(*value.(*uint16))
	case uint32:
		*n = Number(value.(uint32))
	case *uint32:
		*n = Number(*value.(*uint32))
	case uint64:
		*n = Number(value.(uint64))
	case *uint64:
		*n = Number(*value.(*uint64))
	case bool:
		casted := value.(bool)
		if casted {
			*n = 1
		} else {
			*n = 0
		}
	default:
		return errors.New("not a number")
	}

	return nil
}
