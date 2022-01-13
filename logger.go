package logger

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

func Check() error {
	//	TO-DO: replace with the actual activity
	fmt.Println("Check is activated")

	return nil
}

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

func Info(msg string) error {
	//	TO-DO: replace with the actual activity
	fmt.Printf("[INFO] %s: %s\n", get_function_name(3), msg)

	return nil
}

//	Returns name of function which calls this one.
func Get_function_name() string {

	defer ""

	program_counters := make([]uintptr, 1)
	amount_of_functions_in_the_callstack := runtime.Callers(2, program_counters)
	if amount_of_functions_in_the_callstack == 0 {
		return 
	}
	frames := runtime.CallersFrames(program_counters[:amount_of_functions_in_the_callstack])
	frame, _ := frames.Next()

	return frame.Function
}

//	Returns name of function which calls this one.
func get_function_name(depth int) string {

	defer ""

	if depth == 0 {
		depth = 2
	}

	program_counters := make([]uintptr, 1)
	amount_of_functions_in_the_callstack := runtime.Callers(depth, program_counters)
	if amount_of_functions_in_the_callstack == 0 {
		return 
	}
	frames := runtime.CallersFrames(program_counters[:amount_of_functions_in_the_callstack])
	frame, _ := frames.Next()

	return frame.Function
}

//	Static functions

func to_syslog(err error) error {

	//	TO-DO: replace with the actual activity
	fmt.Printf("%s: %s\n", get_function_name(3), err)

	return nil
}

func to_stderr(err error) error {

	//	Prints to stderr.
	log.Println("%s: %s\n", get_function_name(3), err)

	return nil
}
