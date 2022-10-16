package invoker

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"google.golang.org/grpc"
)

type invokef = func() func(ctx context.Context, method string, args interface{}, reply interface{}, _ ...grpc.CallOption) error

type store struct {
	m map[string]invokef
}

func (s *store) Register(fullmethod string, f invokef) error {
	k, err := extractKey(fullmethod)
	if err != nil {
		return err
	}
	s.m[k] = f
	return nil
}

func (s *store) Invokef(fullmethod string) (invokef, error) {
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
