package berr

import (
	"fmt"
	"testing"
)

func TestCause(t *testing.T) {
	tests := []struct {
		name string
		err  error
		code Code
		want bool
	}{
		{
			name: "Failure",
			err:  Wrap(New("test", Unmarshal, "error"), "testWrap", InvalidArgument, "error"),
			code: Marshal,
			want: false,
		},
		{
			name: "Success-Level1",
			err:  Wrap(New("test", Unmarshal, "error"), "testWrap", InvalidArgument, "error"),
			code: InvalidArgument,
			want: true,
		},
		{
			name: "Success-Level2",
			err:  Wrap(New("test", Unmarshal, "error"), "testWrap", InvalidArgument, "error"),
			code: Unmarshal,
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Cause(tt.err, tt.code); got != tt.want {
				t.Errorf("Cause() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrorCode(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want Code
	}{
		{"Not-bError", fmt.Errorf("some error"), ""},
		{"bError", New("test", InvalidArgument, "some error"), InvalidArgument},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ErrorCode(tt.err); got != tt.want {
				t.Errorf("ErrorCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrorMessage(t *testing.T) {
	wrappedError := Wrap(New("test", Unmarshal, "error"), "testWrap", InvalidArgument, "wrapped error")
	tests := []struct {
		name string
		err  error
		want string
	}{
		{"Not-bError", fmt.Errorf("some error"), "some error"},
		{"bError-Level1", New("test", InvalidArgument, "some error"), "some error"},
		{"bError-Level2", wrappedError, "wrapped error"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ErrorMessage(tt.err); got != tt.want {
				t.Errorf("ErrorMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrorLog(t *testing.T) {
	wrappedError := Wrap(New("test", Unmarshal, "error"), "testWrap", InvalidArgument, "wrapped error")
	tests := []struct {
		name string
		err  error
		want string
	}{
		{"Not-bError", fmt.Errorf("some error"), "[\"some error\"]"},
		{"bError-Level1", New("test", InvalidArgument, "some error"), "[\"test: InvalidArgument - some error\"]"},
		{"bError-Level2", wrappedError, "[\"testWrap: InvalidArgument - wrapped error\",\"test: UnmarshalError - error\"]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ErrorLog(tt.err); got != tt.want {
				t.Errorf("ErrorLog() = %v, want %v", got, tt.want)
			}
		})
	}
}
