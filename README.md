# Go Parallels SKD - C API Wrapper
![Parallels Logo](imgs/parallels_small.png)

This is Go wrapper of the
[Parallels Virtualization SDK 9 for Mac](http://www.parallels.com/downloads/desktop/) C API.

## Note

This project is still in active development and most features of the SKD is not implemented. If you are missing some dig in and open a pull request. The API of the wrapper might change, so be warned.

To understand how to use the wrapper be sure to read the [Parallales Deveoper Documentation](http://www.parallels.com/support/docs/).

I'm look forward to hearing from you with any issues or features.  Thank you!

## Requirements

Install [DMG: Parallels Virtualization SDK 9 for Mac](http://download.parallels.com//desktop/v9/pde.hf1/ParallelsVirtualizationSDK-9.0.24172.951362.dmg)
It will install under ```/Library/Frameworks/ParallelsVirtualizationSDK.framework```

## Example

```go
server, err := goprlapi.LocalConnect()
vm, err := server.getVm("My Virtual Machine")
fmt.Printf("VM name: %s", vm.Name())
```

## Copyright

Distributed under the MIT license, see LICENSE.txt.
 
© 2014 Rickard von Essen

## Contributors

* Rickard von Essen `@rickard-von-essen`
