package prlapi

/*
#cgo LDFLAGS: -framework ParallelsVirtualizationSDK
#include <ParallelsVirtualizationSDK/Parallels.h>
*/
import "C"

import (
	"container/list"
	"fmt"
)

type Server struct {
	handle C.PRL_HANDLE
}

func (s *Server) GetVms() (vms *list.List, err error) {

	vms = list.New()
	var hJob = C.PrlSrv_GetVmList(s.handle)
	defer C.PrlHandle_Free(hJob)

	res := C.PrlJob_Wait(hJob, 10000)
	if res < 0 {
		return vms, from_prl_error("PrlJob_Wait", res)
	}

	var nJobReturnCode C.PRL_RESULT
	var hJobResult C.PRL_HANDLE
	defer C.PrlHandle_Free(hJobResult)

	res = C.PrlJob_GetRetCode(hJob, &nJobReturnCode)
	if res < 0 {
		return vms, from_prl_error("PrlJob_GetRetCode", res)
	}
	if nJobReturnCode < 0 {
		return vms, from_prl_error("PrlJob_GetRetCode", nJobReturnCode)
	}

	res = C.PrlJob_GetResult(hJob, &hJobResult)
	if res < 0 {
		return vms, from_prl_error("PrlJob_GetResult", res)
	}

	var nIndex, nCount C.PRL_UINT32

	res = C.PrlResult_GetParamsCount(hJobResult, &nCount)
	if res < 0 {
		return vms, from_prl_error("PrlResult_GetParamsCount", res)
	}

	for nIndex = 0; nIndex < nCount; nIndex++ {

		var hVm *C.PRL_HANDLE = new(C.PRL_HANDLE)
		res = C.PrlResult_GetParamByIndex(hJobResult, nIndex, hVm)
		if res >= 0 {
			vms.PushBack(VirtualMachine{handle: *hVm})
		}
	}
	return vms, nil
}

func (s *Server) GetVm(searchName string) (vm VirtualMachine, err error) {

	vms, err := s.GetVms()
	if err != nil {
		return VirtualMachine{}, err
	}

	for e := vms.Front(); e != nil; e = e.Next() {
		vm, ok := e.Value.(VirtualMachine)
		if ok && vm.Name() == searchName {
			return vm, nil
		}
	}
	return VirtualMachine{}, fmt.Errorf("No virtual machine \"%s\" was found!", searchName)
}

func (s *Server) Disconnect() {
	defer C.PrlHandle_Free(s.handle)
	defer C.PrlApi_Deinit()
}
