package prlapi

/*
#cgo LDFLAGS: -framework ParallelsVirtualizationSDK
#include <ParallelsVirtualizationSDK/Parallels.h>
*/
import "C"

import (
	"fmt"
	"github.com/rickard-von-essen/go-parallels/prlapi/key"
	"unsafe"
)

type VirtualMachine struct {
	handle           C.PRL_HANDLE
	displayConnected bool
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

func (v *VirtualMachine) DisplayConnnect() error {

	C.PrlDevDisplay_ConnectToVm(v.handle, C.enum_PDCT_HIGH_QUALITY_WITHOUT_COMPRESSION)
	//C.PrlDevDisplay_ConnectToVm(v.handle, C.PRL_DISPLAY_CODEC_TYPE(0))
	v.displayConnected = true
	return nil
}

func (v *VirtualMachine) DisplayDisconnnect() error {

	v.displayConnected = false
	res := C.PrlDevDisplay_DisconnectFromVm(v.handle)
	if res < 0 {
		return from_prl_error("PrlDevDisplay_DisconnectFromVm", res)
	}
	return nil
}

func (v *VirtualMachine) SendKeyEvent(key key.Key, event key.KeyEvent) error {

	if !v.displayConnected {
		return fmt.Errorf("Must connect to display before sending keyboard events!")
	}

	res := C.PrlDevKeyboard_SendKeyEvent(v.handle, C.PRL_UINT32(key), C.PRL_KEY_EVENT(event))
	if res < 0 {
		return from_prl_error("PrlDevKeyboard_SendKeyEvent X key X", res)
	}
	return nil
}
