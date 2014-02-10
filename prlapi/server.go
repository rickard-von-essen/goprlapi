package prlapi

/*
#cgo LDFLAGS: -framework ParallelsVirtualizationSDK
#include <ParallelsVirtualizationSDK/Parallels.h>

extern const char * prl_result_to_string(
  PRL_RESULT value
);
*/
import "C"

import (
	"fmt"
	"unsafe"
)

type Server struct {
	handle C.PRL_HANDLE
}

func (s *Server) PrintVmList() error {

	var hJob = C.PrlSrv_GetVmList(s.handle)
	defer C.PrlHandle_Free(hJob)

	res := C.PrlJob_Wait(hJob, 10000)
	if res < 0 {
		return from_prl_error("PrlJob_Wait", res)
	}

	var nJobReturnCode C.PRL_RESULT
	var hJobResult C.PRL_HANDLE
	defer C.PrlHandle_Free(hJobResult)

	res = C.PrlJob_GetRetCode(hJob, &nJobReturnCode)
	if res < 0 {
		return from_prl_error("PrlJob_GetRetCode", res)
	}
	if nJobReturnCode < 0 {
		return from_prl_error("PrlJob_GetRetCode", nJobReturnCode)
	}

	res = C.PrlJob_GetResult(hJob, &hJobResult)
	if res < 0 {
		return from_prl_error("PrlJob_GetResult", res)
	}

	var nIndex, nCount C.PRL_UINT32

	res = C.PrlResult_GetParamsCount(hJobResult, &nCount)
	if res < 0 {
		return from_prl_error("PrlResult_GetParamsCount", res)
	}

	for nIndex = 0; nIndex < nCount; nIndex++ {

		var hVm *C.PRL_HANDLE = new(C.PRL_HANDLE)
		defer C.PrlHandle_Free(*hVm)

		res = C.PrlResult_GetParamByIndex(hJobResult, nIndex, hVm)
		if res >= 0 {

			var buf = make([]byte, 1024)
			var sName C.PRL_STR = (C.PRL_STR)(unsafe.Pointer(&buf))
			var nBufSize C.PRL_UINT32 = 1024
			res = C.PrlVmCfg_GetName(*hVm, sName, &nBufSize)
			name := C.GoStringN((*_Ctype_char)(unsafe.Pointer(sName)), C.int(nBufSize))
			if res >= 0 {
				fmt.Printf("(%d) : \"%s\"\n", nIndex, name)
			}
		}
	}
	return nil
}

func (s *Server) Disconnect() {
	defer C.PrlHandle_Free(s.handle)
	defer C.PrlApi_Deinit()
}
