package spa

type Operator interface {
	Value() any
	Precedence() (int, int)
}

type operator struct {
	value      any
	precedence [2]int
}

func (o operator) Value() any {
	return o.value
}

func (o operator) Precedence() (int, int) {
	return o.precedence[0], o.precedence[1]
}
