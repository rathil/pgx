package pgx

import "reflect"

// RowToStructByNameReflect like RowToStructByName with return any and using reflect
func RowToStructByNameReflect(proto any) RowToFunc[any] {
	return func(row CollectableRow) (any, error) {
		rv := reflect.New(reflect.TypeOf(proto))
		err := (&namedStructRowScanner{ptrToStruct: rv.Interface()}).ScanRow(row)
		return rv.Elem().Interface(), err
	}
}
