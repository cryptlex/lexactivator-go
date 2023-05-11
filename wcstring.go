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
	maxCArrayLength  C.uint = 1000000
	maxGoArrayLength C.int  = 1000000
)

func goToCString(goString string) *C.ushort {
	bytes := []rune(goString)
	encodedBytes := utf16.Encode(bytes)
	cString := (*C.ushort)(unsafe.Pointer(&encodedBytes[0]))
	return cString
}

func ctoGoString(cString *C.ushort) string {
	encodedBytes := C.GoBytes(unsafe.Pointer(cString), maxGoArrayLength)
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

func getCArray() [maxCArrayLength]C.ushort {
	var cArray [maxCArrayLength]C.ushort
	return cArray
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
