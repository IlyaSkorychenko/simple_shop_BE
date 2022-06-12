package pkg

import "reflect"

func CheckWithCustomError(e error, message string) {
	if e != nil {
		panic(message)
	}
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetType(v any) string {
	return reflect.ValueOf(v).Type().Name()
}
