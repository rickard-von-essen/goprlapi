package prlapi

/*
#cgo LDFLAGS: -framework ParallelsVirtualizationSDK
#include <ParallelsVirtualizationSDK/Parallels.h>
*/
import "C"

func LoginLocal() (server Server, err error) {

	var hServer C.PRL_HANDLE
	var res C.PRL_RESULT

	res = C.PrlApi_InitEx(C.PARALLELS_API_VER, C.PAM_DESKTOP, C.PRL_UINT32(0), C.PRL_UINT32(0))
	if res < 0 {
		C.PrlApi_Deinit()
		return Server{}, from_prl_error("PrlApi_InitEx", res)
	}

	//Call the PrlSrv_Create to obtain the handle.
	res = C.PrlSrv_Create(&hServer)
	if res < 0 {
		C.PrlApi_Deinit()
		return Server{}, from_prl_error("PrlSrv_Create", res)
	}

	// Log in (PrlSrv_Login is asynchronous).
	var hJob = C.PrlSrv_LoginLocal(hServer, nil, C.PRL_UINT32(0), C.PSL_LOW_SECURITY)
	defer C.PrlHandle_Free(hJob)
	// Wait for a maximum of 10 seconds for asynchronous function PrlSrv_Login to complete.
	res = C.PrlJob_Wait(hJob, 10000)
	if res < 0 {
		defer C.PrlHandle_Free(hServer)
		defer C.PrlApi_Deinit()
		return Server{}, from_prl_error("PrlJob_Wait", res)
	}
	// Analyse the result of the PrlServer_Login call.
	var nJobResult C.PRL_RESULT
	res = C.PrlJob_GetRetCode(hJob, &nJobResult)
	if nJobResult < 0 {
		defer C.PrlHandle_Free(hServer)
		defer C.PrlApi_Deinit()
		return Server{}, from_prl_error("PrlJob_GetRetCode", res)
	}

	server.handle = hServer
	return server, nil
}
