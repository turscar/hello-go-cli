package main

import (
	"fmt"
	"github.com/fatih/color"
)

func main() {
	_, _ = fmt.Fprintf(color.Output, "%s, %s!\n", color.GreenString("Hello"), color.BlueString(world()))
}

func world() string {
	return "World"
}
