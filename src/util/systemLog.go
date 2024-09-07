package util

import (
	"errors"

	"github.com/sirupsen/logrus"
)

type LogFields struct {
	ProjectName  string
	Message      string
	ErrorMessage string
}

// Print log to terminal, lowest level
func (c *LogFields) Trace() {
	logrus.WithFields(logrus.Fields{
		"error message": c.ErrorMessage,
		"project name":  c.ProjectName,
	}).Trace(c.Message)
}

// Print log for a little information
func (c *LogFields) Info() {
	logrus.WithField(c.ErrorMessage, logrus.FieldMap{
		"project name": c.ProjectName,
	}).Info(c.Message)
}

// Print log to terminal, use for debugging info.
func (c *LogFields) Debug() {
	logrus.WithFields(logrus.Fields{
		"error message": c.ErrorMessage,
		"project name":  c.ProjectName,
	}).Debug(c.Message)
}

// Print log to terminal, warning level
func (c *LogFields) Warn() {
	logrus.WithFields(logrus.Fields{
		"error message": c.ErrorMessage,
		"project name":  c.ProjectName,
	}).Warn(c.Message)
}

// Print log for a error information, but not stop the app
func (c *LogFields) Error() {
	logrus.WithFields(logrus.Fields{
		"error message": c.ErrorMessage,
		"project name":  c.ProjectName,
	}).Error(c.Message)
}

// Print log to terminal and send panic() can handle with recover())
func (c *LogFields) Panic() {
	logrus.WithFields(logrus.Fields{
		"error message": c.ErrorMessage,
		"project name":  c.ProjectName,
	}).Panic(c.Message)
}

// Print log to terminal and terminate the app (os.Exit(1))
func (c *LogFields) Fatal() {
	logrus.WithFields(logrus.Fields{
		"error message": c.ErrorMessage,
		"project name":  c.ProjectName,
	}).Fatal(c.Message)
}

func SystemLog(projectName, message string, errMsg error) *LogFields {

	if errMsg == nil {
		errMsg = errors.New("Additional Info: ")
	}

	return &LogFields{
		Message:      message,
		ErrorMessage: errMsg.Error(),
		ProjectName:  projectName,
	}
}
