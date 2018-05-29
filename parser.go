package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type composerInfo struct {
	pathToLockFile string
	Packages       []Package
}

type Package struct {
	Name    string
	Version string
	Type    string
	Authors []string
}

func (c *composerInfo) parse() {
	raw, err := ioutil.ReadFile(c.pathToLockFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	json.Unmarshal(raw, c)
}

func (c *composerInfo) getPackageByName(name string) Package {
	p := Package{}
	for _, cp := range c.Packages {
		if name == cp.Name {
			p = cp
		}
	}
	return p
}

func (c *composerInfo) getPackageListByPrefix(prefix string) []Package {
	packageList := []Package{}
	for _, cp := range c.Packages {
		if strings.HasPrefix(cp.Name, prefix) {
			packageList = append(packageList, cp)
		}
	}
	return packageList
}
