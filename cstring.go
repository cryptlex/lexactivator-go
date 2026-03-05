// +build linux darwin

package lexactivator

//#include <stdlib.h>
import "C"
import "unsafe"

const cArrayLength C.uint = 1024
const maxCArrayLength C.uint = 4096

func goToCString(data string) *C.char {
	cString := C.CString(data)
	return cString
}

func ctoGoString(cString *C.char) string {
	goString := C.GoString(cString)
	return goString
}

func getCArray(length ...C.uint) []C.char {
	size := cArrayLength
	if len(length) > 0 {
		size = length[0]
	}
	return make([]C.char, size)
}

func freeCString(cString *C.char) {
	defer C.free(unsafe.Pointer(cString))
}
