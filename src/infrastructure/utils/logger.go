package utils

import "log"

type Logger struct {
	isEnabled bool
}

func (l *Logger) LogInfo(message string) {
	if !l.isEnabled {
		return
	}
	log.Println(message)
}

func NewLogger(isEnabled bool) *Logger {
	return &Logger{
		isEnabled: isEnabled,
	}
}
