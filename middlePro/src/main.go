package main

import (
	"github.com/gookit/color"
	"github.com/pengk/summer/bootstrap"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			color.Error.Println("[Start Server Err!!!] ", err)
		}
	}()
	bootstrap.Start()
}
