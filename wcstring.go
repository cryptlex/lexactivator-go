// +build windows

package lexactivator

import "C"
import (
	"unicode/utf16"
	"unsafe"
)

const (
	MaxCArrayLength  C.uint = 256
	MaxGoArrayLength C.int    = 256
)

func ToCString(data string) *C.ushort {
	runeByte := []rune(data)
	encodedByte := utf16.Encode(runeByte)
	cString := (*C.ushort)(unsafe.Pointer(&encodedByte[0]))
	return cString
}

func GetCArray() [MaxCArrayLength]C.ushort {
	var cArray [MaxCArrayLength]C.char
	return cArray;
}
