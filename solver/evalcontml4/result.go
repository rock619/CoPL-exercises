package main

import "fmt"

type Result[T any] struct {
	val T
	err error
}

func OK[T any](val T) Result[T] {
	return Result[T]{val: val}
}

func Err[T any](err error) Result[T] {
	return Result[T]{err: err}
}

func (r Result[T]) Val() T {
	return r.val
}

func (r Result[T]) Err() error {
	return r.err
}

func (r Result[T]) IsOK() bool {
	return r.err == nil
}

func (r Result[T]) IsErr() bool {
	return !r.IsOK()
}

func (r Result[T]) Unwrap() (T, error) {
	return r.val, r.err
}

func AssertResult[T any](v any) Result[T] {
	r, ok := v.(Result[T])
	if !ok {
		return Err[T](fmt.Errorf("not Result[%T]: got %T", (*new(T)), v))
	}
	return r
}
