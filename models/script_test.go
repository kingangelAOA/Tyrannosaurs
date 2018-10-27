package models

import (
	"testing"
	"fmt"
	"os/exec"
)

func TestScript(t *testing.T)  {
	cmd := exec.Command("python",  "-c", "print('aaaa')\nprint('dddd')")
	fmt.Println(cmd.Args)
	out, err := cmd.CombinedOutput()
	if err != nil { fmt.Println(err); }
	fmt.Println(string(out))
}
