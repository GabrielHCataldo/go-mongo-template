package util

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func IsObjectId(v reflect.Value) bool {
	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}
	if _, ok := v.Interface().(primitive.ObjectID); ok {
		return true
	}
	return false
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
	return v.Kind() == reflect.Invalid || v.IsZero() || IsNilValueReflect(v) || !v.CanInterface() ||
		(v.Kind() == reflect.Map && len(v.MapKeys()) == 0) ||
		(v.Kind() == reflect.Struct && v.NumField() == 0) ||
		(v.Kind() == reflect.Slice && v.Len() == 0) ||
		(v.Kind() == reflect.Array && v.Len() == 0)
}

func GetDatabaseNameByStruct(a any) string {
	var result string
	v := reflect.ValueOf(a)
	t := reflect.TypeOf(a)
	if v.Kind() == reflect.Interface || v.Kind() == reflect.Pointer {
		v = v.Elem()
		t = t.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		result = t.Field(i).Tag.Get("database")
		if len(result) != 0 {
			break
		}
	}
	return result
}

func GetDatabaseNameBySlice(a any) string {
	var result string
	v := reflect.ValueOf(a)
	for i := 0; i < v.Len(); i++ {
		indexValue := v.Index(i)
		if !indexValue.CanInterface() {
			continue
		}
		if indexValue.Kind() == reflect.Struct {
			result = GetDatabaseNameByStruct(indexValue.Interface())
			if len(result) != 0 {
				break
			}
		}
	}
	return result
}

func GetCollectionNameByStruct(a any) string {
	var result string
	v := reflect.ValueOf(a)
	t := reflect.TypeOf(a)
	if v.Kind() == reflect.Interface || v.Kind() == reflect.Pointer {
		v = v.Elem()
		t = t.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		result = t.Field(i).Tag.Get("collection")
		if len(result) != 0 {
			break
		}
	}
	return result
}

func GetCollectionNameBySlice(a any) string {
	var result string
	v := reflect.ValueOf(a)
	for i := 0; i < v.Len(); i++ {
		indexValue := v.Index(i)
		if !indexValue.CanInterface() {
			continue
		}
		if indexValue.Kind() == reflect.Struct {
			result = GetCollectionNameByStruct(indexValue.Interface())
			if len(result) != 0 {
				break
			}
		}
	}
	return result
}

func SetInsertedIdsOnDocuments(insertedIds, as []any) {
	for i, a := range as {
		SetInsertedIdOnDocument(insertedIds[i], a)
	}
}

func SetInsertedIdOnDocument(insertedId, a any) {
	rInsertedId := reflect.ValueOf(insertedId)
	v := reflect.ValueOf(a)
	t := reflect.TypeOf(a)
	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i)
		fieldStruct := t.Field(i)
		if strings.Contains(fieldStruct.Tag.Get("bson"), "_id") && IsObjectId(fieldValue) {
			valueToSet := reflect.New(fieldValue.Type())
			valueToSet.Set(rInsertedId)
		}
	}
}

func MinInt64(value, min int64) int64 {
	if value < min {
		return min
	}
	return value
}
