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
	bufferSize256 C.uint = 256
	bufferSize1024 C.uint = 1024
	bufferSize2048 C.uint = 2048
	bufferSize4096 C.uint = 4096
	bufferSizeMax C.uint = 1000000
)

func goToCString(goString string) *C.ushort {
	bytes := []rune(goString)
	encodedBytes := utf16.Encode(bytes)
	// Ensure the slice is null-terminated
	encodedBytes = append(encodedBytes, 0)
	cString := (*C.ushort)(unsafe.Pointer(&encodedBytes[0]))
	return cString
}

func ctoGoString(cString *C.ushort, length C.uint) string {
	encodedBytes := C.GoBytes(unsafe.Pointer(cString), C.int(int(length) * 2))
	goString, _ := decodeUtf16(encodedBytes, binary.LittleEndian)
	return goString
}

func wideCtoGoString(cString *C.ushort) string {
	var encodedRunes []rune
	i := 0
	for ; ; i++ {
		runeVal := rune(*(*uint16)(unsafe.Pointer(uintptr(unsafe.Pointer(cString)) + uintptr(i)*unsafe.Sizeof(*cString))))
		if runeVal == 0 {
			break
		}
		encodedRunes = append(encodedRunes, runeVal)
	}
	goString := string(encodedRunes)
	return goString
}

func getCArray(length C.uint) []C.ushort {
	return make([]C.ushort, int(length))
}

func freeCString(cString *C.ushort) {
	// do nothing
}

func decodeUtf16(b []byte, order binary.ByteOrder) (string, error) {
	ints := make([]uint16, len(b)/2)
	if err := binary.Read(bytes.NewReader(b), order, &ints); err != nil {
		return "", err
	}
	return string(utf16.Decode(ints)), nil
}
