/*
@author: sk
@date: 2023/9/9
*/
package utils

import (
	"fmt"
)

func LogErr(format string, args ...any) {
	fmt.Printf("\x1b[31m"+format+"\n", args...)
}

func LogWarn(format string, args ...any) {
	fmt.Printf("\x1b[33m"+format+"\n", args...)
}

func LogInfo(format string, args ...any) {
	fmt.Printf("\x1b[32m"+format+"\n", args...)
}
