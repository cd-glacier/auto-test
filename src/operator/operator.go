package operator

type Type int

type Operator struct {
	Type    Type
	Literal string
}

const (
	_ = iota
	TO_MAXINT
)
