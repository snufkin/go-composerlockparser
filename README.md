# Go Composer Lock Parser
Basic Composer lockfile parser library written in Golang.

# Usage

Initalise the `composerInfo` object, and use the `Parse(filepath)` method to
read the Composer lockfile and unmarshal it into the object defined in parser.go.

The library defines two helper methods for convenience:
1. `GetPackageByName()`
1. `GetPackageListByPrefix()`

The first one will return the `Package` which has exactly the name as specified,
while the second will return a `[]Package`, where the name starts with the given
string.

For example the following code will parse a composer.lock file in the project,
find a project called "drupal/embed", and also find all drupal specific projects.

```go
func main() {
  c := new(composerInfo)
  c.Parse("/path/to/composer.lock")

  // Get a single package by name.
  p := c.GetPackageByName("drupal/embed")
  fmt.Println(p)

  // Get a list of packages by the prefix.
  pList := c.GetPackageListByPrefix("drupal")
  for _, p := range pList {
    fmt.Println(p.Name)
  }
}
```
