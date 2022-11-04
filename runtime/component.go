package runtime

type ComponentDescription struct{}

type Component struct {
	name   string
	fields map[string]any
}
