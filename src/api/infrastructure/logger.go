package infrastructure

import "fmt"

type Logger struct {}

// TODO: Update implementation
func (logger *Logger) Log(message string) error {
	fmt.Println(message)
	return nil
}
