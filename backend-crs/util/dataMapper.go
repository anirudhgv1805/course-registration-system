package util

import (
	"backend-crs/model"
	"fmt"
)

func FilterAllowedFieldsWithTypes(allowed map[string]any, newData map[string]any) (map[string]any, error) {

	filteredData := make(map[string]any)

	for k, v := range newData {
		expectedType, ok := allowed[k]
		if !ok {
			continue
		}
		switch expectedType {
		case "string":
			if str, ok := v.(string); ok {
				filteredData[k] = str
			} else {
				return nil, fmt.Errorf("unexpected type for key '%s': expected (string), got: %T ", k, v)
			}
		case "int":
			if num, ok := v.(int); ok {
				filteredData[k] = num
			} else if num, ok := v.(float64); ok {
				filteredData[k] = num
			} else {
				return nil, fmt.Errorf("unexpected type for key '%s': expected (int), got: %T ", k, v)
			}
		case "uint":
			switch num := v.(type) {
			case float64:
				if num < 0 || num != float64(uint(num)) {
					return nil, fmt.Errorf("invalid value for key '%s': expected unsigned integer, got %v", k, num)
				}
				filteredData[k] = uint(num)
			case uint:
				filteredData[k] = num
			case int:
				if num < 0 {
					return nil, fmt.Errorf("invalid value for key '%s': negative integer is not allowed for uint", k)
				}
				filteredData[k] = uint(num)
			default:
				return nil, fmt.Errorf("invalid type for key '%s': expected uint, got %T", k, v)
			}
		case "Staff":
			if staff, ok := v.(model.Staff); ok {
				filteredData[k] = staff
			} else {
				return nil, fmt.Errorf("invalid type for key '%s': expected Staff, got %T", k, v)
			}

		}
	}
	return filteredData, nil
}
