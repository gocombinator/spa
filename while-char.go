package spa

// WhileChar matches pattern of runes one or more times returning string.
// See [Char] for pattern of runes.
func WhileChar(list ...string) Parser {
	return Concat(Repeat(Char(list...)))
}
