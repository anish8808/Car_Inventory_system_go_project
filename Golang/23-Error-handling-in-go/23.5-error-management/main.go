package main

/*
	Best Practices for Error management

		1 - check for Errors Immediately (with each fun or data strctures check)
		2 - Return Errors Instead of panicking(not all the case)
		3 - Wrap Errors for Context
		4 - Use sentinel Errors Sparingly
		5 - leverage custom Error Types
		6 - use error.is and error.As for Error Inspection
		7 - Use defer with recover for panic recovery
		8 - Do not hide errors
		9 - Handel errors in Goroutines  , when launching goroutines , handle or propagate errors back to
			the main routine using channels
		10 - Return nil for for success
		11 - Avoid overuse of Panid
*/
