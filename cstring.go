// +build linux darwin

package lexactivator

//#include <stdlib.h>
import "C"
import "unsafe"

const (
	MaxCArrayLength  C.uint = 256
	MaxGoArrayLength C.int  = 256
)

func GoToCString(data string) *C.char {
	cString := C.CString(data)
	return cString
}

func CtoGoString(cString *C.char) string {
	goString := C.GoStringN(cString, MaxGoArrayLength)
	return goString
}

func GetCArray() [MaxCArrayLength]C.char {
	var cArray [MaxCArrayLength]C.char
	return cArray
}

func FreeCString(cString *C.char) {
	defer C.free(unsafe.Pointer(cString))
}
