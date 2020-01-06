package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"
)

func IsEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}

func StructJsonTagNameToMapWithJson(obj interface{}) (target map[string]interface{}) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		//return errors.New(fmt.Sprintf("map to json err = %v", err))
		return nil
	}
	err = json.Unmarshal(bytes, &target)
	return
}

func StructJsonTagNameToMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get("json")
		ks := strings.Split(tag, ",")
		var k string
		if len(ks) < 1 {
			continue
		}
		k = ks[0]
		if strings.Contains(tag, "omitempty") {
			if IsEmptyValue(v.Field(i)) {
				continue
			}
		}
		data[k] = v.Field(i).Interface()
	}
	return data
}

func MapToStructJsonTagName(obj map[string]interface{}, target interface{}) error {
	return MapToStructJsonTagNameInner(obj, reflect.Indirect(reflect.ValueOf(target)))
}
func MapToStructJsonTagNameInner(obj map[string]interface{}, rv reflect.Value) error {
	rt := rv.Type()

	for i := 0; i < rt.NumField(); i++ {
		tag := rt.Field(i).Tag.Get("json")
		ks := strings.Split(tag, ",")
		if len(ks) < 1 {
			continue
		}
		jk := ks[0]
		v, ok := obj[jk]
		if !ok {
			//log.Printf("map not found field key = %s", jk)
			continue
		}

		sfv := rv.Field(i)
		if !sfv.IsValid() {
			return fmt.Errorf("No such field: %s in obj", jk)
		}
		//if !sfv.CanAddr() {
		//	return errors.New("Cannot unaddressable")
		//}

		if !sfv.CanSet() {
			return fmt.Errorf("Cannot set %s field value", jk)
		}

		sft := sfv.Type()
		val := reflect.ValueOf(v)
		if sft != val.Type() {
			//log.Printf("field type = %v value type = %v", sfv.Kind(), val.Kind())
			if sfv.Kind() == reflect.Struct && val.Kind() == reflect.Map{
				err := MapToStructJsonTagNameInner(v.(map[string]interface{}), sfv)
				if err != nil {
					log.Printf("field = %s maptostruct err = %v", jk, err)
				}
			}else {
				return errors.New("Provided value type didn't match obj field type")
			}
		}else {
			sfv.Set(val)
		}
	}
	return nil
}

func MapToStructJsonTagNameWithJson(obj map[string]interface{}, target interface{}) error {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return errors.New(fmt.Sprintf("map to json err = %v", err))
	}
	err = json.Unmarshal(bytes, target)
	return err
}

