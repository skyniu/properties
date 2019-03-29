package properties

import (
	"reflect"
	"fmt"
)


func makeMap(tp reflect.Type,val reflect.Value) map[string]interface{} {
	resultMap:=make(map[string]interface{})
	for idx:=0;idx<tp.NumField();idx++{
		field:=tp.Field(idx)
		fieldVal:=val.Field(idx).Interface()

		tag:=field.Tag.Get("json")
		if tag == ""{
			tag = field.Name
		}
		resultMap[tag] = ObjectToMap(fieldVal)
	}
	return resultMap
}

func makeSlice(tp reflect.Type,val reflect.Value) interface{} {
	fmt.Println(tp.NumField())
	for i:=0;i<tp.NumField();i++{
		fmt.Println(tp.Field(i).Type.String())
	}
	return nil
}

func ObjectToMap(i interface{})interface{} {
	tp := reflect.TypeOf(i)
	switch tp.Kind() {
	case reflect.Struct:
		val := reflect.ValueOf(i)
		return makeMap(tp,val)
	case reflect.Map:
		if resultMap,ok:=i.(map[string]interface{});ok{
			return resultMap
		}
	case reflect.Ptr:
		val := reflect.ValueOf(i)
		return makeMap(tp.Elem(),val.Elem())
	case reflect.Slice:
		makeSlice(tp.Elem(),reflect.ValueOf(i))
	default:
		return i
	}
	return i
}