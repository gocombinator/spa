package spa

// Hex matches single hexadecimal digit (0-9 / a-f / A-F) returning string of length 1.
var Hex = Char(
	"0", "9",
	"a", "f",
	"A", "F",
)
