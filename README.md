# go-debug

**debugger** provides an environment based debugger in code.  

Based heavily on visionmedia's [debug](https://github.com/visionmedia/debug).

## Install
```bash
$ go get github.com/swhite24/go-debug
```

## Usage

**debugger** will output logs based on the `DEBUG` environment variable.  `DEBUG` may be
comma-separated to provided multiple values, each of which may contain a wildcard `*`.  If the provided
key does not contain a match in the environment, no output will be displayed.

```go
package main

import (
	"github.com/swhite24/go-debug"
)

func main () {
	d := debugger.NewDebugger("testing")
	d.Log("test log", 123)
}
```