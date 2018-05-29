package composerlockparser

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

type ComposerInfo struct {
	Packages []Package
}

type Package struct {
	Name    string
	Version string
	Type    string
	Authors []string
}

func (c *ComposerInfo) Parse(path string) error {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(raw, c)
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
