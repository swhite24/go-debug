// package debugger provides an environment based logger
//
// In order for logs to display, the DEBUG environment variable must be set and
// match the key provided to debugger.NewDebugger
//
// DEBUG may be set as a comma-separated list, and also supports wildcards matching (*).
package debugger


import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"regexp"
	"time"
	"github.com/mgutz/ansi"
)

var (
	// Compile collection of regular expressions corresponding to Debug
	debug_r []*regexp.Regexp

	base = "\033[1;30m"

	reset = ansi.ColorCode("reset")

	colors = []string{ "red", "green", "blue", "magenta", "cyan" }

	max = 5

	last = 0
)

type (
	// Debugger provides logging based on DEBUG environment variable
	//
	// If the provided debugger key is found to be a match for the environment, logs will be displayed.
	Debugger interface {

		// Log will output log message to stdout, prefixed with key
		Log (vals ...interface{})
	}

	validDebugger struct {
		key string
		last time.Time
		first bool
		print func(string) string
	}
	invalidDebugger struct {}
)

// NewDebugger provides creates a new Debugger based the provided key.  Each debugger will
// be assigned a different color, up to five colors, then repeat.
//
// key will be compared with the environment variable "DEBUG" in order to determine if logs
// should be output.
func NewDebugger (key string) Debugger {
	// Determine type of Debugger to deliver based on environment
	for _, r := range(debug_r) {
		if r.String() != "" && r.MatchString(key) {

			color := colors[last]
			last = (last + 1) % 5

			// Valid debug
			return &validDebugger{ key, time.Time{}, true, ansi.ColorFunc(color) }
		}
	}
	// Invalid debug
	return &invalidDebugger{}
}


// validDebugger.Log will output if match
func (l *validDebugger) Log (vals ...interface{}) {
	// Determine time since last log
	diff := time.Since(l.last) / 1e6

	// Check if first log, always 0ms
	if l.first {
		diff = 0
		l.first = false
	}

	// Capture time of log
	l.last = time.Now()

	// Build log
	vals = append([]interface{}{ l.print(l.key + ":"), base }, append(vals, reset, l.print(strconv.Itoa(int(diff)) + "ms"))...)

	// Send log
	fmt.Println(vals...)
}

// invalidDebugger.Log will never output
func (l invalidDebugger) Log (vals ...interface{}) {}

func init () {
	// Grab environment and parse into regex's
	debug := strings.Split(os.Getenv("DEBUG"), ",")
	for _, val := range(debug) {
		r, _ := regexp.Compile("^" + strings.Replace(val, "*", ".*?", -1) + "$")
		debug_r = append(debug_r, r)
	}
}