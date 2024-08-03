package cqrs

import (
	"context"
	"fmt"
	"reflect"
)

type Result[T any] struct {
	Value T
	Err   error
}

func (r Result[T]) IsError() bool {
	return r.Err != nil
}

func (r Result[T]) String() string {
	if r.Err != nil {
		return r.Err.Error()
	}

	return fmt.Sprintf("%v", r.Value)
}

type CQRSHandler[TRequest any, TResponse any] interface {
	Handle(TRequest, context.Context) (TResponse, error)
}

var (
	handlers = map[reflect.Type]interface{}{}
)

func RegisterHandler[TRequest any, TResponse any](handler CQRSHandler[TRequest, TResponse]) error {
	var request TRequest
	t := reflect.TypeOf(request)
	if _, exists := handlers[t]; exists {
		return fmt.Errorf("handler already registered for %s", t.String())
	}

	handlers[t] = handler
	return nil
}

func GetHandler[TRequest any, TResponse any](handler CQRSHandler[TRequest, TResponse]) CQRSHandler[TRequest, TResponse] {
	t := reflect.TypeOf(handler)
	if h, ok := handlers[t]; ok {
		return h.(CQRSHandler[TRequest, TResponse])
	}

	return nil
}

func ClearRegistrations() {
	handlers = map[reflect.Type]interface{}{}
}

func Dispatch[TRequest any, TResponse any](request TRequest, cancellationCtx context.Context) Result[TResponse] {
	requestType := reflect.TypeOf(request)
	handler, ok := handlers[requestType]
	if !ok {
		return Result[TResponse]{Err: fmt.Errorf("handler not found for %s", requestType.String())}
	}

	handlerValue, ok := buildHandler[TRequest, TResponse](handler)
	if !ok {
		return Result[TResponse]{Err: fmt.Errorf("handler not found for %s", requestType.String())}
	}

	res, err := handlerValue.Handle(request, cancellationCtx)
	if err != nil {
		return Result[TResponse]{Err: err}
	}

	return Result[TResponse]{Value: res}
}

func buildHandler[TRequest any, TResponse any](handler any) (CQRSHandler[TRequest, TResponse], bool) {
	handlerValue, ok := handler.(CQRSHandler[TRequest, TResponse])
	if !ok {
		return nil, false
	}

	return handlerValue, true
}
