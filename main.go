package main

// #cgo LDFLAGS: -framework ParallelsVirtualizationSDK
// #include <ParallelsVirtualizationSDK/Parallels.h>
// #include <stdio.h>
import "C"

import (
	"fmt"
	//"reflect"
	"unsafe"
)

func main() {

	err := C.PrlApi_InitEx(C.PARALLELS_API_VER, C.PAM_DESKTOP, C.PRL_UINT32(0), C.PRL_UINT32(0))
	if err < 0 {
		C.PrlApi_Deinit()
		fmt.Printf("Fail! %d\n", uint(err))
		return
	} else {
		fmt.Printf("Success! %s\n", C.GoString(C.prl_result_to_string(err)))
	}

	// Declare a handle variable.
	var hServer C.PRL_HANDLE
	//Call the PrlSrv_Create to obtain the handle.
	var res = C.PrlSrv_Create(&hServer)
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
	var ret = C.PrlJob_Wait(hJob, 10000)
	if ret < 0 {
		fmt.Printf("PrlJob_Wait for PrlSrv_Login returned with error: %s\n", C.GoString(C.prl_result_to_string(ret)))
		C.PrlHandle_Free(hJob)
		C.PrlHandle_Free(hServer)
		C.PrlApi_Deinit()
		return
	}
	// Analyse the result of the PrlServer_Login call.
	var nJobResult C.PRL_RESULT
	ret = C.PrlJob_GetRetCode(hJob, &nJobResult)
	if nJobResult < 0 {
		C.PrlHandle_Free(hJob)
		C.PrlHandle_Free(hServer)
		C.PrlApi_Deinit()
		fmt.Printf("Login job returned with error: %s\n", C.GoString(C.prl_result_to_string(nJobResult)))
		return
	}
	fmt.Printf("login successfully performed\n")

	// Do stuff
	//var hStringList C.PRL_HANDLE
	//C.PrlApi_CreateStringsList(&hStringList)
	//C.PrlStrList_AddItem(hStringList, C.CString("/Users/rickard/go-code/testing/output-parallels-iso/"))
	//C.PrlStrList_AddItem(hStringList, C.CString("/Users/rickard/Documents/Parallels/Windows 7.pvm/"))
	// Begin the search operation.
	//hJob = C.PrlSrv_StartSearchVms(hServer, hStringList)
	hJob = C.PrlSrv_GetVmList(hServer)
	// Wait for the job to complete.
	ret = C.PrlJob_Wait(hJob, 10000)
	if ret < 0 {
		C.PrlHandle_Free(hJob)
		C.PrlHandle_Free(hServer)
		C.PrlApi_Deinit()
		fmt.Printf("Search job returned with error: %s\n", C.GoString(C.prl_result_to_string(ret)))
		return
	}

	var nJobReturnCode C.PRL_RESULT
	var hJobResult C.PRL_HANDLE
	// Analyze the result of PrlSrv_StartSearchVms.
	ret = C.PrlJob_GetRetCode(hJob, &nJobReturnCode)
	if ret < 0 {
		C.PrlHandle_Free(hJob)
		C.PrlHandle_Free(hServer)
		C.PrlApi_Deinit()
		fmt.Printf("GetRetCode: Search job returned with error: %s\n", C.GoString(C.prl_result_to_string(ret)))
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
	ret = C.PrlJob_GetResult(hJob, &hJobResult)
	C.PrlHandle_Free(hJob)
	if ret < 0 {
		C.PrlHandle_Free(hJob)
		C.PrlHandle_Free(hServer)
		C.PrlApi_Deinit()
		fmt.Printf("GetResult: Search job returned with error: %s\n", C.GoString(C.prl_result_to_string(ret)))
		return
	}

	// Iterate through the returned list obtaining a
	// handle of type PHT_FOUND_VM_INFO in each iteration containing
	// the information about an individual virtual machine.
	var nIndex, nCount C.PRL_UINT32
	var hFoundVmInfo C.PRL_HANDLE

	C.PrlResult_GetParamsCount(hJobResult, &nCount)
	fmt.Printf("Count: %d\n", nCount)
	for nIndex = 0; nIndex < nCount; nIndex++ {
		C.PrlResult_GetParamByIndex(hJobResult, nIndex, &hFoundVmInfo)
		// Get the virtual machine name.
		//var name = "                                        "
		//var xxName *C.char = (*C.char)(C.CString(name))
		//var sName C.PRL_CHAR = (C.char)(C.CString(*xName))
		//var nBufSize C.PRL_UINT32 = C.sizeof(sName)

		var cName = C.CString("                                          ")
		var name C.PRL_STR = C.PRL_STR(unsafe.Pointer(&cName))
		var nBufSize C.PRL_UINT32 = 25

		ret = C.PrlFoundVmInfo_GetName(hFoundVmInfo, name, &nBufSize)
		var gosName string = C.GoStringN(name, C.int(nBufSize))
		//var gosName string = nameToString(name, nBufSize)
		fmt.Printf("VM %d name: %s\n", nBufSize, gosName)
		//		// Get the name and path of the virtual machine directory.
		//		var sPath [1024]C.PRL_CHAR
		//		nBufSize = C.sizeof(sPath)
		//		ret = C.PrlFoundVmInfo_GetConfigPath(hFoundVmInfo, sPath, &nBufSize)
		//		fmt.Printf("Path: %s\n\n", C.GoString(sPath))
		//		C.PrlHandle_Free(hFoundVmInfo)
	}
	C.PrlHandle_Free(hJobResult)
	//C.PrlHandle_Free(hStringList)

	C.PrlHandle_Free(hJob)
	C.PrlHandle_Free(hServer)
	C.PrlApi_Deinit()

}

/*
func nameToString(name []C.PRL_CHAR, bufLen C.PRL_UINT32) (result string) {

	for i := 0; i < len(name); i++ {
		fmt.Printf("%x ", name[i])
	}
		var chars string
		sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&chars)))
		sliceHeader.Len = int(bufLen)
		sliceHeader.Cap = int(bufLen)
		sliceHeader.Data = uint32(unsafe.Pointer(&name))

		//var result string
		for _, value := range name {
			result += fmt.Sprintf(".%s", value)
		}
		return result[1:]
}
*/
