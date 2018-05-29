package composerlockparser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type ComposerInfo struct {
	PathToLockFile string
	Packages       []Package
}

type Package struct {
	Name    string
	Version string
	Type    string
	Authors []string
}

func (c *ComposerInfo) Parse() {
	raw, err := ioutil.ReadFile(c.PathToLockFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	json.Unmarshal(raw, c)
}

func (c *ComposerInfo) GetPackageByName(name string) Package {
	p := Package{}
	for _, cp := range c.Packages {
		if name == cp.Name {
			p = cp
		}
	}
	return p
}

func (c *ComposerInfo) GetPackageListByPrefix(prefix string) []Package {
	packageList := []Package{}
	for _, cp := range c.Packages {
		if strings.HasPrefix(cp.Name, prefix) {
			packageList = append(packageList, cp)
		}
	}
	return packageList
}
