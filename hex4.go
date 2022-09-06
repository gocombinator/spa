package spa

// Hex4 matches four hexadecimal digits (0-9, a-f, A-F).
var Hex4 = Pipe(
	Hex,
	NTimes(4),
	Concat,
)
