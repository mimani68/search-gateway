package logHandler

import "fmt"

type LogConfig struct {
	Type string
}

func Log(message string) {
	fmt.Println(message)
}
