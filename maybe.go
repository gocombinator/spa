package spa

// Maybe matches zero or one time.
var Maybe = Catch(func(_ error, _ string) (any, int, error) {
	return nil, 0, nil
})
