package json

import "github.com/gocombinator/spa"

// Char bnf:
//
//	unescaped /
//	escape (
//		%x22 /          ; "    quotation mark  U+0022
//		%x5C /          ; \    reverse solidus U+005C
//		%x2F /          ; /    solidus         U+002F
//		%x62 /          ; b    backspace       U+0008
//		%x66 /          ; f    form feed       U+000C
//		%x6E /          ; n    line feed       U+000A
//		%x72 /          ; r    carriage return U+000D
//		%x74 /          ; t    tab             U+0009
//		%x75 4HEXDIG )  ; uXXXX                U+XXXX
var Char = spa.Or(
	Unescaped,
	spa.Concat(spa.Seq(Escape, spa.Or(
		spa.Char("\"\\/bfnrt"),
		spa.Concat(spa.Seq(spa.Char("u"), spa.Hex4)),
	))),
)
