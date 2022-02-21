package logger

//	TO-DO:	use the 'log' package to format (add timestamp) and output string to syslog and stderr.
//	TO-DO:	use ClociConfiguration to set a level of logging (error, warning, info, debug)
//	To-DO:	use a specific format to chain errors

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

//	The Codelearning.platform error type
type ClpError struct {
	Code     uint16
	Msg      string
	Location string
}

//	Makes ClpError a standard error.
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

func Error(err error) {
	//	TO-DO: replace with the actual activity
	fmt.Printf("[ERROR] %s\n", err)

	//	TO-DO: call a function that gracefully finishes the service.
	if cloci_err, is_cloci_err := err.(*ClpError); is_cloci_err {
		os.Exit(int(cloci_err.Code))
	} else {
		os.Exit(1)
	}
}

//	TO-DO: Add cheking of warning flag
func Warning(format string, args ...interface{}) error {
	//	TO-DO: replace with the actual activity
	data := fmt.Sprintf(format, args...)
	if location, func_err := get_function_name(3); func_err == nil {
		fmt.Printf("[WARNING] %s: %s\n", location, data)
	} else {
		fmt.Printf("[WARNING] %s\n", data)
		return func_err
	}

	return nil
}

//	TO-DO: Add cheking of debug flag
func Debug(format string, args ...interface{}) error {
	//	TO-DO: replace with the actual activity
	data := fmt.Sprintf(format, args...)
	if location, func_err := get_function_name(3); func_err == nil {
		fmt.Printf("[DEBUG] %s: %s\n", location, data)
	} else {
		fmt.Printf("[DEBUG] %s\n", data)
		return func_err
	}

	return nil
}

//	TO-DO: Add cheking of info flag
func Info(format string, args ...interface{}) error {
	//	TO-DO: replace with the actual activity
	data := fmt.Sprintf(format, args...)
	if location, func_err := get_function_name(3); func_err == nil {
		fmt.Printf("[INFO] %s: %s\n", location, data)
	} else {
		Warning("%s", func_err)
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
		fmt.Printf("%s\n", err)
		return func_err
	}

	return nil
}

func to_stderr(err error) error {

	//	Prints to stderr.
	if location, func_err := get_function_name(3); func_err == nil {
		log.Printf("%s: %s\n", location, err)
	} else {
		log.Printf("%s\n", err)
		return func_err
	}

	return nil
}
