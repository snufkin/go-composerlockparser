# Go Composer Lock Parser
Basic Composer lockfile parser library written in Golang.

# Usage

First initalise the `composerInfo` object, and set the `pathToLockFile` property
accordingly. `parse()` will read the file and perfrom the unmarshal, while
two helper methods are provided for convenience:
1. `getPackageByName()`
2. `getPackageListByPrefix()`

The first one will return the `Package` which has exactly the name as specified,
while the second will return a `[]Package`, where the name starts with the given
string.

For example the following code will parse a composer.lock file in the project,
find a project called "drupal/embed", and also find all drupal specific projects.

```go
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
```
