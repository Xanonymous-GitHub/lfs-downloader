package utils

import "strings"

func ExtractSpecificStringSliceFrom(raw, startIndicator, endIndicator string) *string {
	idxFind := strings.Index(raw, startIndicator)
	if idxFind == -1 {
		return nil
	}

	left := strings.LastIndex(raw[:idxFind], endIndicator)
	if left == -1 {
		return nil
	}

	right := strings.Index(raw[idxFind:], endIndicator)
	if right == -1 {
		return nil
	}

	return ToPointerType(raw[left : idxFind+right])
}
