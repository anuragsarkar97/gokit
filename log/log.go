package log

import "log"

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"

var infoLog = Green + "[INFO] " + Reset
var warnLog = Yellow + "[WARN] " + Reset
var errorLog = Red + "[ERROR] " + Reset
var debugLog = Blue + "[DEBUG] " + Reset

func Info(query string, val ...interface{}) {
	log.Printf(infoLog+query, val...)
}

func Error(query string, val ...interface{}) {
	log.Printf(errorLog+query, val...)
}

func Debug(query string, val ...interface{}) {
	log.Printf(debugLog+query, val...)
}

func Warn(query string, val ...interface{}) {
	log.Printf(warnLog+query, val...)
}
