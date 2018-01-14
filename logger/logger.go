package logger

import (
	"fmt"
	"runtime"
	"time"

	"github.com/fatih/color"
)

func Error(err error) {
	t := time.Now()
	stamp := t.Format("2006-01-02 15:04:05.000")
	funcName := findCaller()
	cyan := color.New(color.FgCyan).PrintfFunc()
	cyan("%s ", stamp)

	yellow := color.New(color.FgHiYellow).PrintfFunc()
	yellow("%s: ", funcName)

	red := color.New(color.FgRed).PrintfFunc()
	red("%s\n", err.Error())

}

func Log(message string, arguments ...interface{}) {
	t := time.Now()
	stamp := t.Format("2006-01-02 15:04:05.000")
	funcName := findCaller()
	cyan := color.New(color.FgCyan).PrintfFunc()
	cyan("%s ", stamp)

	yellow := color.New(color.FgHiYellow).PrintfFunc()
	yellow("%s: ", funcName)

	fmt.Printf(message+"\n", arguments...)
}

func findCaller() string {
	functionPointer := make([]uintptr, 1)
	n := runtime.Callers(3, functionPointer)
	if n == 0 {
		return "unknown function"
	}

	function := runtime.FuncForPC(functionPointer[0] - 1)
	if function == nil {
		return "unknown function"
	}
	return function.Name()
}
