# go-debug

[![GoDoc](https://godoc.org/github.com/swhite24/go-debug?status.png)](https://godoc.org/github.com/swhite24/go-debug)

**debugger** provides an environment based debugger in code.  

Based heavily on visionmedia's [debug](https://github.com/visionmedia/debug).

## Install
```bash
$ go get github.com/swhite24/go-debug
```

## Usage

**debugger** will output logs based on the `DEBUG` environment variable.  `DEBUG` may be
comma-separated to provide multiple values, each of which may contain a wildcard `*`.  If the provided
key does not contain a match in the environment, no output will be displayed.

For instance, the following script:

```go
package main

import (
	"time"
    "github.com/swhite24/go-debug"
)

func main () {
    hello := debugger.NewDebugger("hello")
    world := debugger.NewDebugger("world")
    foo := debugger.NewDebugger("foo")

    hello.Log("Log from hello")
    world.Log("Log from world")
    foo.Log("Log from foo")

	time.Sleep(time.Duration(1) * time.Second)
	hello.Log("Log from hello after 1 sec")
}
```

when run with the command `DEBUG=hel*,world go run test.go` will output the following:

![](https://i.cloudup.com/-D6btCQh8s.png)

Note the use of wildcards and the omission of logs from the `"foo"` debugger.