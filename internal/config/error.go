package config

import (
	"log"
	"os"
	"runtime"
)

func HandleError(err error, fatal bool) {
	if err != nil {
		pc, filename, line, _ := runtime.Caller(1)
		log.Printf("[error] in %s :: [%s :line %d] %v", runtime.FuncForPC(pc).Name(), filename, line, err)
	}

	if fatal {
		os.Exit(1)
	}
}
