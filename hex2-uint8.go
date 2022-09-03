package spa

// Hex2Uint8 matches two hexadecimal digits returning uint8 result.
var Hex2Uint8 = As(Pair(Hex, Hex), func(vs [2]uint8) uint8 {
	return vs[0]<<4 | vs[1]
})
