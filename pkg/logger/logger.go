package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

var (
	// AppName - set app name for logging identifier
	AppName = ""
)

// Fields - to pass extra fields value in logs.
type Fields map[string]interface{}

// Init - Init Logger Config
func Init(debug bool) {
	// Set Log Formatter to JSON formatter
	log.SetFormatter(&log.JSONFormatter{})
	// Set Logger Output
	log.SetOutput(os.Stdout)
	// Set Log Level -  Debugging is enable/disable
	if debug {
		log.SetLevel(log.DebugLevel) // Debug Level +
	} else {
		log.SetLevel(log.InfoLevel) // Warn Level +
	}
}

// Debug - Custom Debug level wrapper for debug log
func Debug(message string, other ...map[string]interface{}) {
	fields := map[string]interface{}{}
	if other != nil {
		fields = other[0]
	}
	fields["message"], fields["name"] = message, AppName
	log.WithFields(fields).Debug("")
}

// Info - Custom Info level wrapper for info level log
func Info(message string, other ...map[string]interface{}) {
	fields := map[string]interface{}{}
	if other != nil {
		fields = other[0]
	}
	fields["name"] = AppName
	log.WithFields(fields).Info(message)
}

// Error - Custom Error level wrapper for error level log
func Error(err error, message string, other ...map[string]interface{}) {
	fields := map[string]interface{}{}
	if other != nil {
		fields = other[0]
	}
	fields["name"] = AppName
	log.WithFields(fields).WithError(err).Error(message)
}
