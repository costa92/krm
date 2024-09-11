//go:build linux
// +build linux

package zid

import "os"

const (
	// dbusPathEtc is the default path for dbus machine id located in /etc.
	// Some systems (like Fedora 20) only know this path.
	// Sometimes it's the other way round.
	dbusPathEtc = "/etc/machine-id"
	// dbusPath is the default path for dbus machine id.
	//dbusPath = "/var/lib/dbus/machine-id"
	dbusPath = "/sys/class/dmi/id/product_uuid"
)

// https://github.com/denisbrodbeck/machineid/blob/master/id_linux.go
func machineID() (string, error) {
	b, err := os.ReadFile(dbusPathEtc)
	if err != nil || len(b) == 0 {
		b, err = os.ReadFile(dbusPath)
	}
	return string(b), err
}
