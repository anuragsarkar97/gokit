package berr

import (
	"fmt"
	"strings"
)

type Code string

type bError struct {
	Op      string
	Code    Code
	Message string
	Err     error
}

func New(op string, code Code, msg string, arg ...interface{}) error {
	return bError{Op: op, Code: code, Message: fmt.Sprintf(msg, arg...)}
}

func (be bError) Error() string {
	return fmt.Sprintf("%s: %s - %s", be.Op, be.Code, be.Message)
}

func (be bError) String() string {
	return be.Error()
}

func Wrap(err error, op string, code Code, msg string, arg ...interface{}) error {
	return bError{Op: op, Code: code, Message: fmt.Sprintf(msg, arg...), Err: err}
}

// Compares the error code recursively.
func Cause(err error, code Code) bool {
	if err == nil {
		return false
	}
	if e, ok := err.(bError); ok {
		if e.Code == code {
			return true
		}
		return Cause(e.Err, code)
	}
	return false
}

func ErrorCode(err error) Code {
	if err == nil {
		return ""
	}
	if e, ok := err.(bError); ok {
		return e.Code
	}
	return ""
}

func ErrorMessage(err error) string {
	if err == nil {
		return ""
	}
	if e, ok := err.(bError); ok {
		return e.Message
	}
	return err.Error()
}

func ErrorLog(err error) string {
	var errMsg []string
	for {
		if err == nil {
			break
		}
		if e, ok := err.(bError); ok {
			errMsg = append(errMsg, e.Error())
			err = e.Err
		} else {
			errMsg = append(errMsg, err.Error())
			break
		}
	}
	return fmt.Sprintf("[\"%s\"]", strings.Join(errMsg, "\",\""))
}
