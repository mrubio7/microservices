package mapper

import (
	"fmt"
	"reflect"
)

func init() {
	RegisterMappers()
}

type Mapper[T any, U any] struct {
	From func(T) U
	To   func(U) T
}

var registry = make(map[string]interface{})

func key[T any, U any]() string {
	return fmt.Sprintf("%s->%s", reflect.TypeOf((*T)(nil)).Elem(), reflect.TypeOf((*U)(nil)).Elem())
}

func Register[T any, U any](mapper Mapper[T, U]) {
	registry[key[T, U]()] = mapper
}

func Convert[From any, To any](input From) To {
	mapperKey := key[From, To]()
	mapper, ok := registry[mapperKey]
	if !ok {
		panic(fmt.Sprintf("No mapper registered for types: %s", mapperKey))
	}
	typedMapper := mapper.(Mapper[From, To])
	return typedMapper.From(input)
}
