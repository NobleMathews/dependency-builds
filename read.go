package main

import (
	"C"
	"io/ioutil"
	"strings"

	"golang.org/x/mod/modfile"
)

//export getDepVer
func getDepVer(file string) (string, string, string, []string, error) {
	// Analyzes go.mod contents for Dependency-Analyzer
	// go build -buildmode=c-shared -o _gomod.so
	modfilename := strings.TrimSpace(string(file))
	data, err := ioutil.ReadFile(modfilename)
	if err != nil {
		return "", "", "", []string{}, err
	}

	f, err := modfile.Parse(modfilename, data, nil)
	if err != nil {
		return "", "", "", []string{}, err
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
	return minimum_go_version, modPathVer, modDeprecated, dep_ver, nil
}

func main() {}
