package logger

import (
	"testing"
)

func TestLogger(t *testing.T) {
	SetLevel(INFO)
	SetFormatter(NewStandardFormatter())

	Info("x: %s", "hi")
	Info("%s")
	Info("xxx")
	Debug("xxx")
	Error("xxx")
}
