package version

import (
	"fmt"
	"runtime/debug"
	"testing"
)

func Test_Version(t *testing.T) {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		t.Fatal("ReadBuildInfo failed")
	}

	for _, dep := range info.Deps {
		fmt.Println(dep.Path, dep.Version)
	}
}
