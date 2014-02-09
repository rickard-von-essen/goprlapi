package main

// #cgo LDFLAGS: -framework ParallelsVirtualizationSDK
// #include <ParallelsVirtualizationSDK/Parallels.h>
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {

	hServer := LoginLocal()
	GetVmList(hServer)
}

func LoginLocal() (hServer C.PRL_HANDLE) {

	err := C.PrlApi_InitEx(C.PARALLELS_API_VER, C.PAM_DESKTOP, C.PRL_UINT32(0), C.PRL_UINT32(0))
	if err < 0 {
		C.PrlApi_Deinit()
		fmt.Printf("Fail! %d\n", uint(err))
		return
	} else {
		fmt.Printf("Success! %s\n", C.GoString(C.prl_result_to_string(err)))
	}

	//Call the PrlSrv_Create to obtain the handle.
	res := C.PrlSrv_Create(&hServer)
	// Examine the function return code.
	// PRL_FAILED is a macro that evaluates a variable of type PRL_RESULT.
	// A return value of True indicates success; False indicates failure.
	if res < 0 {
		C.PrlApi_Deinit()
		fmt.Printf("PrlSrv_Create returned error: %s\n", C.GoString(C.prl_result_to_string(res)))
		return
	}
	fmt.Printf("PrlSrv_Create return fine: %s\n", C.GoString(C.prl_result_to_string(res)))

	// Log in (PrlSrv_Login is asynchronous).
	var hJob = C.PrlSrv_LoginLocal(hServer, nil, C.PRL_UINT32(0), C.PSL_LOW_SECURITY)
	// Wait for a maximum of 10 seconds for
	// asynchronous function PrlSrv_Login to complete.
	res = C.PrlJob_Wait(hJob, 10000)
	if res < 0 {
		fmt.Printf("PrlJob_Wait for PrlSrv_Login returned with error: %s\n", C.GoString(C.prl_result_to_string(res)))
		C.PrlHandle_Free(hJob)
		C.PrlHandle_Free(hServer)
		C.PrlApi_Deinit()
		return
	}
	// Analyse the result of the PrlServer_Login call.
	var nJobResult C.PRL_RESULT
	res = C.PrlJob_GetRetCode(hJob, &nJobResult)
	if nJobResult < 0 {
		C.PrlHandle_Free(hJob)
		C.PrlHandle_Free(hServer)
		C.PrlApi_Deinit()
		fmt.Printf("Login job returned with error: %s\n", C.GoString(C.prl_result_to_string(nJobResult)))
		return
	}

	fmt.Printf("login successfully performed\n")
	return hServer
}

func GetVmList(hServer C.PRL_HANDLE) {
	// Begin the search operation.
	var hJob = C.PrlSrv_GetVmList(hServer)
	// Wait for the job to complete.
	res := C.PrlJob_Wait(hJob, 10000)
	if res < 0 {
		C.PrlHandle_Free(hJob)
		C.PrlHandle_Free(hServer)
		C.PrlApi_Deinit()
		fmt.Printf("Search job returned with error: %s\n", C.GoString(C.prl_result_to_string(res)))
		return
	}

	var nJobReturnCode C.PRL_RESULT
	var hJobResult C.PRL_HANDLE
	// Analyze the result of PrlSrv_StartSearchVms.
	res = C.PrlJob_GetRetCode(hJob, &nJobReturnCode)
	if res < 0 {
		C.PrlHandle_Free(hJob)
		C.PrlHandle_Free(hServer)
		C.PrlApi_Deinit()
		fmt.Printf("GetRetCode: Search job returned with error: %s\n", C.GoString(C.prl_result_to_string(res)))
		return
	}
	// Check the job return code.
	if nJobReturnCode < 0 {
		C.PrlHandle_Free(hJob)
		C.PrlHandle_Free(hServer)
		C.PrlApi_Deinit()
		fmt.Printf("GetRetCode 2: Search job returned with error: %s\n", C.GoString(C.prl_result_to_string(nJobReturnCode)))
		return
	}
	// Get job result.
	res = C.PrlJob_GetResult(hJob, &hJobResult)
	C.PrlHandle_Free(hJob)
	if res < 0 {
		C.PrlHandle_Free(hJob)
		C.PrlHandle_Free(hServer)
		C.PrlApi_Deinit()
		fmt.Printf("GetResult: Search job returned with error: %s\n", C.GoString(C.prl_result_to_string(res)))
		return
	}

	// Iterate through the returned list obtaining a
	// handle of type PHT_FOUND_VM_INFO in each iteration containing
	// the information about an individual virtual machine.
	var nIndex, nCount C.PRL_UINT32

	res = C.PrlResult_GetParamsCount(hJobResult, &nCount)
	if res < 0 {
		fmt.Printf("PrlResult_GetParamsCount returned: %s\n", C.GoString(C.prl_result_to_string(res)))
	}
	fmt.Printf("Count: %d\n", nCount)
	for nIndex = 0; nIndex < nCount; nIndex++ {

		var hVm *C.PRL_HANDLE = new(C.PRL_HANDLE)
		res = C.PrlResult_GetParamByIndex(hJobResult, nIndex, hVm)
		if res < 0 {
			fmt.Printf("PrlResult_GetParamByIndex returned: %s\n", C.GoString(C.prl_result_to_string(res)))
		}

		var buf = make([]byte, 1024)
		var sName C.PRL_STR = (C.PRL_STR)(unsafe.Pointer(&buf))
		var nBufSize C.PRL_UINT32 = 1024
		res = C.PrlVmCfg_GetName(*hVm, sName, &nBufSize)
		fmt.Printf("(%d) ", nIndex)
		name := C.GoStringN((*_Ctype_char)(unsafe.Pointer(sName)), C.int(nBufSize))
		fmt.Printf(": \"%s\"\n", name)
		if res < 0 {
			fmt.Printf("PrlFoundVmInfo_GetName returned: %s, %d\n", C.GoString(C.prl_result_to_string(res)), nBufSize)
		}
	}
	C.PrlHandle_Free(hJobResult)

	C.PrlHandle_Free(hJob)
	C.PrlHandle_Free(hServer)
	C.PrlApi_Deinit()
}
