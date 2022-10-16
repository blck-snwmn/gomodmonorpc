package invoker

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"google.golang.org/grpc"
)

func Cp[T any](out, in *T) {
	*out = *in
}

type invoke = func(ctx context.Context, method string, args interface{}, reply interface{}, _ ...grpc.CallOption) error

type Store struct {
	m map[string]invoke
}

func NewStore() *Store {
	return &Store{m: map[string]invoke{}}
}

func (s *Store) Register(fullmethod string, f invoke) error {
	k, err := extractKey(fullmethod)
	if err != nil {
		return err
	}
	s.m[k] = f
	return nil
}

func (s *Store) RegisterByServiceName(serviceName string, f invoke) {
	s.m[serviceName] = f
}

func (s *Store) Invokef(fullmethod string) (invoke, error) {
	k, err := extractKey(fullmethod)
	if err != nil {
		return nil, err
	}
	f, ok := s.m[k]
	if !ok {
		return nil, fmt.Errorf("%q method doesn't exist", fullmethod)
	}
	return f, nil
}

func extractKey(fullmethod string) (string, error) {
	s := strings.Split(fullmethod, "/")
	if len(s) != 3 {
		return "", errors.New("fullmethod is invalid fomat")
	}
	return s[1], nil
}
