package runtime

// Type enumeration defines all primitive types of Bauft
type Type byte

const (
	// TypeNil represents "nil" value - a null-pointer
	TypeNil Type = iota
	// TypeNumber represents number type - 64-bit float number
	TypeNumber
	// TypeBoolean represents boolean type
	TypeBoolean
	// TypeString represents UTF-8 string value
	TypeString
	// TypeList represents list - dynamic array (slice) which contains any other Bauft values as elements
	TypeList
	// TypeDictionary represents dictionary - map whose keys and values are any other Bauft values
	TypeDictionary
)

// Value describes any Bauft value - elementary data unit
type Value interface {
	// Type returns type code which allows to detect concrete value type
	Type() Type
	// Wrap receive native Go value and tries to set it as Bauft value.
	// It's returns error on failure.
	Wrap(value any) error
	// Unwrap returns native Go value from Bauft value container
	Unwrap() any
}
