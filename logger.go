package frogproxy

type Logger interface {
	Printf(format string, v ...interface{})
}
