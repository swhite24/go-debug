package debugger

import (
	"testing"
)

func TestDebug (t *testing.T) {
	var d = NewDebugger("hello")
	var d2 = NewDebugger("world")
	d.Log("hello", 123)
	d2.Log("world", 345)
}