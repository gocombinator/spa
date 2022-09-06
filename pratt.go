package spa

import "fmt"

type PrecedenceResult struct {
	Value       any
	Left, Right int
}

type PrefixResult struct {
	Op, Value any
}

func (r PrefixResult) String() string {
	return fmt.Sprintf("(%v %v)", r.Op, r.Value)
}

type InfixResult struct {
	Left, Op, Right any
}

func (r InfixResult) String() string {
	return fmt.Sprintf("(%v %v %v)", r.Op, r.Left, r.Right)
}

type SuffixResult struct {
	Value, Op any
}

func (r SuffixResult) String() string {
	return fmt.Sprintf("(%v %v)", r.Op, r.Value)
}

func pratt(p, prefix, infix, postfix Parser, min int) Parser {
	return func(in string) (any, int, error) {
		var lhs any
		var o = 0
		if op, w, err := prefix(in); err == nil {
			o += w
			var precedence = 0
			switch op := op.(type) {
			case PrecedenceResult:
				precedence = op.Right
			}
			if rhs, w, err := pratt(p, prefix, infix, postfix, precedence)(in[o:]); err == nil {
				o += w
				lhs = &PrefixResult{Op: op.(PrecedenceResult).Value, Value: rhs}
			} else {
				return nil, 0, err
			}
		} else if v, w, err := p(in); err == nil {
			o += w
			lhs = v
		} else {
			return nil, 0, err
		}
	outer:
		for {
			if op, w, err := postfix(in[o:]); err == nil {
				switch op := op.(type) {
				case PrecedenceResult:
					if op.Left < min {
						break outer
					}
				default:
					return nil, 0, fmt.Errorf("postfix operator must return PrecedenceResult")
				}
				o += w
				lhs = &SuffixResult{Value: lhs, Op: op.(PrecedenceResult).Value}
				continue
			}
			if op, w, err := infix(in[o:]); err == nil {
				var lp, rp int
				switch op := op.(type) {
				case PrecedenceResult:
					lp, rp = op.Left, op.Right
				}
				if lp < min {
					break
				}
				o += w
				if rhs, w, err := pratt(p, prefix, infix, postfix, rp)(in[o:]); err == nil {
					o += w
					lhs = &InfixResult{Left: lhs, Op: op.(PrecedenceResult).Value, Right: rhs}
				} else {
					return nil, 0, err
				}
				continue
			}
			break
		}
		return lhs, o, nil
	}
}

func Pratt(p, prefix, infix, postfix Parser) Parser {
	return pratt(p, prefix, infix, postfix, 0)
}
