package common

import (
	"log"
	"runtime/debug"
)

func LogErr(err error) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if err != nil {
		debug.PrintStack()
		log.Fatal(err)
	}
}
