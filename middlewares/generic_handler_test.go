package middlewares

import (
	"testing"
)

func TestGenericMiddleware(t *testing.T) {
	print(generateRequestId())
}
