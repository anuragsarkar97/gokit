package log

import "log"

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"

func Info(query string, val ...interface{}) {
	log.Printf(Green+"[INFO] "+Reset+query, val...)
}

func Error(query string, val ...interface{}) {
	log.Printf(Red+"[ERROR] "+Reset+query, val...)
}

func Debug(query string, val ...interface{}) {
	log.Printf(Yellow+"[DEBUG] "+Reset+query, val...)
}

func Warn(query string, val ...interface{}) {
	log.Printf(Blue+"[WARN] "+Reset+query, val...)
}
