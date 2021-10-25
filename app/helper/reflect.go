package helper

import "reflect"

// BuildReflectValueParams 构建反射函数触发参数 */
func BuildReflectValueParams(i ...interface{}) []reflect.Value {
	var parameters []reflect.Value
	for _, value := range i {
		parameters = append(parameters, reflect.ValueOf(value))
	}
	return parameters
}

// DynamicValue 动态获取类型值 /**
func DynamicValue(value reflect.Value) interface{} {
	if "int" == value.Type().String() {
		return value.Int()
	}
	return value.String()
}
