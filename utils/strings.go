package utils

import "github.com/fatih/color"

func BoldString(format string, a ...interface{}) string {
	return color.New(color.Bold).Sprintf(format, a...)
}
