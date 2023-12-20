package util

import (
	"reflect"
	"strings"
)

func IsNotStruct(v any) bool {
	r := reflect.ValueOf(v)
	if r.Kind() == reflect.Interface || r.Kind() == reflect.Pointer {
		r = r.Elem()
	}
	return r.Kind() != reflect.Struct
}

func IsNotPointer(v any) bool {
	r := reflect.ValueOf(v)
	return r.Kind() != reflect.Pointer
}

func IsNilValueReflect(v reflect.Value) bool {
	if (v.Kind() == reflect.Interface ||
		v.Kind() == reflect.Pointer ||
		v.Kind() == reflect.Slice ||
		v.Kind() == reflect.Chan ||
		v.Kind() == reflect.Func ||
		v.Kind() == reflect.UnsafePointer ||
		v.Kind() == reflect.Map) && v.IsNil() {
		return true
	}
	return false
}

func IsZero(a any) bool {
	v := reflect.ValueOf(a)
	if v.Kind() == reflect.Pointer && !IsNilValueReflect(v) {
		v = v.Elem()
	}
	return v.Kind() == reflect.Invalid || v.IsZero() || IsNilValueReflect(v) || !v.CanInterface() ||
		(v.Kind() == reflect.Map && len(v.MapKeys()) == 0) ||
		(v.Kind() == reflect.Struct && v.NumField() == 0) ||
		(v.Kind() == reflect.Slice && v.Len() == 0) ||
		(v.Kind() == reflect.Array && v.Len() == 0)
}

func GetDatabaseNameByStruct(a any) string {
	var result string
	t := reflect.TypeOf(a)
	if t.Kind() == reflect.Interface || t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	for i := 0; i < t.NumField(); i++ {
		result = t.Field(i).Tag.Get("database")
		if len(result) != 0 {
			break
		}
	}
	return result
}

func GetDatabaseNameBySlice(a any) string {
	v := reflect.ValueOf(a)
	if v.Kind() == reflect.Pointer || v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	return GetDatabaseNameByStruct(reflect.New(v.Type().Elem()).Interface())
}

func GetCollectionNameByStruct(a any) string {
	var result string
	t := reflect.TypeOf(a)
	if t.Kind() == reflect.Interface || t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	for i := 0; i < t.NumField(); i++ {
		result = t.Field(i).Tag.Get("collection")
		if len(result) != 0 {
			break
		}
	}
	return result
}

func GetCollectionNameBySlice(a any) string {
	v := reflect.ValueOf(a)
	if v.Kind() == reflect.Pointer || v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	return GetCollectionNameByStruct(reflect.New(v.Type().Elem()).Interface())
}

func SetInsertedIdOnDocument(insertedId, a any) {
	rInsertedId := reflect.ValueOf(insertedId)
	v := reflect.ValueOf(a)
	vr := reflect.ValueOf(a)
	t := reflect.TypeOf(a)
	if v.Kind() == reflect.Pointer || v.Kind() == reflect.Interface {
		vr = v.Elem()
		t = t.Elem()
	}
	for i := 0; i < vr.NumField(); i++ {
		fieldValue := vr.Field(i)
		fieldStruct := t.Field(i)
		if fieldValue.Kind() == reflect.Pointer || fieldValue.Kind() == reflect.Interface {
			fieldValue = fieldValue.Elem()
		}
		if strings.Contains(fieldStruct.Tag.Get("bson"), "_id") &&
			rInsertedId.Kind() == fieldValue.Kind() {
			fieldValue.Set(rInsertedId)
		}
	}
}

func MinInt64(value, min int64) int64 {
	if value < min {
		return min
	}
	return value
}
