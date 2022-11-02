[![Go Reference](https://pkg.go.dev/badge/aead.dev/mem.svg)](https://pkg.go.dev/aead.dev/mem)
![Github CI](https://github.com/aead/mem/actions/workflows/build.yml/badge.svg?branch=main)

The `mem` package provides types and functions for measuring and displaying memory throughput and capacity.

### Example

Walk the directory tree of the current directory and list all (non-hidden) files and the corresponding file sizes.

```go
package main

import (
    "fmt"
    "io/fs"
    "path/filepath"
    "strings"

    "aead.dev/mem"
)

func main() {
    const Directory = "."
    filepath.WalkDir(Directory, func(path string, entry fs.DirEntry, err error) error {
        if entry.IsDir() {
            if entry.Name() != Directory && strings.HasPrefix(entry.Name(), ".") { // Skip hidden sub directories
                return filepath.SkipDir
            }
            fmt.Println(entry.Name())
            return nil
        }

        stat, err := entry.Info()
        if err == nil {
            var size mem.Size = mem.Size(stat.Size())
            fmt.Printf("  %-20s %10s\n", stat.Name(), mem.FormatSize(size, 'D', 2))
        }
        return nil
    })
}
```
<details><summary>Example Ouput</summary>

```
.
  .gitignore              269.00B
  .golangci.yml           420.00B
  LICENSE                  1.08KB
  benchmark_test.go        1.47KB
  bitsize.go               2.79KB
  bitsize_test.go          3.38KB
cmd
example
  main.go                 605.00B
  doc.go                   2.91KB
  examples_test.go         1.76KB
  format.go               10.15KB
  format_test.go           4.24KB
  go.mod                   29.00B
  math.go                 831.00B
  math_test.go             2.28KB
  progress.go             842.00B
  size.go                  3.00KB
  size_test.go             3.73KB
```

</details>

## Install

Add `aead.dev/mem` as dependency to your `go.mod` file:

```
go mod download aead.dev/mem@latest
```
