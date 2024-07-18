package utils

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

const ExcelTagName = "excel"

func BuildHeaderMap(headers []string) (map[int]string, error) {
	if IsRowEmpty(headers) {
		return nil, errors.New("headers row is empty")
	}
	columnIdToName := make(map[int]string)
	for i := 0; i < len(headers); i++ {
		columnIdToName[i] = strings.TrimSpace(headers[i])
	}
	return columnIdToName, nil
}

func BuildFieldNameByTagMap(item interface{}) map[string]string {
	e := reflect.ValueOf(item).Elem()
	tagToFieldName := make(map[string]string)
	for i := 0; i < e.NumField(); i++ {
		varName := e.Type().Field(i).Name
		varTag := e.Type().Field(i).Tag
		tagToFieldName[varTag.Get(ExcelTagName)] = varName
	}

	return tagToFieldName
}

func SetValueToField(field reflect.Value, value string) error {
	if isColumnEmpty(value) {
		return nil
	}
	switch field.Kind() {
	case reflect.Int64:
		x, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		field.SetInt(x)

	case reflect.String:
		field.SetString(value)
	case reflect.Float64:
		x, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		field.SetFloat(x)
	case reflect.Slice:
		return setValueToSlice(field, value)

	case reflect.Ptr:
		return setValueToPointer(field, value)
	}
	return nil
}

func setValueToPointer(field reflect.Value, value string) error {
	switch field.Type().Elem().Kind() {
	case reflect.Int64:
		x, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		field.Set(reflect.ValueOf(&x))

	case reflect.String:
		field.Set(reflect.ValueOf(&value))
	case reflect.Float64:
		x, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		field.Set(reflect.ValueOf(&x))
	case reflect.Slice:
		return setValueToSlice(field, value)
	}

	return nil
}

func setValueToSlice(field reflect.Value, value string) error {
	data := strings.Split(value, ",")
	switch field.Type().Elem().Kind() {
	case reflect.Int64:
		var arr []int64
		for i := 0; i < len(data); i++ {
			parseInt, err := strconv.ParseInt(data[i], 10, 64)
			if err != nil {
				return err
			}
			arr = append(arr, parseInt)
		}
		field.Set(reflect.ValueOf(arr))
	case reflect.String:
		field.Set(reflect.ValueOf(data))
	case reflect.Float64:
		var arr []float64
		for i := 0; i < len(data); i++ {
			parseInt, err := strconv.ParseFloat(data[i], 64)
			if err != nil {
				return err
			}
			arr = append(arr, parseInt)
		}
		field.Set(reflect.ValueOf(arr))
	}
	return nil
}

func isColumnEmpty(column string) bool {
	column = strings.TrimSpace(column)
	return column == ""
}

func IsRowEmpty(row []string) bool {
	columnEmptyCount := 0
	for _, column := range row {
		if isColumnEmpty(column) {
			columnEmptyCount++
		}
	}
	return columnEmptyCount == len(row)
}
