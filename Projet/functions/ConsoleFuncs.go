package functions

import (
	"log"
	"os"
	"os/exec"
	"runtime"
)

var ShouldLogInfo = false

// GetArgs return the arguments given to the program
func GetArgs() []string {
	return os.Args[1:]
}

// ClearCmd clear the console
func ClearCmd() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			return
		}
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			return
		}
	}
}

// InfoPrintf print the given info message formatted with the given arguments
func InfoPrintf(format string, a ...interface{}) {
	if ShouldLogInfo {
		log.Printf("\033[1;34m[Info] :\033[0m "+format, a...)
	}
}

// InfoPrintln print the given info message with a new line
func InfoPrintln(s string) {
	if ShouldLogInfo {
		log.Println("\033[1;34m[Info] :\033[0m " + s)
	}
}

// ErrorPrintf print the given error message formatted with the given arguments
func ErrorPrintf(format string, a ...interface{}) {
	log.Printf("\033[1;31m[Error] :\033[0m "+format, a...)
}

// ErrorPrintln print the given error message with a new line
func ErrorPrintln(s string) {
	log.Println("\033[1;31m[Error] :\033[0m " + s)
}

// WarningPrintf print the given warning message formatted with the given arguments
func WarningPrintf(format string, a ...interface{}) {
	log.Printf("\033[1;33m[Warning] :\033[0m "+format, a...)
}

// WarningPrintln print the given warning message with a new line
func WarningPrintln(s string) {
	log.Println("\033[1;33m[Warning] :\033[0m " + s)
}

// SuccessPrintf print the given success message formatted with the given arguments
func SuccessPrintf(format string, a ...interface{}) {
	log.Printf("\033[1;32m[Success] :\033[0m "+format, a...)
}

// SuccessPrintln print the given success message with a new line
func SuccessPrintln(s string) {
	log.Println("\033[1;32m[Success] :\033[0m " + s)
}

// FatalPrintf print the given fatal message  formatted with the given arguments and exit the program
func FatalPrintf(format string, a ...interface{}) {
	log.Fatalf("\033[1;31m[Fatal] :\033[0m "+format, a...)
}

// FatalPrintln print the given fatal message with a new line and exit the program
func FatalPrintln(s string) {
	log.Fatalln("\033[1;31m[Fatal] :\033[0m " + s)
}
