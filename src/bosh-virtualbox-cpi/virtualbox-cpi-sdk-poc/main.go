package main

import (
	"fmt"
)

// #cgo CFLAGS: -g -Wall
// #include <stdlib.h>
// #include "VBoxCAPI_v7_1.h"
// #include "VBoxCAPIGlue.h"
// #include "VBoxCAPIGlue.c"
import "C"

func main() {
	fmt.Println("test")
	C.IVirtualBox.composeMachineFilename()
	//name := C.CString("Gopher")
	//defer C.free(unsafe.Pointer(name))
}
