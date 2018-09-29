package util

import (
	"testing"
	"go/importer"
	"fmt"
	//"time"
)

func TestReflect(t *testing.T) {
	pkg, err := importer.Default().Import("tyrannosaurs/util")
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return
	}
	fmt.Println(pkg.Path())
	for _, declName := range pkg.Scope().Names() {
		fmt.Println(declName)
	}
}