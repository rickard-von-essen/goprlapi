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
)

func from_prl_error(function string, result C.PRL_RESULT) (err error) {
	if result == 0 {
		err = nil
	} else {
		err = fmt.Errorf("Parallels API: %s error: %s", function, C.GoString(C.prl_result_to_string(result)))
	}
	return
}
