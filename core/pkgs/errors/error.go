package errors

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// List of errors
var (
	ErrRecordNotFound = errors.New("record not found")
)

// Wrap error
func Wrap(err error, format string, args ...interface{}) error {
	message := fmt.Sprintf(format, args...)
	logrus.Warn(message, err)
	return errors.Wrap(err, message)
}

// New create a new error
func New(format string, args ...interface{}) error {
	message := fmt.Sprintf(format, args...)
	logrus.Warn(message)
	return errors.New(message)
}
