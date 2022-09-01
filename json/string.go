package json

import "github.com/gocombinator/spa"

// String = quotation-mark *char quotation-mark
var String = spa.Concat(spa.Seq(
	QuotationMark,
	spa.Repeat0(Char),
	QuotationMark,
))
