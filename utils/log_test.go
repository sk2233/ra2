/*
@author: sk
@date: 2023/9/9
*/
package utils

import (
	"fmt"
	"os"
	"testing"

	"github.com/fatih/color"
)

func TestLog(t *testing.T) {
	LogErr("test err")
	LogWarn("test warn")
	LogInfo("test info")
	const escape = "\x1b"
	f := fmt.Sprintf("%s[%sm", escape, "34")
	fmt.Fprint(os.Stdout, f)
	fmt.Fprintln(os.Stdout, "Hello World in blue")
	fmt.Fprintln(os.Stdout, "Hello World in blue")
	fmt.Println("\x1b[34msfsfsdf")
	color.Blue("test")
}
