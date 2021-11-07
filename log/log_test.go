package log

import (
	"testing"
)

func TestInfo(t *testing.T) {
	Info("haha %s %s", "world", "!")
	Error("haha %s %s", "world", "!")
	Debug("haha %s %s", "world", "!")
	Warn("haha %s %s", "world", "!")
}
