package utils

import "encoding/json"

func FlattenedFrom(origin any) (map[string]any, error) {
	serializedProfile, err := json.Marshal(origin)
	if err != nil {
		return nil, err
	}

	var flattenedProfile map[string]any
	err = json.Unmarshal(serializedProfile, &flattenedProfile)
	if err != nil {
		return nil, err
	}

	return flattenedProfile, nil
}
