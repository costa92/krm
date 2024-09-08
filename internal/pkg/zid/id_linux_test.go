//go:build linux
// +build linux

package zid

func TestMachineID(t *testing.T) {
	t.Log(machineID())
}
