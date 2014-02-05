#!/usr/bin/env python
# -*- coding: utf-8 -*-
#
# (c) Parallels Software International, Inc. 2005-2009
#
#
# Example of prlsdkapi usage.
#
# Import the main Parallels Python API package.
import prlsdkapi

# Import some of the standard Python modules.
# We will not use all of them in this sample, but
# we will use them in other samples later.
import sys, time, getopt, operator, re, random
# Define constants for easy referencing of the Parallels Python API modules.
consts = prlsdkapi.prlsdk.consts
# An exception class to use to terminate the program.
class Halt(Exception):
    pass

"""
    Obtain a list of the existing virtual machines and print it
    on the screen.
    @param server: An instance of prlsdkapi.Server
                   identifying the Parallels Service.
"""
def get_vm_list(server):
    # Obtain the virtual machine list.
    # get_vm_list is an asynchronous method that returns
    # a prlsdkapi.Result object containing the list of virtual machines.
    job = server.get_vm_list()
    result = job.wait()
    print "Virtual Machine" + "                 " + "State"
    print "--------------------------------------------------"
    # Iterate through the Result object parameters.
    # Each parameter is an instance of the prlsdkapi.Vm class.
    for i in range(result.get_params_count()):
        vm = result.get_param_by_index(i)
        # Obtain the prlsdkapi.VmConfig object containing
        # the virtual machine
        # configuration information.
        vm_config = vm.get_config()
        # Get the name of the virtual machine.
        vm_name = vm_config.get_name()
        # Obtain the VmInfo object containing the
        # virtual machine state info.
        # The object is obtained from the Result object returned by
        # the vm.get_state() method.
        try:
            state_result = vm.get_state().wait()
        except prlsdkapi.PrlSDKError, e:
            print "Error: %s" % e
            return
        # Now obtain the VmInfo object.
        vm_info = state_result.get_param()
        # Get the virtual machine state code.
        state_code = vm_info.get_state()
        state_desc = "unknown status"
        # Translate the state code into a readable description.
        # For the complete list of states, see the
        # VMS_xxx constants in the Python API Reference guide.
        if state_code == consts.VMS_RUNNING:
            state_desc = "running"
        elif state_code == consts.VMS_STOPPED:
            state_desc = "stopped"
        elif state_code == consts.VMS_PAUSED:
            state_desc = "paused"
        elif state_code == consts.VMS_SUSPENDED:
            state_desc = "suspended"
        # Print the virtual machine name and status on the screen.
        vm_name = vm_name + "                         "
        print vm_name[:25] + "\t" + state_desc
        print "--------------------------------------------------"

"""
Parallels Service login.
    @param server: A new instance of the prlsdkapi.Server class.
    @param host: The host machine IP address. For local login specify "localhost".
    @param user: User name (must be a valid host OS user).
    @param password: User password.
    @param security_level: Connection secuirty level. Must be one of the
                           prlsdk.consts.PSL_xxx constants.
"""
def login_server(server, host, user, password, security_level):
    # Local or remote login?
    if host=="localhost":
        try:
            # The call returns a prlsdkapi.Result object on success.
            result = server.login_local('', 0, security_level).wait()
        except prlsdkapi.PrlSDKError, e:
            print "Login error: %s" % e
            raise Halt
    else:
        try:
            # The call returns a prlsdkapi.Result object on success.
            result = server.login(host, user, password, '', 0, 0, security_level).wait()
        except prlsdkapi.PrlSDKError, e:
            print "Login error: %s" % e
            print "Error code: " + str(e.error_code)
            raise Halt
    # Obtain a LoginResponse object contained in the Result object.
    # LoginResponse contains the results of the login operation.
    login_response = result.get_param()
    # Get the Parallels virtualization product version number.
    product_version = login_response.get_product_version()
    # Get the host operating system version.
    host_os_version = login_response.get_host_os_version()
    # Get the host UUID.
    host_uuid = login_response.get_server_uuid()
    print""
    print "Login successful"
    print""
    print "Parallels product version: " + product_version
    print "Host OS verions:           " + host_os_version
    print "Host UUID:                 " + host_uuid
    print ""

###################################################################################
def main():
    # Initialize the library for Parallels Server.
    # prlsdkapi.init_server_sdk()
    # Create a Server object and log in to a remote Parallels Server.
    # server = prlsdkapi.Server()
    # login_server(server, "10.30.18.99", "root", "secret", consts.PSL_NORMAL_SECURITY);
    # Initialize the library for Parallels Desktop.
    prlsdkapi.init_desktop_sdk()
    # Initialize the library for Parallels Desktop for Windows and Linux.
    # prlsdkapi.init_desktop_wl_sdk()
    # Initialize the library for Parallels Workstation.
    # prlsdkapi.init_workstation_sdk()
    # Initialize the library for Parallels Player.
    # prlsdkapi.init_player_sdk()
    # Create a Server object and log in to a local
    # Parallels Desktop/Workstation/Desktop for Windows and Linux/Player.
    #
    server = prlsdkapi.Server()
    login_server(server, "localhost", "", "", consts.PSL_NORMAL_SECURITY);
    get_vm_list(server)
    # Log off and deinitialize the library.
    server.logoff()
    prlsdkapi.deinit_sdk()

if __name__ == "__main__":
    try:
        sys.exit(main())
    except Halt:
        pass



