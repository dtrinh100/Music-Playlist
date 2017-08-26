package common

import "log"

func Fatal(err error, msg string) {
	if err != nil {
		log.Fatal(msg+":", err)
	}
}
