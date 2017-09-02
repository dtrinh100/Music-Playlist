package common

import "log"

// Fatal helps write cleaner code by packaging the if-statement into a function.
func Fatal(err error, msg string) {
	if err != nil {
		log.Fatal(msg+":", err)
	}
}
