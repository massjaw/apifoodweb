package util

import (
	"errors"

	"github.com/sirupsen/logrus"
)

type LogInterface interface {
	Trace()
	Info()
	Debug()
	Warn()
	Error()
	Panic()
	Fatal()
}
type LogFields struct {
	ProjectName  string
	Message      string
	ErrorMessage string
}

func SystemLog(projectName, message string, errMsg error) LogInterface {

	if errMsg == nil {
		errMsg = errors.New("Additional Info: ")
	}

	return &LogFields{
		Message:      message,
		ErrorMessage: errMsg.Error(),
		ProjectName:  projectName,
	}
}

// Print log to terminal, lowest level
func (c *LogFields) Trace() {
	if c.ErrorMessage != "" {

		logrus.WithFields(logrus.Fields{
			"project name": c.ProjectName,
		}).Trace(c.Message)
	} else {

		logrus.WithFields(logrus.Fields{
			"error message": c.ErrorMessage,
			"project name":  c.ProjectName,
		}).Trace(c.Message)
	}
}

// Print log for a little information
func (c *LogFields) Info() {
	if c.ErrorMessage != "" {

		logrus.WithFields(logrus.Fields{
			"project name": c.ProjectName,
		}).Info(c.Message)
	} else {

		logrus.WithFields(logrus.Fields{
			"error message": c.ErrorMessage,
			"project name":  c.ProjectName,
		}).Info(c.Message)
	}
}

// Print log to terminal, use for debugging info.
func (c *LogFields) Debug() {
	if c.ErrorMessage != "" {

		logrus.WithFields(logrus.Fields{
			"project name": c.ProjectName,
		}).Debug(c.Message)
	} else {

		logrus.WithFields(logrus.Fields{
			"error message": c.ErrorMessage,
			"project name":  c.ProjectName,
		}).Debug(c.Message)
	}
}

// Print log to terminal, warning level
func (c *LogFields) Warn() {
	if c.ErrorMessage != "" {

		logrus.WithFields(logrus.Fields{
			"project name": c.ProjectName,
		}).Warn(c.Message)
	} else {

		logrus.WithFields(logrus.Fields{
			"error message": c.ErrorMessage,
			"project name":  c.ProjectName,
		}).Warn(c.Message)
	}
}

// Print log for a error information, but not stop the app
func (c *LogFields) Error() {
	if c.ErrorMessage == "" {

		logrus.WithFields(logrus.Fields{
			"project name": c.ProjectName,
		}).Error(c.Message)
	} else {

		logrus.WithFields(logrus.Fields{
			"error message": c.ErrorMessage,
			"project name":  c.ProjectName,
		}).Error(c.Message)
	}
}

// Print log to terminal and send panic() can handle with recover())
func (c *LogFields) Panic() {
	if c.ErrorMessage != "" {

		logrus.WithFields(logrus.Fields{
			"project name": c.ProjectName,
		}).Panic(c.Message)
	} else {

		logrus.WithFields(logrus.Fields{
			"error message": c.ErrorMessage,
			"project name":  c.ProjectName,
		}).Panic(c.Message)
	}
}

// Print log to terminal and terminate the app (os.Exit(1))
func (c *LogFields) Fatal() {
	if c.ErrorMessage != "" {

		logrus.WithFields(logrus.Fields{
			"project name": c.ProjectName,
		}).Fatal(c.Message)
	} else {

		logrus.WithFields(logrus.Fields{
			"error message": c.ErrorMessage,
			"project name":  c.ProjectName,
		}).Fatal(c.Message)
	}
}
