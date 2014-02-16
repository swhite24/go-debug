package debugger_test

import (
	"testing"
	"time"
	"github.com/swhite24/go-debug"
)

func TestValidDebug (t *testing.T) {
	var d = debugger.NewDebugger("hello")
	var d2 = debugger.NewDebugger("world")
	d.Log("hello", 123)
	d2.Log("world", 345)

	time.Sleep(time.Duration(2) * time.Second)

	d.Log("hello2", 123)

	time.Sleep(time.Duration(1) * time.Second)

	d2.Log("world2", 345)
	d.Log("hello3", 123)
}

func TestInvalidDebug (t *testing.T) {
	var d = debugger.NewDebugger("notset")
	d.Log("hello")
}
