package logger

//	TO-DO:	use the 'log' package to format (add timestamp) and output string to syslog and stderr.

import (
	"fmt"
	"log"
	"runtime"
)

//	the Codelearning.platform error type
type ClpError struct {
	Code     uint16
	Msg      string
	Location string
}

func (e *ClpError) Error() string {
	return fmt.Sprintf("%s: %d - %s", e.Location, e.Code, e.Msg)
}

//	Exported functions

//	Checks whetner logger works (can write to the log sink: syslog, stderr) or not.
func Check() error {
	//	TO-DO: replace with the actual activity

	/*
		location, _ := Get_function_name()
		return &ClpError{1, "Check", location}
	*/
	return nil
}

//	TO-DO:	think about making interface like the 'Info' function has.
func Debug(err error) error {
	//	TO-DO: replace with the actual activity
	fmt.Printf("[DEBUG] %s\n", err)

	return nil
}

func Error(err error) error {
	//	TO-DO: replace with the actual activity
	fmt.Printf("[ERROR] %s\n", err)

	return nil
}

func Info(format string, args ...interface{}) error {
	//	TO-DO: replace with the actual activity
	data := fmt.Sprintf(format, args...)
	if location, func_err := get_function_name(3); func_err == nil {
		fmt.Printf("[INFO] %s: %s\n", location, data)
	} else {
		//	TO-DO: Add cheking of debug flag
		Debug(func_err)
		fmt.Printf("[INFO] %s\n", data)
	}

	return nil
}

//	Returns name of function which calls this one.
func Get_function_name() (string, error) {

	program_counters := make([]uintptr, 1)
	amount_of_functions_in_the_callstack := runtime.Callers(2, program_counters)
	if amount_of_functions_in_the_callstack == 0 {
		return "", &ClpError{2, "There are no program counters", "Get_function_name()"}
	}

	frames := runtime.CallersFrames(program_counters[:amount_of_functions_in_the_callstack])
	frame, _ := frames.Next()

	return frame.Function, nil
}

//	Returns name of function which calls this one or an error instead.
func get_function_name(depth int) (string, error) {

	if depth <= 0 {
		return "", &ClpError{1, "Depth should be greater than 0", "get_function_name(depth int)"}
	}

	program_counters := make([]uintptr, 1)
	amount_of_functions_in_the_callstack := runtime.Callers(depth, program_counters)
	if amount_of_functions_in_the_callstack == 0 {
		return "", &ClpError{2, "There are no program counters. Check value of the depth argument", "get_function_name(depth int)"}
	}

	frames := runtime.CallersFrames(program_counters[:amount_of_functions_in_the_callstack])
	frame, _ := frames.Next()

	return frame.Function, nil
}

//	Static functions

func to_syslog(err error) error {

	//	TO-DO: replace with the actual activity
	if location, func_err := get_function_name(3); func_err == nil {
		fmt.Printf("%s: %s\n", location, err)
	} else {
		//	TO-DO: Add checking of debug flag
		Debug(func_err)
		fmt.Printf("%s\n", err)
	}

	return nil
}

func to_stderr(err error) error {

	//	Prints to stderr.
	if location, func_err := get_function_name(3); func_err == nil {
		log.Printf("%s: %s\n", location, err)
	} else {
		//	TO-DO: Add checking of debug flag
		Debug(func_err)
		log.Printf("%s\n", err)
	}

	return nil
}
