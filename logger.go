package gohonualogger

import (
	"fmt"
	"os"
	"time"
)

var (
	LEVEL string = os.Getenv("LOGLEVEL")
)

func Print(message string, debug bool) {
	if debug && !is_debug_active() {
		return
	}
	location, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		fmt.Printf("[ERROR]: %s\n", err.Error())
		return
	}
	currentTime := time.Now().In(location).Format("02/01/2006 15:04:05")
	var printableMessage string = ""
	if debug {
		printableMessage = fmt.Sprintf("%s - [DEBUG]: %s\n", currentTime, message)
	} else {
		printableMessage = fmt.Sprintf("%s - [INFO]: %s\n", currentTime, message)
	}
	fmt.Print(printableMessage)

	file, err := os.OpenFile("./honua-logs.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("[ERROR]: %s\n", err.Error())
		return
	}
	defer file.Close()

	file.WriteString(printableMessage)
}

func is_debug_active() bool {
	return LEVEL == "DEBUG"
}
