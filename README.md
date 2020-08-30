# go-dig

[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/mnogu/go-dig)](https://pkg.go.dev/mod/github.com/mnogu/go-dig)
[![GitHub Actions](https://github.com/mnogu/go-dig/workflows/Go/badge.svg)](https://github.com/mnogu/go-dig/actions?query=workflow%3AGo)

Go version of [`Hash#dig`](https://docs.ruby-lang.org/en/2.7.0/Hash.html#method-i-dig) and [`Array#dig`](https://docs.ruby-lang.org/en/2.7.0/Array.html#method-i-dig) in Ruby

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
	var v interface{}
	if err := json.Unmarshal(jsonBlob, &v); err != nil {
		fmt.Println(err)
	}
	success, err := dig.Dig(v, "foo", "bar", "baz")
	if err != nil {
		fmt.Println(err)
	}
	// foo.bar.baz = 1
	fmt.Println("foo.bar.baz =", success)

	failure, err := dig.Dig(v, "foo", "qux", "quux")
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
	var v interface{}
	if err := json.Unmarshal(jsonBlob, &v); err != nil {
		fmt.Println(err)
	}
	success, err := dig.Dig(v, "foo", 1)
	if err != nil {
		fmt.Println(err)
	}
	// foo.1 = 11
	fmt.Println("foo.1 =", success)

	failure, err := dig.Dig(v, "foo", 1, 0)
	if err != nil {
		// 11 isn't a slice
		fmt.Println(err)
	}
	// foo.1.0 = <nil>
	fmt.Println("foo.1.0 =", failure)

	failure2, err := dig.Dig(v, "foo", "bar")
	if err != nil {
		// [10 11 12] isn't a map
		fmt.Println(err)
	}
	// foo.bar = <nil>
	fmt.Println("foo.bar =", failure2)
}
```
