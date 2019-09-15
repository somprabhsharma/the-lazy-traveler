package logger

import (
	"encoding/json"
	"log"
	"runtime"
	"strings"
	"time"
)

//logObj is struct with all the log info
type logObj struct {
	Timestamp int64  `json:"timestamp"`
	LogLevel  string `json:"logLevel"`
	Class     string `json:"class"`
	UserID    string `json:"userId"`
	Message   string `json:"msg"`
	Error     string `json:"error,omitempty"`
}

//Info just logs with logLevel INFO
func Info(userID, message string) {
	logLine(userID, message, nil, "INFO")
}

//Err just logs with logLevel ERROR with error
func Err(userID, message string, err error) {
	logLine(userID, message, err, "ERROR")
}

//Warn just logs with logLevel WARNING with error
func Warn(userID, message string, err error) {
	logLine(userID, message, err, "WARNING")
}

var logLine = func(userID, message string, err error, logLevel string) {
	_, className, _, _ := runtime.Caller(2)
	parts := strings.Split(className, "/")
	part := parts[len(parts)-1]
	arr := strings.Split(part, ".")
	class := arr[0]
	logObj := logObj{
		Timestamp: getCurrentTime(),
		UserID:    userID,
		Message:   message,
		Class:     class,
		LogLevel:  "INFO",
	}
	if err != nil {
		logObj.Error = err.Error()
		logObj.LogLevel = logLevel
	}
	json, _ := json.Marshal(logObj)
	log.Println(string(json))
}

//getCurrentTime return current time in millis
func getCurrentTime() int64 {
	return time.Now().UnixNano() * int64(time.Nanosecond) / int64(time.Millisecond)
}
