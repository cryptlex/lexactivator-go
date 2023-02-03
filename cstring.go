// +build linux darwin

package lexactivator

//#include <stdlib.h>
import "C"
import "unsafe"

const (
	maxCArrayLength  C.uint = 1024
	maxGoArrayLength C.int  = 1024
)

func goToCString(data string) *C.char {
	cString := C.CString(data)
	return cString
}

func ctoGoString(cString *C.char) string {
	goString := C.GoStringN(cString, maxGoArrayLength)
	return goString
}

func getCArray() [maxCArrayLength]C.char {
	var cArray [maxCArrayLength]C.char
	return cArray
}

func freeCString(cString *C.char) {
	defer C.free(unsafe.Pointer(cString))
}
