package order

import "context"

type (
	Input  struct{}
	Output struct{}
)

type Order struct{}

func (oh *Order) Do(ctx context.Context, input Input) (Output, error) {
	return Output{}, nil
}
