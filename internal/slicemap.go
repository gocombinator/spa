package internal

func Slicemap[T, U any](vs []T, f func(T) U) []U {
	var rs []U
	for _, v := range vs {
		rs = append(rs, f(v))
	}
	return rs
}
