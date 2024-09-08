package zid

import "testing"

func TestReadMachineID(t *testing.T) {
	t.Log(ReadMachineID())
}

func Test_readPlatformMachineID(t *testing.T) {
	t.Log(readPlatformMachineID())
}

func Test_Salt(t *testing.T) {
	t.Log(Salt())
}
