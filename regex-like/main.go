package main

import (
	"flag"
	"github.com/teejays/clog"
	"time"
)

// LOG const determines whether the program logs to the Std. Output
const LOG bool = false

const (
	DEBUG uint = iota
	INFO
	NOTICE
)

func main() {
	// allow for pattern and the str to be passed as arguments to the binary execuatble
	var pattern, str string
	flag.StringVar(&pattern, "pattern", "", "the pattern used to match the string")
	flag.StringVar(&str, "string", "", "the string which is matched using the pattern")
	flag.Parse()

	run(pattern, str)
}

func run(pattern, str string) bool {
	start := time.Now()
	answer := Match(pattern, str)
	elapsed := time.Now().Sub(start)
	clog.Infof("Pattern: '%s' || String: '%s' ||  Time: %s => Answer: %t", pattern, str, elapsed, answer)
	return answer
}

func log(logger uint, message string, args ...interface{}) {
	if LOG {
		switch logger {
		case DEBUG:
			clog.Debugf(message, args...)
		case INFO:
			clog.Infof(message, args...)
		case NOTICE:
			clog.Warningf(message, args...)
		}

	}
}
