package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"strings"
	"unsafe"

	"golang.org/x/mod/modfile"
)

//export getDepVer
func getDepVer(filePtr *C.char) *C.char {
	// Analyzes go.mod contents for Dependency-Analyzer
	// go build -buildmode=c-shared -o _gomod.so
	file := C.GoString(filePtr)
	f, err := modfile.Parse("", []byte(file), nil)
	if err != nil {
		return C.CString("")
	}
	dep_ver := []string{}
	for _, dep := range f.Require {
		if len(dep.Mod.Path) > 0 {
			dep_ver = append(dep_ver, dep.Mod.Path+";"+dep.Mod.Version)
		}
	}
	var minimum_go_version, modPathVer, modDeprecated string
	if f.Go != nil {
		minimum_go_version = f.Go.Version
	}
	if f.Module != nil {
		modPathVer = f.Module.Mod.Path + ";" + f.Module.Mod.Version
		modDeprecated = f.Module.Deprecated
	}
	retString := minimum_go_version + ";" + modPathVer + ";" + modDeprecated + ";" + strings.Join(dep_ver, ";")
	cstr := C.CString(retString)
	return cstr
}

//export freeCByte
func freeCByte(b unsafe.Pointer) {
	C.free(b)
}

func main() {}
