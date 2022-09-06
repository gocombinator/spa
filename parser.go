package spa

// Parser represents string parser – function mapping string to arbitrary result.
type Parser func(in string) (value any, eaten int, err error)
