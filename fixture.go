package fix

import (
	"context"
	"testing"
	"unsafe"
)

type Fixture[V any] interface {
	Value(t *testing.T) V
}

func New[V any](createValue func(t *testing.T) V) Fixture[V] {
	return &fixture[V]{createValue}
}

type fixture[V any] struct {
	createValue func(t *testing.T) V
}

func (f *fixture[V]) Value(t *testing.T) V {
	if value, ok := t.Context().Value(f).(V); ok {
		return value
	}

	value := f.createValue(t)

	// Look away
	ctx := (*context.Context)(unsafe.Add(unsafe.Pointer(t), 400))
	*ctx = context.WithValue(t.Context(), f, value)

	return value
}
