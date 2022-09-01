package spa

// Ws1 matches one or more whitespace characters.
// See [Ws0] for zero or more whitespaces.
var Ws1 = MinLen(1, Ws0)
