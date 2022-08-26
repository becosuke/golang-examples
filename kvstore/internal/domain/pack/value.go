package pack

type Value string

func NewValue(value string) *Value {
	v := Value(value)
	return &v
}

func (v *Value) String() string {
	if v == nil {
		return ""
	}
	return string(*v)
}
