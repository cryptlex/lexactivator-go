// +build windows

package lexactivator

import "C"
import (
	"unicode/utf16"
	"unsafe"
	"bytes"
	"encoding/binary"
)

const (
	MaxCArrayLength  C.uint = 256
	MaxGoArrayLength C.int  = 256
)

func GoToCString(goString string) *C.ushort {
	runeByte := []rune(goString)
	encodedByte := utf16.Encode(runeByte)
	cString := (*C.ushort)(unsafe.Pointer(&encodedByte[0]))
	return cString
}

func CtoGoString(cString *C.ushort) string {
	encodedBytes := C.GoBytes(unsafe.Pointer(cString), MaxGoArrayLength)
	goString,_ := decodeUtf16(encodedBytes, binary.LittleEndian)
	return goString
	//decodedBytes := utf16.Decode(encodedBytes)
	//return string(decodedBytes)
}

func GetCArray() [MaxCArrayLength]C.ushort {
	var cArray [MaxCArrayLength]C.ushort
	return cArray
}

func decodeUtf16(b []byte, order binary.ByteOrder) (string, error) {
	ints := make([]uint16, len(b)/2)
	if err := binary.Read(bytes.NewReader(b), order, &ints); err != nil {
		return "", err
	}
	return string(utf16.Decode(ints)), nil
}
