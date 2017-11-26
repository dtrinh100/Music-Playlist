package usecases

type FaultError struct {
	FaultEntity int
	Message     string
}

func (err *FaultError) Status() int {
	return err.FaultEntity
}

func (err *FaultError) Error() string {
	return err.Message
}

