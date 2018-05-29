package composerlockparser

import (
	"fmt"
)

func main() {
	c := new(ComposerInfo)
	c.PathToLockFile = "composer.lock"
	c.Parse()

	// Get a single package by name.
	p := c.GetPackageByName("drupal/embed")
	fmt.Println(p)

	// Get a list of packages by the prefix.
	pList := c.GetPackageListByPrefix("drupal")
	for _, p := range pList {
		fmt.Println(p.Name)
	}
}
