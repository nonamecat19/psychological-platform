package utils

func UintToPointer(id int) *uint {
	u := uint(id)
	return &u
}
