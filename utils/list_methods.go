package utils

import "reflect"

// ListMethods list all method by given struct
func ListMethods(input interface{}) (result []string) {
	inputType := reflect.TypeOf(input)
	for i := 0; i < inputType.NumMethod(); i++ {
		result = append(result, inputType.Method(i).Name)
	}

	return result
}

// ListFields list all fields by given struct
func ListFields(input interface{}) (result []string) {
	inputType := reflect.TypeOf(input)
	for i := 0; i < inputType.NumField(); i++ {
		result = append(result, inputType.Field(i).Name)
	}

	return result
}
