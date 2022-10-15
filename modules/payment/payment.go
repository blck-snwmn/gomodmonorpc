package payment

import "context"

type (
	Input  struct{}
	Output struct{}
)

type Payment struct{}

func (oh *Payment) Do(ctx context.Context, input Input) (Output, error) {
	return Output{}, nil
}
