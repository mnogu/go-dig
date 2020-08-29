# go-dig

[![GoDev](https://img.shields.io/static/v1?label=godev&message=reference&color=00add8)](https://pkg.go.dev/mod/github.com/mnogu/go-dig)

Go version of [`Hash#dig`](https://docs.ruby-lang.org/en/2.7.0/Hash.html#method-i-dig) in Ruby

## Download and Install

```
$ go get -u github.com/mnogu/go-dig
```

## Examples

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/mnogu/go-dig"
)

func main() {
	var jsonBlob = []byte(`{"foo": {"bar": {"baz": 1}}}`)
	var h map[string]interface{}
	if err := json.Unmarshal(jsonBlob, &h); err != nil {
		fmt.Println(err)
	}
	success, err := dig.Dig(h, "foo", "bar", "baz")
	if err != nil {
		fmt.Println(err)
	}
	// foo.bar.baz = 1
	fmt.Println("foo.bar.baz =", success)

	failure, err := dig.Dig(h, "foo", "qux", "quux")
	if err != nil {
		// key qux not found in <nil>
		fmt.Println(err)
	}
	// foo.qux.quux = <nil>
	fmt.Println("foo.qux.quux =", failure)
}
```

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/mnogu/go-dig"
)

func main() {
	var jsonBlob = []byte(`{"foo": [10, 11, 12]}`)
	var h map[string]interface{}
	if err := json.Unmarshal(jsonBlob, &h); err != nil {
		fmt.Println(err)
	}
	success, err := dig.Dig(h, "foo", 1)
	if err != nil {
		fmt.Println(err)
	}
	// foo.1 = 11
	fmt.Println("foo.1 =", success)

	failure, err := dig.Dig(h, "foo", 1, 0)
	if err != nil {
		// 11 isn't a slice
		fmt.Println(err)
	}
	// foo.1.0 = <nil>
	fmt.Println("foo.1.0 =", failure)

	failure2, err := dig.Dig(h, "foo", "bar")
	if err != nil {
		// [10 11 12] isn't a map
		fmt.Println(err)
	}
	// foo.bar = <nil>
	fmt.Println("foo.bar =", failure2)
}
```