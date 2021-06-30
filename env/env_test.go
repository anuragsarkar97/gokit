package env

import (
	"os"
	"reflect"
	"strings"
	"testing"
	"time"
)

func Test_getEnvAsString(t *testing.T) {
	os.Clearenv()
	const property = "ENV_TEST"
	t.Run("EnvNotSet", func(t *testing.T) {
		if got := AsString(property, "default"); got != "default" {
			t.Errorf("got=%s, want=%s", got, "default")
		}
	})
	t.Run("EnvSet", func(t *testing.T) {
		_ = os.Setenv(property, "test")
		if got := AsString(property, "default"); got != "test" {
			t.Errorf("got=%s, want=%s", got, "test")
		}
	})
}

func Test_getEnvAsInt(t *testing.T) {
	os.Clearenv()
	const property = "ENV_TEST"
	t.Run("EnvNotSet", func(t *testing.T) {
		if got := AsInt(property, 2); got != 2 {
			t.Errorf("got=%d, want=%d", got, 2)
		}
	})
	t.Run("EnvSet", func(t *testing.T) {
		_ = os.Setenv(property, "1")
		if got := AsInt(property, 2); got != 1 {
			t.Errorf("got=%d, want=%d", got, 1)
		}
	})
}

func Test_getEnvAsBool(t *testing.T) {
	os.Clearenv()
	const property = "ENV_TEST"
	t.Run("EnvNotSet", func(t *testing.T) {
		if got := AsBool(property, false); got != false {
			t.Errorf("got=%v, want=%v", got, false)
		}
	})
	t.Run("EnvSet", func(t *testing.T) {
		_ = os.Setenv(property, "true")
		if got := AsBool(property, false); got != true {
			t.Errorf("got=%v, want=%v", got, true)
		}
	})
}

func Test_getEnvAsMillisecondDuration(t *testing.T) {
	os.Clearenv()
	defaultValue := time.Millisecond * 5
	const property = "ENV_TEST"
	t.Run("EnvNotSet", func(t *testing.T) {
		if got := AsMillisecondDuration(property, defaultValue); got != defaultValue {
			t.Errorf("got=%v, want=%v", got, defaultValue)
		}
	})
	t.Run("EnvSet", func(t *testing.T) {
		want := time.Millisecond * 4
		_ = os.Setenv(property, "4")
		if got := AsMillisecondDuration(property, defaultValue); got != want {
			t.Errorf("got=%v, want=%v", got, want)
		}
	})
}

func Test_getEnvAsSecondDuration(t *testing.T) {
	os.Clearenv()
	defaultValue := time.Second * 5
	const property = "ENV_TEST"
	t.Run("EnvNotSet", func(t *testing.T) {
		if got := AsSecondDuration(property, defaultValue); got != defaultValue {
			t.Errorf("got=%v, want=%v", got, false)
		}
	})
	t.Run("EnvSet", func(t *testing.T) {
		want := time.Second * 4
		_ = os.Setenv(property, "4")
		if got := AsSecondDuration(property, defaultValue); got != want {
			t.Errorf("got=%v, want=%v", got, want)
		}
	})
}

func Test_getEnvAsStringArray(t *testing.T) {
	os.Clearenv()
	var defaultValue = []string{"one", "two"}
	const property = "ENV_TEST"
	t.Run("EnvNotSet", func(t *testing.T) {
		if got := AsStringArray(property, ",", defaultValue); !reflect.DeepEqual(got, defaultValue) {
			t.Errorf("got=%v, want=%v", got, defaultValue)
		}
	})
	t.Run("EnvSet", func(t *testing.T) {
		want := []string{"three", "four"}
		_ = os.Setenv(property, strings.Join(want, ","))
		if got := AsStringArray(property, ",", defaultValue); !reflect.DeepEqual(got, want) {
			t.Errorf("got=%v, want=%v", got, want)
		}
	})
}
