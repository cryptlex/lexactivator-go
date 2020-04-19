// +build windows

package lexactivator

import "C"
import (
	"bytes"
	"encoding/binary"
	"unicode/utf16"
	"unsafe"
)

const (
	MaxCArrayLength  C.uint = 256
	MaxGoArrayLength C.int  = 256
)

func GoToCString(goString string) *C.ushort {
	bytes := []rune(goString)
	encodedBytes := utf16.Encode(bytes)
	cString := (*C.ushort)(unsafe.Pointer(&encodedBytes[0]))
	return cString
}

func CtoGoString(cString *C.ushort) string {
	encodedBytes := C.GoBytes(unsafe.Pointer(cString), MaxGoArrayLength)
	goString, _ := decodeUtf16(encodedBytes, binary.LittleEndian)
	return goString
}

func GetCArray() [MaxCArrayLength]C.ushort {
	var cArray [MaxCArrayLength]C.ushort
	return cArray
}

func FreeCString(cString *C.ushort) {
	// do nothing
}

func decodeUtf16(b []byte, order binary.ByteOrder) (string, error) {
	ints := make([]uint16, len(b)/2)
	if err := binary.Read(bytes.NewReader(b), order, &ints); err != nil {
		return "", err
	}
	return string(utf16.Decode(ints)), nil
}
