package main

import (
	"log"
	"runtime"
)

func doCaller() {
	programCounter, sourceFileName, sourceFileLineNum, ok := runtime.Caller(1)
	log.Printf("programCounter: %v\n", programCounter)
	log.Printf("filename: %s\n", sourceFileName)
	log.Printf("linenum: %d\n", sourceFileLineNum)
	log.Printf("ok: %t\n", ok)
}

func callCaller() {
	doCaller()
}

func main() {
	callCaller()
}
