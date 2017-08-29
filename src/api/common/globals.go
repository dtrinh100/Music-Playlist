package common

type Str2mapstr map[string](map[string]string)
type ErrMap map[string]string

type ErrorList struct {
	Errors map[string]string    `json:"errors"`
}

type configuration struct {
	Server server
	DB     database
}

type server struct {
	Address  string
	LogLevel int
}

type database struct {
	Name      string
	UserTable table
}

type table struct {
	Name string
}
