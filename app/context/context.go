// Package context provides functions for setting and getting values from
// a request's context.
package context

import (
	"context"
	"net/http"

	"github.com/thecodingmachine/gotenberg/app/converter"
)

type key uint32

const (
	converterKey key = iota
	resultFilePathKey
)

// WithConverter populates a request's context with the given converter
// and returns the updated request.
func WithConverter(r *http.Request, converter *converter.Converter) *http.Request {
	ctx := r.Context()
	ctx = context.WithValue(ctx, converterKey, converter)
	r = r.WithContext(ctx)

	return r
}

type converterNotFoundError struct{}

const converterNotFoundErrorMessage = "The converter was not found in request context"

func (e *converterNotFoundError) Error() string {
	return converterNotFoundErrorMessage
}

// GetConverter returns the converter if found in
// the request's context. Otherwise throws an error.
func GetConverter(r *http.Request) (*converter.Converter, error) {
	c, ok := r.Context().Value(converterKey).(*converter.Converter)
	if !ok {
		return nil, &converterNotFoundError{}
	}

	return c, nil
}

// WithResultFilePath populates a request's context with the given result file path
// and returns the updated request.
func WithResultFilePath(r *http.Request, resultFilePath string) *http.Request {
	ctx := r.Context()
	ctx = context.WithValue(ctx, resultFilePathKey, resultFilePath)
	r = r.WithContext(ctx)

	return r
}

type resultFilePathNotFoundError struct{}

const resultFilePathNotFoundErrorMessage = "The result file path was not found in request context"

func (e *resultFilePathNotFoundError) Error() string {
	return resultFilePathNotFoundErrorMessage
}

// GetResultFilePath returns the result file path if found in
// the request's context. Otherwise throws an error.
func GetResultFilePath(r *http.Request) (string, error) {
	path, ok := r.Context().Value(resultFilePathKey).(string)
	if !ok {
		return "", &resultFilePathNotFoundError{}
	}

	return path, nil
}
