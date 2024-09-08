package zid

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"hash/fnv"
	"os"
)

func Salt() uint64 {
	// Calculate the hash value of the string using the FNV-1a hash algorithm
	h := fnv.New64a()
	h.Write(ReadMachineID())

	// Convert the hash value to a salt of type uint64
	hash := h.Sum64()
	return hash
}

func ReadMachineID() []byte {
	id := make([]byte, 3)
	hid, err := readPlatformMachineID()
	if err != nil || len(hid) == 0 {
		hid, err = os.Hostname()
	}
	if err == nil && len(hid) != 0 {
		hw := sha256.New()
		hw.Write([]byte(hid))
		copy(id, hw.Sum(nil))
	} else {
		// Fallback to rand number if machine id can't be gathered
		if _, randErr := rand.Reader.Read(id); randErr != nil {
			panic(fmt.Errorf("id: cannot get hostname nor generate a random number: %w; %w", err, randErr))
		}
	}
	return id
}

// ReadMachineID reads the machine ID from the platform
// zh: 从平台读取机器ID
func readPlatformMachineID() (string, error) {
	// https://github.com/denisbrodbeck/machineid
	return machineID()
}
