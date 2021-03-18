package main

import (
	"errors"
	"log"
	"reflect"
)

func AssignStructFields(in interface{}, values map[string]interface{}) error {
	if reflect.ValueOf(values).Kind() != reflect.Map {
		return errors.New("Values is not a map")
	}
	if in == nil {
		return errors.New("struct is empty")
	}

	val := reflect.ValueOf(in)
	if val.Kind() != reflect.Ptr {
		return errors.New("We can treat only pointers to the struct")
	}

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return errors.New("this is not a struct")
	}

	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)

		if typeField.Type.Kind() == reflect.Struct {
			log.Printf("nested field: %v", typeField.Name)
			AssignStructFields(val.Field(i).Interface(), values)
			continue
		}

		if mapValue, ok := values[typeField.Name]; ok {
			newValue := reflect.ValueOf(mapValue)
			log.Println(newValue)
			val.Field(i).Set(newValue)
		}

		log.Printf("\tname=%s, type=%s, value=%v, tag=`%s`\n",
			typeField.Name,
			typeField.Type,
			val.Field(i),
			typeField.Tag,
		)

	}
	return nil
}
