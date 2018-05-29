package composerlockparser

import (
	"fmt"
	"os"
)

func main() {
	c := new(ComposerInfo)
	err := c.Parse("composer.lock")

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Get a single package by name.
	p := c.GetPackageByName("drupal/embed")
	fmt.Println(p)

	// Get a list of packages by the prefix.
	pList := c.GetPackageListByPrefix("drupal")
	for _, p := range pList {
		fmt.Println(p.Name)
	}
}
