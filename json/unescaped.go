package json

import "github.com/gocombinator/spa"

// Unescaped = %x20-21 / %x23-5B / %x5D-10FFFF
var Unescaped = spa.Char(
	"\u0020", "\u0021",
	"\u0023", "\u005B",
	"\u005D", "\u10FF",
)
