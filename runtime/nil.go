package runtime

type Nil struct{}

func (n Nil) Type() Type {
	return TypeNil
}

func (n Nil) Wrap(value any) error {
	return nil
}

func (n Nil) Unwrap() any {
	return nil
}
