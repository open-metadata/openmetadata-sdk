package ometa

import (
	"encoding/json"
	"fmt"
	"reflect"
)


func Deserialize(jsonResponse map[string]any, targetType any) (any, error) {
	if reflect.TypeOf(targetType).Kind() != reflect.Ptr {
		return nil, fmt.Errorf("targetType must be a pointer to a struct")
	}

	jsonBytes, err := json.Marshal(jsonResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JSON response: %w", err)
	}

	runTimeType := reflect.New(reflect.TypeOf(targetType).Elem()).Interface() // Create a new instance of the target type

	if err := json.Unmarshal(jsonBytes, runTimeType); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON into target type: %w", err)
	}

	return runTimeType, nil
}

func DeserializeSlice(jsonArray []any, targetType any) (any, error) {
	if reflect.TypeOf(targetType).Kind() != reflect.Ptr {
		return nil, fmt.Errorf("targetType must be a pointer to a struct")
	}

	jsonBytes, err := json.Marshal(jsonArray)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JSON array: %w", err)
	}

	sliceType := reflect.SliceOf(reflect.TypeOf(targetType).Elem())
	resultSlice := reflect.New(sliceType)
	resultSlice.Elem().Set(reflect.MakeSlice(sliceType, 0, 0))

	if err := json.Unmarshal(jsonBytes, resultSlice.Interface()); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON into target slice: %w", err)
	}

	return resultSlice.Elem().Interface(), nil
}

func Serialize(target any) (map[string]any, error) {
	jsonBytes, err := json.Marshal(target)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal target: %w", err)
	}
	var result map[string]any
	if err := json.Unmarshal(jsonBytes, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON into map[string]any: %w", err)
	}
	return result, nil
}

func SerializeSlice(target any) ([]map[string]any, error) {
	jsonBytes, err := json.Marshal(target)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal target: %w", err)
	}
	var result []map[string]any
	if err := json.Unmarshal(jsonBytes, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON into []map[string]any: %w", err)
	}
	return result, nil
}
