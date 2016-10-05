package cmd

// Illustrates a bug I found in os/exec (I think) where an argument
//	containing curly braces is handled with any matching pairs of 
//	braces removed
//
// originally discovered/reported on github/hub:
//	https://github.com/github/hub/issues/1268
// ref:
//	https://golang.org/pkg/os/exec	

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"reflect"
)

// Spawn runs command with spawn(3)
func Spawn() error {
	//originally pulled from hub cmd/cmd.go Spawn()
	my_args := []string{"stash", "show", "stash@{0}"}

	c := exec.Command("git", my_args...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	fmt.Println(c.Args)

	result := c.Run()
	fmt.Println(reflect.TypeOf(result))

	return result
}
