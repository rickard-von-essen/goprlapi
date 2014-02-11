package prlapi

/*
#cgo LDFLAGS: -framework ParallelsVirtualizationSDK
#include <ParallelsVirtualizationSDK/Parallels.h>
*/
import "C"

import (
	"unsafe"
)

type VirtualMachine struct {
	handle C.PRL_HANDLE
}

func (v *VirtualMachine) Name() string {

	var buf = make([]byte, 1024)
	var sName C.PRL_STR = (C.PRL_STR)(unsafe.Pointer(&buf))
	var nBufSize C.PRL_UINT32 = 1024
	res := C.PrlVmCfg_GetName(v.handle, sName, &nBufSize)
	name := C.GoStringN((*_Ctype_char)(unsafe.Pointer(sName)), C.int(nBufSize-1))
	if res >= 0 {
		return name
	}
	return ""
}
