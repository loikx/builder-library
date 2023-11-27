package builder

import (
	"errors"
	"fmt"
	"reflect"
)

type Builder[T any] struct {
	item   *T
	typeOf reflect.Type

	err error
}

func NewBuilder[T any]() *Builder[T] {
	var item T

	return &Builder[T]{
		item:   new(T),
		typeOf: reflect.TypeOf(item),
	}
}

func (b *Builder[T]) With(key string, value any) *Builder[T] {
	if b.item == nil {
		return b
	}

	_, ok := b.typeOf.FieldByName(key)
	if !ok {
		b.err = errors.New(
			fmt.Sprintf("does not have field %s with type %v", key, reflect.TypeOf(value)),
		)

		b.item = nil

		return b
	}

	field := reflect.ValueOf(b.item).Elem().FieldByName(key)
	if !field.CanSet() {
		b.err = errors.New(
			fmt.Sprintf("can not set field %s with type %v", key, reflect.TypeOf(value)),
		)

		return b
	}

	field.Set(reflect.ValueOf(value))

	return b
}

func (b *Builder[T]) Build() (*T, error) {
	return b.item, b.err
}
