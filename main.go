package main

import (
	"fmt"
	"github.com/rickard-von-essen/go-parallels/prlapi"
	//"github.com/rickard-von-essen/go-parallels/prlapi/key"
)

func main() {

	server, err := prlapi.LoginLocal()
	if err == nil {
		vms, err := server.GetVms()
		if err == nil {
			i := 0
			for e := vms.Front(); e != nil; e = e.Next() {
				vm, ok := e.Value.(prlapi.VirtualMachine)
				if ok {
					i++
					fmt.Printf("(%d): %s\n", i, vm.Name())
				} else {
					fmt.Printf("fail\n")
				}
			}
		} else {
			fmt.Printf("Error: %s\n", err)
		}
		vm, err := server.GetVm("puppet-management_1392128731")
		if err == nil {
			fmt.Printf("Vm: %s\n", vm.Name())
			/*err = vm.SendKeyEvent(key.PRL_KEY_A, key.PKE_PRESS)
			if err != nil {
				fmt.Printf("error: %s\n", err)
			}
			err = vm.SendKeyEvent(key.PRL_KEY_A, key.PKE_RELEASE)
			if err != nil {
				fmt.Printf("error: %s\n", err)
			} */
		} else {
			fmt.Printf("Error: %s\n", err)
		}
		server.Disconnect()
	} else {
		fmt.Printf("Error: %s\n", err)
	}
}
