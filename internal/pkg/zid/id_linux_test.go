//go:build linux
// +build linux

package zid

import "testing"

func TestMachineID(t *testing.T) {
	t.Log(machineID())
}
