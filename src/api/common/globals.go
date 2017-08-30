package common

type Str2mapstr map[string](map[string]string)
type ErrMap map[string]string

type ErrorList struct {
	Errors ErrMap    `json:"errors"`
}

type ServerConfig struct {
	Address  string
	LogLevel int
}



