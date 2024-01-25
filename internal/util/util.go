package util

import (
	"github.com/GabrielHCataldo/go-helper/helper"
	"reflect"
	"strings"
)

func GetDatabaseNameByStruct(a any) string {
	var result string
	t := reflect.TypeOf(a)
	if helper.IsPointer(a) || helper.IsInterface(a) {
		t = t.Elem()
	}
	for i := 0; helper.IsLessThan(i, t.NumField()); i++ {
		result = t.Field(i).Tag.Get("database")
		if helper.IsNotEmpty(result) {
			break
		}
	}
	return result
}

func GetDatabaseNameBySlice(a any) string {
	v := reflect.ValueOf(a)
	if helper.IsPointer(a) || helper.IsInterface(a) {
		v = v.Elem()
	}
	return GetDatabaseNameByStruct(reflect.New(v.Type().Elem()).Interface())
}

func GetCollectionNameByStruct(a any) string {
	var result string
	t := reflect.TypeOf(a)
	if helper.IsPointer(a) || helper.IsInterface(a) {
		t = t.Elem()
	}
	for i := 0; helper.IsLessThan(i, t.NumField()); i++ {
		result = t.Field(i).Tag.Get("collection")
		if helper.IsNotEmpty(result) {
			break
		}
	}
	return result
}

func GetCollectionNameBySlice(a any) string {
	v := reflect.ValueOf(a)
	if helper.IsPointer(a) || helper.IsInterface(a) {
		v = v.Elem()
	}
	return GetCollectionNameByStruct(reflect.New(v.Type().Elem()).Interface())
}

func SetInsertedIdOnDocument(insertedId, a any) {
	rInsertedId := reflect.ValueOf(insertedId)
	v := reflect.ValueOf(a)
	vr := reflect.ValueOf(a)
	t := reflect.TypeOf(a)
	if helper.IsPointer(a) || helper.IsInterface(a) {
		vr = v.Elem()
		t = t.Elem()
	}
	for i := 0; helper.IsLessThan(i, vr.NumField()); i++ {
		fieldValue := vr.Field(i)
		fieldStruct := t.Field(i)
		if strings.Contains(fieldStruct.Tag.Get("bson"), "_id") &&
			helper.Equals(rInsertedId.Kind(), fieldValue.Kind()) {
			fieldValue.Set(rInsertedId)
		}
	}
}
