package errors

import "fmt"

// MinerStatusErrorPtr converts a MinerStatusError to a pointer.
type MinerError struct {
	Reason  MinerStatusError
	Message string
}

func (e *MinerError) Error() string {
	return e.Message
}

// Some error builders for ease of use. They set the appropriate "Reason"
// value, and all arguments are Printf-style varargs fed into Sprintf to
// construct the Message.

// InvalidMinerConfiguration creates a new error when a Miner has invalid configuration.
func InvalidMinerConfiguration(msg string, args ...any) *MinerError {
	return &MinerError{
		Reason:  InvalidConfigurationMinerError,
		Message: fmt.Sprintf(msg, args...),
	}
}

// CreateMiner creates a new error for when creating a Miner.
func CreateMiner(msg string, args ...any) *MinerError {
	return &MinerError{
		Reason:  CreateMinerError,
		Message: fmt.Sprintf(msg, args...),
	}
}

// UpdateMiner creates a new error for when updating a Miner.
func UpdateMiner(msg string, args ...any) *MinerError {
	return &MinerError{
		Reason:  UpdateMinerError,
		Message: fmt.Sprintf(msg, args...),
	}
}

// DeleteMiner creates a new error for when deleting a Miner.
func DeleteMiner(msg string, args ...any) *MinerError {
	return &MinerError{
		Reason:  DeleteMinerError,
		Message: fmt.Sprintf(msg, args...),
	}
}
