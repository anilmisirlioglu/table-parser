# Table Parser [![Made With Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg?color=007EC6)](http://golang.org)

Easily parse your cool data table in the CLI.

## Example

```go
package main

import (
	"fmt"
	"strings"
	
	"github.com/anilmisirlioglu/table-parser"
)

const input = `
REPOSITORY      TAG          IMAGE ID       CREATED         SIZE
foo             latest       cf508acd919c   26 hours ago    24.5MB
bar             latest       382715ecff56   2 months ago    705MB
baz             v2.3.5       cc88abbad18b   2 months ago    317MB
`

func main() {
	// Read the all table
	t := table.ReadAll(input)
	fmt.Printf("table header len: %d\n", len(t.Header.Cells))
	
	// With using Reader
	r := table.NewReader(strings.NewReader(input))
	fmt.Printf("table header len: %d\n", len(r.Header().Cells))
	for r.Next() {
		fmt.Printf("Row: %+v\n", r.Row())
	}
}
```

## Roadmap

- [ ] Table Writer
- [ ] Optimize the Header parser algorithm