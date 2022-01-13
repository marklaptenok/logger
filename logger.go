package logger

import (
	"fmt"
	"log"
	"runtime"
)

//	the Codelearning.platform error type
type ClpError struct {
	code uint16
	msg  string
}

func (e *ClpError) Error() string {
	return fmt.Sprintf("%d - %s", e.code, e.msg)
}

//	Exported functions

func Check() error {
	//	TO-DO: replace with the actual activity
	fmt.Println("Check is activated")

	return nil
}

func Debug(err *ClpError) error {
	//	TO-DO: replace with the actual activity
	fmt.Printf("%s: %s\n", get_function_name(), err)

	return nil
}

//	Static functions

//	Returns name of function which calls this one.
func get_function_name() string {

	program_counters := make([]uintptr, 1)
	amount_of_functions_in_the_callstack := runtime.Callers(2, program_counters)
	frames := runtime.CallersFrames(program_counters[:amount_of_functions_in_the_callstack])
	frame, _ := frames.Next()

	return frame.Function
}

func to_syslog(err *ClpError) error {

	//	TO-DO: replace with the actual activity
	fmt.Printf("%s: %s\n", get_function_name(), err)

	return nil
}

func to_stderr(err *ClpError) error {

	//	Prints to stderr.
	log.Println("%s: %s\n", get_function_name(), err)

	return nil
}
