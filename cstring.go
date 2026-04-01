// +build linux darwin

package lexactivator

//#include <stdlib.h>
import "C"
import "unsafe"

const (
	bufferSize256 C.uint = 256
	bufferSize1024 C.uint = 1024
	bufferSize2048 C.uint = 2048
	bufferSize4096 C.uint = 4096
	bufferSizeMax C.uint = 4096
)

func goToCString(data string) *C.char {
	cString := C.CString(data)
	return cString
}

func ctoGoString(cString *C.char, length C.uint) string {
	goString := C.GoString(cString)
	return goString
}

func getCArray(length C.uint) []C.char {
	return make([]C.char, int(length))
}

func freeCString(cString *C.char) {
	defer C.free(unsafe.Pointer(cString))
}
