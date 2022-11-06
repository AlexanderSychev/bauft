package data

type String string

func (s *String) TypeId() uint64 {
	return TypeIdString
}

func (s *String) BytesSize() int {
	return Int64Size + len(*s)
}
