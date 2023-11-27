package util

import (
	"errors"
	"strconv"
)

// ParseIDFromString takes a string and attempts to parse it into an integer ID.
func ParseIDFromString(idStr string) (int, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, errors.New("invalid ID format")
	}
	return id, nil
}

// ValidateUrgency checks if the given urgency value is valid (either "u" or "nu").
func ValidateUrgency(urgency string) bool {
	return urgency == "u" || urgency == "nu"
}

// ValidateImportance checks if the given importance value is valid (either "i" or "ni").
func ValidateImportance(importance string) bool {
	return importance == "i" || importance == "ni"
}

// Other utility functions can be added here as needed.
