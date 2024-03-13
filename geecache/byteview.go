package geecache

type ByteView struct {
	b []byte
}

// return the view's length
func (v ByteView) Len() int {
	return len(v.b)
}

// return a copy of the data
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

// return the string type of data
func (v ByteView) String() string {
	return string(v.b)
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
