package fix

import (
	"context"
	"testing"
	"unsafe"
)

var ctxOffset int

type Fixture[V any] func(t *testing.T) V

func New[V any](createValue func(t *testing.T) V) Fixture[V] {
	f := &fixture[V]{createValue}

	return func(t *testing.T) V {
		return f.value(t)
	}
}

type fixture[V any] struct {
	createValue func(t *testing.T) V
}

func (f *fixture[V]) value(t *testing.T) V {
	if value, ok := t.Context().Value(f).(V); ok {
		return value
	}

	value := f.createValue(t)

	// Look away
	ctx := (*context.Context)(unsafe.Add(unsafe.Pointer(t), ctxOffset))
	*ctx = context.WithValue(t.Context(), f, value)

	return value
}
