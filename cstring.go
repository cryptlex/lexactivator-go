// +build linux darwin

package lexactivator

import "C"

const(
	MaxCArrayLength C.uint = 256
	MaxGoArrayLength C.int = 256
)

func ToCString(data string) *C.char {
	cString := C.CString(data)
	return cString
}

func GetCArray() [MaxCArrayLength]C.char {
	var cArray [MaxCArrayLength]C.char
	return cArray;
}
