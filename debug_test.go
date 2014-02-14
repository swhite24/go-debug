package debugger

import (
	"testing"
	"time"
)

func TestDebug (t *testing.T) {
	var d = NewDebugger("hello")
	var d2 = NewDebugger("world")
	d.Log("hello", 123)
	d2.Log("world", 345)

	time.Sleep(time.Duration(2) * time.Second)

	d.Log("hello2", 123)

	time.Sleep(time.Duration(1) * time.Second)

	d2.Log("world2", 345)
	d.Log("hello3", 123)
}