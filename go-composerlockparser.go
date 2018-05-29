package composerlockaparser

import (
	"fmt"
)

func main() {
	c := new(composerInfo)
	c.pathToLockFile = "composer.lock"
	c.parse()

	// Get a single package by name.
	p := c.getPackageByName("drupal/embed")
	fmt.Println(p)

	// Get a list of packages by the prefix.
	pList := c.getPackageListByPrefix("drupal")
	for _, p := range pList {
		fmt.Println(p.Name)
	}
}
