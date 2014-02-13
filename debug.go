// package debugger provides an e
package debugger


import (
	"fmt"
	"os"
	"strings"
	"regexp"
	"github.com/mgutz/ansi"
)

var (
	// Compile collection of regular expressions corresponding to Debug
	debug_r []*regexp.Regexp

	colors = []string{ "red", "green", "blue", "magenta", "cyan" }

	max = 5

	last = 0
)

type (
	// Debugger provides logging based on DEBUG environment variable
	//
	// If the provided debugger key is found to be a match for the environment, logs will be displayed.
	Debugger interface {
		Log (vals ...interface{})
	}

	validDebugger struct {
		key string
		print func(string) string
	}
	invalidDebugger struct {}
)

// NewDebugger provides creates a new Debugger based the provided key.
//
// Subsequent calls to Log will be displayed if the environment is found to be a match.
func NewDebugger (key string) Debugger {
	// Determine type of Debugger to deliver based on environment
	for _, r := range(debug_r) {
		if r.String() != "" && r.MatchString(key) {

			color := colors[last]
			last = (last + 1) % 5

			// Valid debug
			return validDebugger{ key, ansi.ColorFunc(color) }
		}
	}
	// Invalid debug
	return invalidDebugger{}
}


// validDebugger.Log will output if match
func (l validDebugger) Log (vals ...interface{}) {
	// Prepend key
	vals = append([]interface{}{l.print(l.key + ":")}, vals...)

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