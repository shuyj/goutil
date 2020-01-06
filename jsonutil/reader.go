package jsonutil

import (
	"encoding/json"
	"errors"
	"log"
	"reflect"
)

type (
	Reader interface {
		HasField(name string) bool
		Int(name string) (int,bool)
		String(name string) (string,bool)
		Float32(name string) (float32,bool)
		Float64(name string) (float64,bool)
		IntArray(name string) ([]int,bool)
		StringArray(name string) ([]string,bool)
		Interface(name string)(interface{}, bool)
		ToReader(name string) (Reader,error)
		ToReaders(name string) ([]Reader,error)
		ToStruct(value interface{})error
	}

	readImpl struct {
		fields map[string]interface{}
	}
)

func NewJsonReader(jsonstr string) (Reader,error) {
	fields := map[string]interface{}{}

	err := json.Unmarshal([]byte(jsonstr), &fields)
	if err != nil {
		return nil,err
	}

	for k,v := range fields{
		log.Printf("ParseJson %s=%T", k, v)
	}

	return &readImpl{
		fields: fields,
	},nil
}

func NewJsonReaderWithDict(dict map[string]interface{}) (Reader,error) {
	//for k,v := range dict{
		//logging.Debugf("ParseJson %s=%T", k, v)
	//}

	return &readImpl{
		fields: dict,
	},nil
}

func (self *readImpl) HasField(name string) bool {
	_, ok := self.fields[name]
	return ok
}

func (self *readImpl) Int(name string) (int,bool){
	val, ok := self.fields[name]
	if ok {
		return digitalToInt(val)
	}
	return 0, false
}
func (self *readImpl) String(name string) (string,bool){
	val, ok := self.fields[name]
	if ok {
		sval,ok := val.(string)
		return sval,ok
	}
	return "",false
}
func (self *readImpl) Float32(name string) (float32,bool){
	val, ok := self.fields[name]
	if ok {
		fval,ok := val.(float32)
		return fval,ok
	}
	return 0.0, false
}
func (self *readImpl) Float64(name string) (float64,bool){
	val, ok := self.fields[name]
	if ok {
		fval,ok := val.(float64)
		return fval,ok
	}
	return 0.0, false
}
func (self *readImpl) IntArray(name string) ([]int,bool){
	val, ok := self.fields[name]
	if ok {
		return interfaceToIntArray(val)
	}
	return nil, false
}

func (self *readImpl) StringArray(name string) ([]string,bool){
	val, ok := self.fields[name]
	if ok {
		return interfaceToStringArray(val)
	}
	return nil, false
}

func (self *readImpl) Interface(name string)(interface{}, bool){
	val, ok := self.fields[name]
	if ok {
		return val, true
	}
	return nil, false
}

func interfaceToIntArray(val interface{}) ([]int, bool) {
	valueOf := reflect.Indirect(reflect.ValueOf(val))
	typeOf := valueOf.Type()
	if typeOf.Kind() != reflect.Slice && typeOf.Kind() != reflect.Array{
		return nil, false
	}
	var ivals []int
	var ok bool = false

	iarr, iok := val.([]interface{})
	if !iok{
		return nil, false
	}

	for _, iv := range iarr {
		switch reflect.TypeOf(iv).Kind() {
		case reflect.Int, reflect.Int64, reflect.Int8, reflect.Int16, reflect.Int32:
			ivals = append(ivals, iv.(int))
			ok = true
		case reflect.Float64:
			fval,_ := iv.(float64)
			ivals = append(ivals, int(fval))
			ok = true
		case reflect.Float32:
			fval,_ := iv.(float32)
			ivals = append(ivals, int(fval))
			ok = true
		}
	}
	return ivals,ok
}

func interfaceToStringArray(val interface{}) ([]string, bool) {
	valueOf := reflect.Indirect(reflect.ValueOf(val))
	typeOf := valueOf.Type()
	if typeOf.Kind() != reflect.Slice && typeOf.Kind() != reflect.Array{
		return nil, false
	}
	var svals []string
	var ok bool = false

	iarr, iok := val.([]interface{})
	if !iok{
		return nil, false
	}

	for _, iv := range iarr {
		switch reflect.TypeOf(iv).Kind() {
		case reflect.String:
			svals = append(svals, iv.(string))
			ok = true
		}
	}
	return svals,ok
}


func digitalToInt(value interface{})(int,bool){
	valueOf := reflect.Indirect(reflect.ValueOf(value))
	typeOf := valueOf.Type()
	var ival int
	var ok bool = false
	switch typeOf.Kind(){
	case reflect.Int, reflect.Int64, reflect.Int8, reflect.Int16, reflect.Int32:
		ival,ok = value.(int)
	case reflect.Float64:
		fval,_ := value.(float64)
		ival = int(fval)
		ok = true
	case reflect.Float32:
		fval,_ := value.(float32)
		ival = int(fval)
		ok = true
	}
	return ival,ok
}

func (self *readImpl) ToStruct(value interface{}) error {
	valueOf := reflect.ValueOf(value)

	if valueOf.Kind() != reflect.Ptr || valueOf.IsNil() {
		return errors.New("ToStruct: expected a pointer as an argument")
	}

	valueOf = valueOf.Elem()
	typeOf := valueOf.Type()

	if valueOf.Kind() != reflect.Struct {
		return errors.New("ToStruct: expected a pointer to struct as an argument")
	}

	for i := 0; i < valueOf.NumField(); i++ {
		fieldType := typeOf.Field(i)
		fieldValue := valueOf.Field(i)

		original, ok := self.fields[fieldType.Name]
		if !ok {
			continue
		}
		valueOf := reflect.Indirect(reflect.ValueOf(original))

		if fieldValue.CanSet() && self.haveSameTypes(valueOf.Type(), fieldValue.Type()) {
			fieldValue.Set(valueOf)
		}
	}

	return nil
}

//func (r readImpl) ToMapReaderOfReaders() map[interface{}]Reader {
//	valueOf := reflect.Indirect(reflect.ValueOf(r.value))
//	typeOf := valueOf.Type()
//
//	if typeOf.Kind() != reflect.Map {
//		return nil
//	}
//
//	readers := map[interface{}]Reader{}
//
//	for _, keyValue := range valueOf.MapKeys() {
//		readers[keyValue.Interface()] = NewReader(valueOf.MapIndex(keyValue).Interface())
//	}
//
//	return readers
//}

func (self *readImpl) haveSameTypes(first reflect.Type, second reflect.Type) bool {
	if first.Kind() != second.Kind() {
		return false
	}

	switch first.Kind() {
	case reflect.Ptr:
		return self.haveSameTypes(first.Elem(), second.Elem())
	case reflect.Struct:
		return first.PkgPath() == second.PkgPath() && first.Name() == second.Name()
	case reflect.Slice:
		return self.haveSameTypes(first.Elem(), second.Elem())
	case reflect.Map:
		return self.haveSameTypes(first.Elem(), second.Elem()) && self.haveSameTypes(first.Key(), second.Key())
	default:
		return first.Kind() == second.Kind()
	}
}

func (self *readImpl) ToReader(name string) (Reader,error){
	val, ok := self.fields[name]
	if ok {
		fval,ok := val.(map[string]interface{})
		if ok {
			return &readImpl{fields: fval}, nil
		}else{
			return nil, errors.New("value is not Reader")
		}
	}
	return nil, errors.New("key is not found")
}
func (self *readImpl) ToReaders(name string) ([]Reader,error){
	val, ok := self.fields[name]
	if ok {
		valueOf := reflect.Indirect(reflect.ValueOf(val))
		typeOf := valueOf.Type()
		if typeOf.Kind() != reflect.Slice && typeOf.Kind() != reflect.Array{
			return nil, errors.New("value is not Slice")
		}
		fval,ok := val.([]map[string]interface{})
		if ok {
			var readers []Reader
			for _,val := range fval{
				readers = append(readers, &readImpl{fields:val})
			}
			return readers, nil
		}else{
			return nil, errors.New("value is not []Reader")
		}
	}
	return nil, errors.New("key is not found")
}
