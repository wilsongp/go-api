package utils

//Coalesce returns the first non-nil value in the supplied arguments
func Coalesce(a, b interface{}) interface{} {
	if a != nil {
		return a
	}
	if b != nil {
		return b
	}
	return nil
}
