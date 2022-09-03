package spa

// Hex2 matches two hexadecimal digits (0-9, a-f, A-F).
var Hex2 = Concat(NTimes(2, Hex))
