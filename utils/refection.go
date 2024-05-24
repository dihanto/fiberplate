package utils

import "reflect"

func Convert[T interface{}](destStructType T, srcp interface{}) T {
	dest := reflect.New(reflect.TypeOf(destStructType)).Elem()
	srcv := reflect.ValueOf(srcp).Elem()

	destt := dest.Type()
	for i := 0; i < destt.NumField(); i++ {
		sf := destt.Field(i)
		v := srcv.FieldByName(sf.Name)
		if !v.IsValid() || !v.Type().AssignableTo(sf.Type) {
			continue
		}
		dest.Field(i).Set(v)
	}

	return dest.Interface().(T)
}
